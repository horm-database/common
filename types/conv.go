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

package types

import (
	"errors"
	"reflect"
	"strconv"
	"time"
	"unsafe"

	"github.com/horm-database/common/structs"

	"github.com/spf13/cast"
)

// InterfaceToString 接口转字符串
func InterfaceToString(value interface{}) string {
	return cast.ToString(value)
}

// InterfaceToBytes 接口转字节码
func InterfaceToBytes(value interface{}) []byte {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return v
	case *[]byte:
		return *v
	default:
		str := InterfaceToString(value)
		return []byte(str)
	}
}

func InterfaceToBool(value interface{}) bool {
	switch v := value.(type) {
	case []byte:
		b, _ := strconv.ParseBool(BytesToString(v))
		return b
	case *[]byte:
		b, _ := strconv.ParseBool(BytesToString(*v))
		return b
	default:
		return cast.ToBool(value)
	}
}

func InterfaceToInt64(value interface{}) (int64, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case []byte:
		return strconv.ParseInt(BytesToString(v), 10, 64)
	case *[]byte:
		return strconv.ParseInt(BytesToString(*v), 10, 64)
	default:
		return cast.ToInt64E(value)
	}
}

func InterfaceToUint64(value interface{}) (uint64, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case []byte:
		return InterfaceToUint64(BytesToString(v))
	case *[]byte:
		return InterfaceToUint64(BytesToString(*v))
	default:
		return cast.ToUint64E(value)
	}
}

func InterfaceToFloat64(value interface{}) (float64, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case []byte:
		return InterfaceToFloat64(BytesToString(v))
	case *[]byte:
		return InterfaceToFloat64(BytesToString(*v))
	default:
		return cast.ToFloat64E(value)
	}
}

func InterfaceToInt(value interface{}) (int, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case []byte:
		return InterfaceToInt(BytesToString(v))
	case *[]byte:
		return InterfaceToInt(BytesToString(*v))
	default:
		return cast.ToIntE(value)
	}
}

func InterfaceToInt8(value interface{}) (int8, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case []byte:
		return InterfaceToInt8(BytesToString(v))
	case *[]byte:
		return InterfaceToInt8(BytesToString(*v))
	default:
		return cast.ToInt8E(value)
	}
}

func InterfaceToInt16(value interface{}) (int16, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case []byte:
		return InterfaceToInt16(BytesToString(v))
	case *[]byte:
		return InterfaceToInt16(BytesToString(*v))
	default:
		return cast.ToInt16E(value)
	}
}

func InterfaceToInt32(value interface{}) (int32, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case []byte:
		return InterfaceToInt32(BytesToString(v))
	case *[]byte:
		return InterfaceToInt32(BytesToString(*v))
	default:
		return cast.ToInt32E(value)
	}
}

func InterfaceToUint(value interface{}) (uint, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case []byte:
		return InterfaceToUint(BytesToString(v))
	case *[]byte:
		return InterfaceToUint(BytesToString(*v))
	default:
		return cast.ToUintE(value)
	}
}

func InterfaceToUint8(value interface{}) (uint8, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case []byte:
		return InterfaceToUint8(BytesToString(v))
	case *[]byte:
		return InterfaceToUint8(BytesToString(*v))
	default:
		return cast.ToUint8E(value)
	}
}

func InterfaceToUint16(value interface{}) (uint16, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case []byte:
		return InterfaceToUint16(BytesToString(v))
	case *[]byte:
		return InterfaceToUint16(BytesToString(*v))
	default:
		return cast.ToUint16E(value)
	}
}

func InterfaceToUint32(value interface{}) (uint32, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case []byte:
		return InterfaceToUint32(BytesToString(v))
	case *[]byte:
		return InterfaceToUint32(BytesToString(*v))
	default:
		return cast.ToUint32E(value)
	}
}

func InterfaceToTime(value interface{}, layout string, loc ...*time.Location) (time.Time, error) {
	if value == nil {
		return time.Time{}, nil
	}

	l := time.Local
	if len(loc) > 0 {
		l = loc[0]
	}

	return ParseTime(value, l, layout)
}

// InterfaceToArray 接口转数组
func InterfaceToArray(value interface{}) (ret []interface{}, err error) {
	if value == nil {
		return nil, nil
	}

	switch val := value.(type) {
	case []interface{}:
		return val, nil
	case []map[string]interface{}:
		ret = make([]interface{}, len(val))
		for k, v := range val {
			ret[k] = v
		}
		return
	default:
		v := reflect.ValueOf(value)
		if IsNil(v) {
			return nil, nil
		}

		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		if !IsArray(v) {
			return nil, errors.New("value is not array")
		}

		l := v.Len()
		ret = make([]interface{}, l)

		for i := 0; i < l; i++ {
			ret[i] = Interface(v.Index(i))
		}
		return ret, nil
	}
}

