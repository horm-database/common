// Copyright (c) 2024 The horm-database Authors (such as CaoHao <18500482693@163.com>). All rights reserved.
//
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
	"fmt"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Levels is the map from string to zapcore.Level.
var Levels = map[string]zapcore.Level{
	"":         zapcore.DebugLevel,
	LevelDebug: zapcore.DebugLevel,
	LevelInfo:  zapcore.InfoLevel,
	LevelWarn:  zapcore.WarnLevel,
	LevelError: zapcore.ErrorLevel,
	LevelFatal: zapcore.FatalLevel,
}

// newZapLog new a zap logger
func newZapLog(c []*Config) Logger {
	var cores []zapcore.Core

	for _, o := range c {
		writer := GetWriter(o.Writer)
		if writer == nil {
			panic("log writer " + o.Writer + " no registered")
		}

		core, err := writer.Setup(o)
		if err != nil {
			panic("log writer " + o.Writer + " setup error: " + err.Error())
		}
		cores = append(cores, core)
	}

	return &zapLog{logger: zap.New(zapcore.NewTee(cores...))}
}

func newEncoder(c *Config) zapcore.Encoder {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        GetLogEncoderKey("time", c.EncoderConfig.TimeKey),
		LevelKey:       GetLogEncoderKey("level", c.EncoderConfig.LevelKey),
		MessageKey:     GetLogEncoderKey("msg", c.EncoderConfig.MessageKey),
		NameKey:        GetLogEncoderKey(zapcore.OmitKey, c.EncoderConfig.NameKey),
		CallerKey:      GetLogEncoderKey(zapcore.OmitKey, c.EncoderConfig.CallerKey),
		FunctionKey:    GetLogEncoderKey(zapcore.OmitKey, c.EncoderConfig.FunctionKey),
		StacktraceKey:  GetLogEncoderKey(zapcore.OmitKey, c.EncoderConfig.StacktraceKey),
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     NewTimeEncoder(c.EncoderConfig.TimeFmt),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	switch c.Encoder {
	case "separator":
		return newSepEncoder(encoderCfg, c.Field, c.EncoderConfig.MessageKey, c.Escape)
	case "json":
		return zapcore.NewJSONEncoder(encoderCfg)
	default:
		return zapcore.NewConsoleEncoder(encoderCfg)
	}
}

// GetLogEncoderKey gets user defined log output name, uses defKey if empty.
func GetLogEncoderKey(defKey, key string) string {
	if key == "" {
		return defKey
	}
	return key
}

func newConsoleCore(c *Config) zapcore.Core {
	return zapcore.NewCore(
		newEncoder(c),
		zapcore.Lock(os.Stdout),
		zap.NewAtomicLevelAt(Levels[c.Level]))
}

func newFileCore(c *Config) zapcore.Core {
	writer := &lumberjack.Logger{
		Filename:   c.FileConfig.Filename,
		MaxSize:    c.FileConfig.MaxSize,
		MaxBackups: c.FileConfig.MaxBackups,
		MaxAge:     c.FileConfig.MaxDay,
		LocalTime:  c.FileConfig.LocalTime,
		Compress:   c.FileConfig.Compress,
	}

	var ws zapcore.WriteSyncer

	if c.FileConfig.WriteMode == WriteSync {
		ws = zapcore.AddSync(writer)
	} else {
		ws = NewAsyncFileWriter(writer, c.FileConfig.WriteMode == WriteFast)
	}

	return zapcore.NewCore(
		newEncoder(c), ws,
		zap.NewAtomicLevelAt(Levels[c.Level]))
}

// NewTimeEncoder 时间编码格式
func NewTimeEncoder(format string) zapcore.TimeEncoder {
	switch format {
	case "":
		return func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendByteString(DefaultTimeFormat(t))
		}
	case "seconds":
		return zapcore.EpochTimeEncoder
	case "milliseconds":
		return zapcore.EpochMillisTimeEncoder
	case "standard":
		return zapcore.ISO8601TimeEncoder
	default:
		return func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(format))
		}
	}
}

// zapLog is a Logger implementation based on zaplogger.
type zapLog struct {
	logger *zap.Logger
}

