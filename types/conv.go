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

	"github.com/spf13/cast"
)

// ToString 接口转字符串
func ToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case *string:
		return *v
	case []byte:
		return string(v)
	case *[]byte:
		return string(*v)
	case time.Time:
		return v.Format(time.RFC3339Nano)
	default:
		return cast.ToString(value)
	}
}

// ToBytes 接口转字节码
func ToBytes(value interface{}) []byte {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return v
	case *[]byte:
		return *v
	default:
		str := ToString(value)
		return []byte(str)
	}
}

func ToBool(value interface{}) bool {
	switch v := value.(type) {
	case bool:
		return v
	case *bool:
		return *v
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

func ToInt64(value interface{}) (int64, error) {
	if value == nil {
		return 0, nil
	}
	switch v := value.(type) {
	case int64:
		return v, nil
	case *int64:
		return *v, nil
	case []byte:
		return strconv.ParseInt(BytesToString(v), 10, 64)
	case *[]byte:
		return strconv.ParseInt(BytesToString(*v), 10, 64)
	default:
		return cast.ToInt64E(value)
	}
}

func ToUint64(value interface{}) (uint64, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case uint64:
		return v, nil
	case *uint64:
		return *v, nil
	case []byte:
		return ToUint64(BytesToString(v))
	case *[]byte:
		return ToUint64(BytesToString(*v))
	default:
		return cast.ToUint64E(value)
	}
}

func ToFloat64(value interface{}) (float64, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case float64:
		return v, nil
	case *float64:
		return *v, nil
	case []byte:
		return ToFloat64(BytesToString(v))
	case *[]byte:
		return ToFloat64(BytesToString(*v))
	default:
		return cast.ToFloat64E(value)
	}
}

func ToInt(value interface{}) (int, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case int:
		return v, nil
	case *int:
		return *v, nil
	case []byte:
		return ToInt(BytesToString(v))
	case *[]byte:
		return ToInt(BytesToString(*v))
	default:
		return cast.ToIntE(value)
	}
}

func ToInt8(value interface{}) (int8, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case int8:
		return v, nil
	case *int8:
		return *v, nil
	case []byte:
		return ToInt8(BytesToString(v))
	case *[]byte:
		return ToInt8(BytesToString(*v))
	default:
		return cast.ToInt8E(value)
	}
}

func ToInt16(value interface{}) (int16, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case int16:
		return v, nil
	case *int16:
		return *v, nil
	case []byte:
		return ToInt16(BytesToString(v))
	case *[]byte:
		return ToInt16(BytesToString(*v))
	default:
		return cast.ToInt16E(value)
	}
}

func ToInt32(value interface{}) (int32, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case int32:
		return v, nil
	case *int32:
		return *v, nil
	case []byte:
		return ToInt32(BytesToString(v))
	case *[]byte:
		return ToInt32(BytesToString(*v))
	default:
		return cast.ToInt32E(value)
	}
}

func ToUint(value interface{}) (uint, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case uint:
		return v, nil
	case *uint:
		return *v, nil
	case []byte:
		return ToUint(BytesToString(v))
	case *[]byte:
		return ToUint(BytesToString(*v))
	default:
		return cast.ToUintE(value)
	}
}

func ToUint8(value interface{}) (uint8, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case uint8:
		return v, nil
	case *uint8:
		return *v, nil
	case []byte:
		return ToUint8(BytesToString(v))
	case *[]byte:
		return ToUint8(BytesToString(*v))
	default:
		return cast.ToUint8E(value)
	}
}

func ToUint16(value interface{}) (uint16, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case uint16:
		return v, nil
	case *uint16:
		return *v, nil
	case []byte:
		return ToUint16(BytesToString(v))
	case *[]byte:
		return ToUint16(BytesToString(*v))
	default:
		return cast.ToUint16E(value)
	}
}

func ToUint32(value interface{}) (uint32, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case uint32:
		return v, nil
	case *uint32:
		return *v, nil
	case []byte:
		return ToUint32(BytesToString(v))
	case *[]byte:
		return ToUint32(BytesToString(*v))
	default:
		return cast.ToUint32E(value)
	}
}