// InterfaceToMap 接口转 map
func InterfaceToMap(value interface{}) (map[string]interface{}, error) {
	if value == nil {
		return map[string]interface{}{}, nil
	}

	result, ok := value.(map[string]interface{})
	if ok {
		return result, nil
	}

	result = map[string]interface{}{}

	rv := reflect.Indirect(reflect.ValueOf(value))

	switch rv.Kind() {
	case reflect.Struct:
		ss := structs.GetStructSpec("", rv.Type())
		for _, f := range ss.Fs {
			val := rv.Field(f.I)
			field := f.Column
			if field == "" {
				field = f.Name
			}
			result[field] = Interface(val)
		}
		return result, nil
	case reflect.Map:
		for _, k := range rv.MapKeys() {
			result[k.String()] = Interface(rv.MapIndex(k))
		}

		return result, nil
	default:
		return map[string]interface{}{}, errors.New("value is not map")
	}
}

// InterfaceToStringArray 接口转字符串数组
func InterfaceToStringArray(value interface{}) ([]string, error) {
	if value == nil {
		return nil, nil
	}

	switch v := value.(type) {
	case []string:
		return v, nil
	case []interface{}:
		ret := make([]string, len(v))
		for k, item := range v {
			ret[k] = InterfaceToString(item)
		}
		return ret, nil
	}

	rv := reflect.Indirect(reflect.ValueOf(value))

	if rv.Kind() == reflect.Array || rv.Kind() == reflect.Slice {
		l := rv.Len()
		ret := make([]string, l)

		for i := 0; i < l; i++ {
			ret[i] = InterfaceToString(Interface(rv.Index(i)))
		}

		return ret, nil
	}

	return nil, errors.New("value is not array")
}

// InterfaceToInt64Array 接口转int64数组
func InterfaceToInt64Array(value interface{}) (ret []int64, err error) {
	if value == nil {
		return nil, nil
	}

	switch v := value.(type) {
	case []int64:
		return v, nil
	case []interface{}:
		ret = make([]int64, len(v))
		for k, item := range v {
			ret[k], err = InterfaceToInt64(item)
			if err != nil {
				return nil, err
			}
		}
		return ret, nil
	}

	rv := reflect.Indirect(reflect.ValueOf(value))

	if rv.Kind() == reflect.Array || rv.Kind() == reflect.Slice {
		l := rv.Len()
		ret = make([]int64, l)

		for i := 0; i < l; i++ {
			ret[i], err = InterfaceToInt64(Interface(rv.Index(i)))
			if err != nil {
				return nil, err
			}
		}

		return ret, nil
	}

	return nil, errors.New("value is not array")
}

// InterfaceToUint64Array 接口转 uint64 数组
func InterfaceToUint64Array(value interface{}) (ret []uint64, err error) {
	if value == nil {
		return nil, nil
	}

	switch v := value.(type) {
	case []uint64:
		return v, nil
	case []interface{}:
		ret = make([]uint64, len(v))
		for k, item := range v {
			ret[k], err = InterfaceToUint64(item)
			if err != nil {
				return nil, err
			}
		}
		return ret, nil
	}

	rv := reflect.Indirect(reflect.ValueOf(value))

	if rv.Kind() == reflect.Array || rv.Kind() == reflect.Slice {
		l := rv.Len()
		ret = make([]uint64, l)

		for i := 0; i < l; i++ {
			ret[i], err = InterfaceToUint64(Interface(rv.Index(i)))
			if err != nil {
				return nil, err
			}
		}

		return ret, nil
	}

	return nil, errors.New("value is not array")
}

// InterfaceToFloat64Array 接口转 float64 数组
func InterfaceToFloat64Array(value interface{}) (ret []float64, err error) {
	if value == nil {
		return nil, nil
	}

	switch v := value.(type) {
	case []float64:
		return v, nil
	case []interface{}:
		ret = make([]float64, len(v))
		for k, item := range v {
			ret[k], err = InterfaceToFloat64(item)
			if err != nil {
				return nil, err
			}
		}
		return ret, nil
	}

	rv := reflect.Indirect(reflect.ValueOf(value))

	if rv.Kind() == reflect.Array || rv.Kind() == reflect.Slice {
		l := rv.Len()
		ret = make([]float64, l)

		for i := 0; i < l; i++ {
			ret[i], err = InterfaceToFloat64(Interface(rv.Index(i)))
			if err != nil {
				return nil, err
			}
		}

		return ret, nil
	}

	return nil, errors.New("value is not array")
}

// BytesToString converts byte slice to a string without memory allocation.
//
// Note it may break if the implementation of string or slice header changes in the future go versions.
func BytesToString(b []byte) string {
	/* #nosec G103 */
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes converts string to a byte slice without memory allocation.
//
// Note it may break if the implementation of string or slice header changes in the future go versions.
func StringToBytes(s string) (b []byte) {
	/* #nosec G103 */
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	/* #nosec G103 */
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return b
}