func (l *zapLog) With(fields ...Field) Logger {
	if len(fields) == 0 {
		return l
	}

	return &zapLog{logger: l.logger.With(getZapField(fields...)...)}
}

// Debug logs to DEBUG log.
func (l *zapLog) Debug(msg string, fields ...Field) {
	if l.logger.Core().Enabled(zapcore.DebugLevel) {
		l.logger.Debug(msg, getZapField(fields...)...)
	}
}

// Info logs to INFO log.
func (l *zapLog) Info(msg string, fields ...Field) {
	if l.logger.Core().Enabled(zapcore.InfoLevel) {
		l.logger.Info(msg, getZapField(fields...)...)
	}
}

// Warn logs to WARNING log.
func (l *zapLog) Warn(msg string, fields ...Field) {
	if l.logger.Core().Enabled(zapcore.WarnLevel) {
		l.logger.Warn(msg, getZapField(fields...)...)
	}
}

// Error logs to ERROR log.
func (l *zapLog) Error(msg string, fields ...Field) {
	if l.logger.Core().Enabled(zapcore.ErrorLevel) {
		l.logger.Error(msg, getZapField(fields...)...)
	}
}

// Fatal logs to FATAL log.
func (l *zapLog) Fatal(msg string, fields ...Field) {
	if l.logger.Core().Enabled(zapcore.FatalLevel) {
		l.logger.Fatal(msg, getZapField(fields...)...)
	}
}

// Debugf debug logs
func (l *zapLog) Debugf(msg string, args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.DebugLevel) {
		l.logger.Debug(fmt.Sprintf(msg, args...))
	}
}

// Infof info logs
func (l *zapLog) Infof(msg string, args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.InfoLevel) {
		l.logger.Info(fmt.Sprintf(msg, args...))
	}
}

// Warnf warn logs
func (l *zapLog) Warnf(msg string, args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.WarnLevel) {
		l.logger.Warn(fmt.Sprintf(msg, args...))
	}
}

// Errorf error logs
func (l *zapLog) Errorf(msg string, args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.ErrorLevel) {
		l.logger.Error(fmt.Sprintf(msg, args...))
	}
}

// Fatalf fatal logs
func (l *zapLog) Fatalf(msg string, args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.FatalLevel) {
		l.logger.Fatal(fmt.Sprintf(msg, args...))
	}
}

// Sync calls the zap log's Sync method, and flushes any buffered log entries.
// Applications should take care to call Sync before exiting.
func (l *zapLog) Sync() error {
	return l.logger.Sync()
}

func getZapField(fields ...Field) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for k, field := range fields {
		zapFields[k] = zap.Any(field.Key, field.Value)
	}
	return zapFields
}

// DefaultTimeFormat 默认时间编码格式 2006-01-02 15:04:05.000
func DefaultTimeFormat(t time.Time) []byte {
	t = t.Local()
	year, month, day := t.Date()
	hour, minute, second := t.Clock()
	micros := t.Nanosecond() / 1000

	buf := make([]byte, 23)
	buf[0] = byte((year/1000)%10) + '0'
	buf[1] = byte((year/100)%10) + '0'
	buf[2] = byte((year/10)%10) + '0'
	buf[3] = byte(year%10) + '0'
	buf[4] = '-'
	buf[5] = byte((month)/10) + '0'
	buf[6] = byte((month)%10) + '0'
	buf[7] = '-'
	buf[8] = byte((day)/10) + '0'
	buf[9] = byte((day)%10) + '0'
	buf[10] = ' '
	buf[11] = byte((hour)/10) + '0'
	buf[12] = byte((hour)%10) + '0'
	buf[13] = ':'
	buf[14] = byte((minute)/10) + '0'
	buf[15] = byte((minute)%10) + '0'
	buf[16] = ':'
	buf[17] = byte((second)/10) + '0'
	buf[18] = byte((second)%10) + '0'
	buf[19] = '.'
	buf[20] = byte((micros/100000)%10) + '0'
	buf[21] = byte((micros/10000)%10) + '0'
	buf[22] = byte((micros/1000)%10) + '0'
	return buf
}