// ToArray 接口转数组
func ToArray(value interface{}) (ret []interface{}, err error) {
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
	case []Map:
		ret = make([]interface{}, len(val))
		for k, v := range val {
			ret[k] = v
		}
		return
	case []string:
		ret = make([]interface{}, len(val))
		for k, v := range val {
			ret[k] = v
		}
		return
	case []int:
		ret = make([]interface{}, len(val))
		for k, v := range val {
			ret[k] = v
		}
		return
	case []int64:
		ret = make([]interface{}, len(val))
		for k, v := range val {
			ret[k] = v
		}
		return
	case []uint64:
		ret = make([]interface{}, len(val))
		for k, v := range val {
			ret[k] = v
		}
		return
	case []float32:
		ret = make([]interface{}, len(val))
		for k, v := range val {
			ret[k] = v
		}
		return
	case []float64:
		ret = make([]interface{}, len(val))
		for k, v := range val {
			ret[k] = v
		}
		return
	case []byte:
		return nil, errors.New("value is a bytes")
	default:
		v := reflect.Indirect(reflect.ValueOf(value))
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

// ToMap 接口转 map
func ToMap(value interface{}, tag string, op ...int8) (Map, error) {
	if value == nil {
		return nil, nil
	}

	val := Indirect(value)
	switch v := val.(type) {
	case Map:
		return v, nil
	case map[string]interface{}:
		return v, nil
	}

	rv := reflect.ValueOf(val)
	if IsStruct(rv.Type()) {
		return StructToMap(rv, tag, op...), nil
	}

	if rv.Kind() != reflect.Map {
		return nil, errors.New("value is not map")
	}

	// 将非 string 的 key 转化为 string
	ret := make(map[string]interface{}, rv.Len())
	for _, k := range rv.MapKeys() {
		ret[ToString(Interface(k))] = Interface(rv.MapIndex(k))
	}

	return ret, nil
}

// ToMapArray 接口转map数组
func ToMapArray(value interface{}, tag string, op ...int8) (ret []Map, err error) {
	if value == nil {
		return nil, nil
	}

	val := Indirect(value)

	switch v := val.(type) {
	case []Map:
		return v, nil
	case []map[string]interface{}:
		ret = make([]Map, len(v))
		for i, m := range v {
			ret[i] = m
		}
		return ret, nil
	}

	rv := reflect.ValueOf(value)
	if IsStructArray(rv) {
		v := StructsToMaps(rv, tag, op...)
		ret = make([]Map, len(v))
		for i, m := range v {
			ret[i] = m
		}
		return ret, nil
	}

	return nil, errors.New("value is not map array")
}

// ToStringArray 接口转字符串数组
func ToStringArray(value interface{}) ([]string, error) {
	if value == nil {
		return nil, nil
	}

	switch v := value.(type) {
	case []string:
		return v, nil
	case []interface{}:
		ret := make([]string, len(v))
		for k, item := range v {
			ret[k] = ToString(item)
		}
		return ret, nil
	}

	rv := reflect.Indirect(reflect.ValueOf(value))
	if IsArray(rv) {
		l := rv.Len()
		ret := make([]string, l)

		for i := 0; i < l; i++ {
			ret[i] = ToString(Interface(rv.Index(i)))
		}

		return ret, nil
	}

	return nil, errors.New("value is not array")
}

// ToInt64Array 接口转int64数组
func ToInt64Array(value interface{}) (ret []int64, err error) {
	if value == nil {
		return nil, nil
	}

	switch v := value.(type) {
	case []int64:
		return v, nil
	case []interface{}:
		ret = make([]int64, len(v))
		for k, item := range v {
			ret[k], err = ToInt64(item)
			if err != nil {
				return nil, err
			}
		}
		return ret, nil
	}

	rv := reflect.Indirect(reflect.ValueOf(value))
	if IsArray(rv) {
		l := rv.Len()
		ret = make([]int64, l)

		for i := 0; i < l; i++ {
			ret[i], err = ToInt64(Interface(rv.Index(i)))
			if err != nil {
				return nil, err
			}
		}

		return ret, nil
	}

	return nil, errors.New("value is not array")
}

// ToUint64Array 接口转 uint64 数组
func ToUint64Array(value interface{}) (ret []uint64, err error) {
	if value == nil {
		return nil, nil
	}

	switch v := value.(type) {
	case []uint64:
		return v, nil
	case []interface{}:
		ret = make([]uint64, len(v))
		for k, item := range v {
			ret[k], err = ToUint64(item)
			if err != nil {
				return nil, err
			}
		}
		return ret, nil
	}

	rv := reflect.Indirect(reflect.ValueOf(value))
	if IsArray(rv) {
		l := rv.Len()
		ret = make([]uint64, l)

		for i := 0; i < l; i++ {
			ret[i], err = ToUint64(Interface(rv.Index(i)))
			if err != nil {
				return nil, err
			}
		}

		return ret, nil
	}

	return nil, errors.New("value is not array")
}

// ToFloat64Array 接口转 float64 数组
func ToFloat64Array(value interface{}) (ret []float64, err error) {
	if value == nil {
		return nil, nil
	}

	switch v := value.(type) {
	case []float64:
		return v, nil
	case []interface{}:
		ret = make([]float64, len(v))
		for k, item := range v {
			ret[k], err = ToFloat64(item)
			if err != nil {
				return nil, err
			}
		}
		return ret, nil
	}

	rv := reflect.Indirect(reflect.ValueOf(value))
	if IsArray(rv) {
		l := rv.Len()
		ret = make([]float64, l)

		for i := 0; i < l; i++ {
			ret[i], err = ToFloat64(Interface(rv.Index(i)))
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
