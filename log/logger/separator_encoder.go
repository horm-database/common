// Copyright (c) 2024 The horm-database Authors. All rights reserved.
// This file Author:  CaoHao <18500482693@163.com> .
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logger

import (
	"encoding/base64"
	"math"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/json-iterator/go"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

const _hex = "0123456789abcdef"
const maxLength = 60 * 1024 // 包大小限制

var bufferPool = buffer.NewPool()
var nullLiteralBytes = []byte("-") //空字符采用什么代替
var separatorBytes = []byte("\t")  //分隔符

var _encoderPool = sync.Pool{
	New: func() interface{} {
		return &separatorEncoder{}
	},
}

func getEncoder() *separatorEncoder {
	return _encoderPool.Get().(*separatorEncoder)
}

func putEncoder(enc *separatorEncoder) {
	if enc.reflectBuf != nil {
		enc.reflectBuf.Free()
	}

	enc.EncoderConfig = nil
	enc.field = nil
	enc.messageKey = ""
	enc.escape = false

	enc.buf = nil
	enc.fieldMap = nil

	enc.reflectBuf = nil
	enc.reflectEnc = nil

	_encoderPool.Put(enc)
}

type separatorEncoder struct {
	*zapcore.EncoderConfig
	field      []string // 打印字段，严格有序
	messageKey string
	escape     bool // 是否转义

	buf      *buffer.Buffer    // encode过程中复用的buffer
	fieldMap map[string]string // 每个field对应的log内容

	// for encoding generic values by reflection
	reflectBuf *buffer.Buffer
	reflectEnc *jsoniter.Encoder
}

func newSepEncoder(cfg zapcore.EncoderConfig, field []string, messageKey string, escape bool) *separatorEncoder {
	return &separatorEncoder{
		EncoderConfig: &cfg,
		field:         field,
		messageKey:    messageKey,
		escape:        escape,
		buf:           bufferPool.Get(),
		fieldMap:      make(map[string]string),
	}
}

// EncodeEntry 日志数据 encode 入口
func (enc *separatorEncoder) EncodeEntry(ent zapcore.Entry, field []zapcore.Field) (*buffer.Buffer, error) {
	final := enc.clone()

	if final.LevelKey != "" {
		cur := final.buf.Len()
		final.EncodeLevel(ent.Level, final)
		if cur == final.buf.Len() {
			// User-supplied EncodeLevel was a no-op. Fall back to strings to keep
			// output JSON valid.
			final.AppendString(ent.Level.String())
		}

		final.saveBuf(final.LevelKey)
	}
	if final.TimeKey != "" {
		final.AddTime(final.TimeKey, ent.Time)
	}
	if ent.LoggerName != "" && final.NameKey != "" {
		nameEncoder := final.EncodeName

		// if no name encoder provided, fall back to FullNameEncoder for backwards
		// compatibility
		if nameEncoder == nil {
			nameEncoder = zapcore.FullNameEncoder
		}

		cur := final.buf.Len()
		nameEncoder(ent.LoggerName, final)
		if cur == final.buf.Len() {
			// User-supplied EncodeName was a no-op. Fall back to strings to
			// keep output JSON valid.
			final.AppendString(ent.LoggerName)
		}
		final.saveBuf(final.NameKey)
	}
	if ent.Caller.Defined && final.CallerKey != "" {
		cur := final.buf.Len()
		final.EncodeCaller(ent.Caller, final)
		if cur == final.buf.Len() {
			// User-supplied EncodeCaller was a no-op. Fall back to strings to
			// keep output JSON valid.
			final.AppendString(ent.Caller.String())
		}
		final.saveBuf(final.CallerKey)
	}
	if final.MessageKey != "" {
		if enc.escape {
			final.AddString(enc.MessageKey, ent.Message)
		} else {
			final.fieldMap[enc.MessageKey] = ent.Message
		}
	}

	if len(enc.fieldMap) > 0 {
		for k, v := range enc.fieldMap {
			final.fieldMap[k] = v
		}
	}

	addFields(final, field)

	if ent.Stack != "" && final.StacktraceKey != "" {
		final.AddString(final.StacktraceKey, ent.Stack)
	}

	final.generateLog(0)

	l := final.buf.Len()
	if l > maxLength {
		final.reset()

		// 日志截断msg字段
		extraLength := l - maxLength
		final.generateLog(extraLength)
	}

	ret := final.buf
	putEncoder(final)

	return ret, nil
}

func (enc *separatorEncoder) saveBuf(key string) {
	enc.fieldMap[key] = enc.buf.String()
	enc.reset()
}

// AddArray encode 数组
func (enc *separatorEncoder) AddArray(key string, arr zapcore.ArrayMarshaler) error {
	err := arr.MarshalLogArray(enc)
	if err != nil {
		return err
	}

	enc.saveBuf(key)
	return nil
}

// AddObject encode 结构体
func (enc *separatorEncoder) AddObject(key string, obj zapcore.ObjectMarshaler) error {
	err := obj.MarshalLogObject(enc)
	if err != nil {
		return err
	}

	enc.saveBuf(key)
	return nil
}

// AddBinary encode 二进制
func (enc *separatorEncoder) AddBinary(key string, val []byte) {
	enc.AddString(key, base64.StdEncoding.EncodeToString(val))
}

// AddByteString encode byte字符串
func (enc *separatorEncoder) AddByteString(key string, val []byte) {
	enc.AppendByteString(val)
	enc.saveBuf(key)
}

// AddBool encode bool
func (enc *separatorEncoder) AddBool(key string, val bool) {
	enc.AppendBool(val)

	enc.saveBuf(key)
}

// AddComplex128 encode Complex128
func (enc *separatorEncoder) AddComplex128(key string, val complex128) {
	enc.AppendComplex128(val)
	enc.saveBuf(key)
}

// AddDuration encode Duration
func (enc *separatorEncoder) AddDuration(key string, val time.Duration) {
	enc.AppendDuration(val)

	enc.saveBuf(key)
}

// AddFloat64 encode float64
func (enc *separatorEncoder) AddFloat64(key string, val float64) {
	enc.AppendFloat64(val)
	enc.saveBuf(key)
}

// AddInt64 encode int64
func (enc *separatorEncoder) AddInt64(key string, val int64) {
	enc.AppendInt64(val)
	enc.saveBuf(key)
}

// AddReflected encode interface
func (enc *separatorEncoder) AddReflected(key string, obj interface{}) error {
	err := enc.AppendReflected(obj)
	if err != nil {
		return err
	}

	enc.saveBuf(key)
	return nil
}

// OpenNamespace 命名空间，不需要
func (enc *separatorEncoder) OpenNamespace(key string) {
}

// AddString encode 字符串
func (enc *separatorEncoder) AddString(key, val string) {
	enc.AppendString(val)
	enc.saveBuf(key)
}

// AddTime encode time
func (enc *separatorEncoder) AddTime(key string, val time.Time) {
	enc.AppendTime(val)
	enc.saveBuf(key)
}

// AddUint64 encode uint64
func (enc *separatorEncoder) AddUint64(key string, val uint64) {
	enc.AppendUint64(val)
	enc.saveBuf(key)
}

// AddComplex64 encode complex64
func (enc *separatorEncoder) AddComplex64(k string, v complex64) { enc.AddComplex128(k, complex128(v)) }

// AddFloat32 encode float32
func (enc *separatorEncoder) AddFloat32(k string, v float32) { enc.AddFloat64(k, float64(v)) }

// AddInt encode int
func (enc *separatorEncoder) AddInt(k string, v int) { enc.AddInt64(k, int64(v)) }

// AddInt32 encode int32
func (enc *separatorEncoder) AddInt32(k string, v int32) { enc.AddInt64(k, int64(v)) }

// AddInt16 encode int16
func (enc *separatorEncoder) AddInt16(k string, v int16) { enc.AddInt64(k, int64(v)) }

// AddInt8 encode int8
func (enc *separatorEncoder) AddInt8(k string, v int8) { enc.AddInt64(k, int64(v)) }

// AddUint encode uint
func (enc *separatorEncoder) AddUint(k string, v uint) { enc.AddUint64(k, uint64(v)) }

// AddUint32 encode uint32
func (enc *separatorEncoder) AddUint32(k string, v uint32) { enc.AddUint64(k, uint64(v)) }

// AddUint16 encode uint16
func (enc *separatorEncoder) AddUint16(k string, v uint16) { enc.AddUint64(k, uint64(v)) }

// AddUint8 encode uint8
func (enc *separatorEncoder) AddUint8(k string, v uint8) { enc.AddUint64(k, uint64(v)) }

// AddUintptr encode uintptr
func (enc *separatorEncoder) AddUintptr(k string, v uintptr) { enc.AddUint64(k, uint64(v)) }

// Clone separatorEncoder拷贝
func (enc *separatorEncoder) Clone() zapcore.Encoder {
	clone := enc.clone()
	clone.buf.Write(enc.buf.Bytes())
	return clone
}

func (enc *separatorEncoder) clone() *separatorEncoder {
	clone := getEncoder()
	clone.EncoderConfig = enc.EncoderConfig
	clone.field = enc.field
	clone.messageKey = enc.messageKey
	clone.escape = enc.escape
	clone.buf = bufferPool.Get()
	clone.fieldMap = make(map[string]string)
	for k, v := range enc.fieldMap {
		clone.fieldMap[k] = v
	}
	return clone
}

func addFields(enc zapcore.ObjectEncoder, field []zapcore.Field) {
	for i := range field {
		field[i].AddTo(enc)
	}
}

func (enc *separatorEncoder) generateLog(extraLength int) {
	endIndex := len(enc.field) - 1

	for i, field := range enc.field {
		b, ok := enc.fieldMap[field]
		if ok {
			if field == enc.messageKey {
				remainLength := len(b) - extraLength
				if remainLength > 0 {
					enc.buf.Write([]byte(b[0:remainLength]))
				} else {
					enc.buf.Write(nullLiteralBytes)
				}
			} else {
				enc.buf.Write([]byte(b))
			}
		} else {
			enc.buf.Write(nullLiteralBytes)
		}
		if i != endIndex {
			enc.buf.Write(separatorBytes)
		}
	}

	enc.buf.WriteByte('\n')
}

func (enc *separatorEncoder) reset() {
	enc.buf.Reset()
}

func (enc *separatorEncoder) resetReflectBuf() {
	if enc.reflectBuf == nil {
		enc.reflectBuf = bufferPool.Get()
		enc.reflectEnc = jsoniter.NewEncoder(enc.reflectBuf)

		// For consistency with our custom JSON encoder.
		enc.reflectEnc.SetEscapeHTML(false)
	} else {
		enc.reflectBuf.Reset()
	}
}

// AppendArray encode 数组
func (enc *separatorEncoder) AppendArray(arr zapcore.ArrayMarshaler) error {
	return arr.MarshalLogArray(enc)
}

// AppendObject encode 结构体
func (enc *separatorEncoder) AppendObject(obj zapcore.ObjectMarshaler) error {
	//return proto.MarshalLogObject(enc)
	return nil
}

// AppendBool encode bool
func (enc *separatorEncoder) AppendBool(val bool) {
	enc.buf.AppendBool(val)
}

// AppendByteString encode bytes
func (enc *separatorEncoder) AppendByteString(val []byte) {
	enc.safeAddByteString(val)
}

// AppendComplex128 encode complex128
func (enc *separatorEncoder) AppendComplex128(val complex128) {
	// Cast to enc platform-independent, fixed-size type.
	r, i := float64(real(val)), float64(imag(val))
	enc.buf.AppendByte('"')
	// Because we're always in enc quoted string, we can use strconv without
	// special-casing NaN and +/-Inf.
	enc.buf.AppendFloat(r, 64)
	enc.buf.AppendByte('+')
	enc.buf.AppendFloat(i, 64)
	enc.buf.AppendByte('i')
	enc.buf.AppendByte('"')
}

// AppendDuration encode time.Duration
func (enc *separatorEncoder) AppendDuration(val time.Duration) {
	cur := enc.buf.Len()
	enc.EncodeDuration(val, enc)
	if cur == enc.buf.Len() {
		// User-supplied EncodeDuration is enc no-op. Fall back to nanoseconds to keep
		// JSON valid.
		enc.AppendInt64(int64(val))
	}
}

// AppendInt64 encode int64
func (enc *separatorEncoder) AppendInt64(val int64) {
	enc.buf.AppendInt(val)
}

// Only invoke the standard JSON encoder if there is actually something to
// encode; otherwise write JSON null literal directly.
func (enc *separatorEncoder) encodeReflected(obj interface{}) ([]byte, error) {
	if obj == nil {
		return nullLiteralBytes, nil
	}
	enc.resetReflectBuf()
	if err := enc.reflectEnc.Encode(obj); err != nil {
		return nil, err
	}
	enc.reflectBuf.TrimNewline()
	return enc.reflectBuf.Bytes(), nil
}

// AppendReflected encode interface
func (enc *separatorEncoder) AppendReflected(val interface{}) error {
	valueBytes, err := enc.encodeReflected(val)
	if err != nil {
		return err
	}
	_, err = enc.buf.Write(valueBytes)
	return err
}

// AppendString encode string
func (enc *separatorEncoder) AppendString(val string) {
	enc.safeAddString(val)
}

// AppendTime encode time.Time
func (enc *separatorEncoder) AppendTime(val time.Time) {
	cur := enc.buf.Len()
	enc.EncodeTime(val, enc)
	if cur == enc.buf.Len() {
		enc.AppendInt64(val.UnixNano())
	}
}

// AppendUint64 encode uint64
func (enc *separatorEncoder) AppendUint64(val uint64) {
	enc.buf.AppendUint(val)
}

func (enc *separatorEncoder) appendFloat(val float64, bitSize int) {
	switch {
	case math.IsNaN(val):
		enc.buf.AppendString(`"NaN"`)
	case math.IsInf(val, 1):
		enc.buf.AppendString(`"+Inf"`)
	case math.IsInf(val, -1):
		enc.buf.AppendString(`"-Inf"`)
	default:
		enc.buf.AppendFloat(val, bitSize)
	}
}

// safeAddString escapes a string and appends it to the internal buffer.
func (enc *separatorEncoder) safeAddString(s string) {
	for i := 0; i < len(s); {
		if enc.tryAddRuneSelf(s[i]) {
			i++
			continue
		}
		r, size := utf8.DecodeRuneInString(s[i:])
		if enc.tryAddRuneError(r, size) {
			i++
			continue
		}
		enc.buf.AppendString(s[i : i+size])
		i += size
	}
}

// safeAddByteString is no-alloc equivalent of safeAddString(string(s)) for s []byte.
func (enc *separatorEncoder) safeAddByteString(s []byte) {
	for i := 0; i < len(s); {
		if enc.tryAddRuneSelf(s[i]) {
			i++
			continue
		}
		r, size := utf8.DecodeRune(s[i:])
		if enc.tryAddRuneError(r, size) {
			i++
			continue
		}
		enc.buf.Write(s[i : i+size])
		i += size
	}
}

func (enc *separatorEncoder) tryAddRuneSelf(b byte) bool {
	if b >= utf8.RuneSelf { //多字节，例如中文，utf8.DecodeRuneInString/utf8.DecodeRune 处理
		return false
	}
	if 0x20 <= b && b != '\\' && b != '|' { //除 '\' 和 '|' // 可见字符
		enc.buf.AppendByte(b)
		return true
	}
	switch b { //不可见字符
	case '\\':
		enc.buf.AppendByte('\\')
		enc.buf.AppendByte('\\')
	case '|':
		enc.buf.AppendByte('\\')
		enc.buf.AppendByte('|')
	case '\n':
		enc.buf.AppendByte('\\')
		enc.buf.AppendByte('n')
	case '\r':
		enc.buf.AppendByte('\\')
		enc.buf.AppendByte('r')
	case '\t': //日志里面 tab 转为 4个空格
		enc.buf.AppendByte(' ')
		enc.buf.AppendByte(' ')
		enc.buf.AppendByte(' ')
		enc.buf.AppendByte(' ')
	default:
		// 除上面转义之外的不可见字符
		enc.buf.AppendString(`\u00`)
		enc.buf.AppendByte(_hex[b>>4])
		enc.buf.AppendByte(_hex[b&0xF])
	}
	return true
}

func (enc *separatorEncoder) tryAddRuneError(r rune, size int) bool {
	if r == utf8.RuneError && size == 1 {
		enc.buf.AppendString(`\uffd`)
		return true
	}
	return false
}

// AppendComplex64 encode complex64
func (enc *separatorEncoder) AppendComplex64(v complex64) { enc.AppendComplex128(complex128(v)) }

// AppendFloat64 encode float64
func (enc *separatorEncoder) AppendFloat64(v float64) { enc.appendFloat(v, 64) }

// AppendFloat32 encode float32
func (enc *separatorEncoder) AppendFloat32(v float32) { enc.appendFloat(float64(v), 32) }

// AppendInt encode int
func (enc *separatorEncoder) AppendInt(v int) { enc.AppendInt64(int64(v)) }

// AppendInt32 encode int32
func (enc *separatorEncoder) AppendInt32(v int32) { enc.AppendInt64(int64(v)) }

// AppendInt16 encode int16
func (enc *separatorEncoder) AppendInt16(v int16) { enc.AppendInt64(int64(v)) }

// AppendInt8 encode int8
func (enc *separatorEncoder) AppendInt8(v int8) { enc.AppendInt64(int64(v)) }

// AppendUint encode uint
func (enc *separatorEncoder) AppendUint(v uint) { enc.AppendUint64(uint64(v)) }

// AppendUint32 encode uint32
func (enc *separatorEncoder) AppendUint32(v uint32) { enc.AppendUint64(uint64(v)) }

// AppendUint16 encode uint16
func (enc *separatorEncoder) AppendUint16(v uint16) { enc.AppendUint64(uint64(v)) }

// AppendUint8 encode uint8
func (enc *separatorEncoder) AppendUint8(v uint8) { enc.AppendUint64(uint64(v)) }

// AppendUintptr encode uintptr
func (enc *separatorEncoder) AppendUintptr(v uintptr) { enc.AppendUint64(uint64(v)) }
