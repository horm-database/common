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
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/araddon/dateparse"
)

// IsArray 判断是否 Array 或 Slice
func IsArray(v reflect.Value) bool {
	return v.Kind() == reflect.Slice || v.Kind() == reflect.Array
}

// IsNil 判断变量是否为 nil
func IsNil(v reflect.Value) bool {
	k := v.Kind()
	if k == reflect.Chan || k == reflect.Func || k == reflect.Map || k == reflect.Pointer ||
		k == reflect.UnsafePointer || k == reflect.Interface || k == reflect.Slice {
		return v.IsNil()
	}
	return false
}

// Interface returns v's current value as an interface{}.
func Interface(v reflect.Value) interface{} {
	if IsNil(v) {
		return nil
	}

	return v.Interface()
}

// IsEmpty val 是否是空值
func IsEmpty(v reflect.Value) bool {
	vk := v.Kind()
	switch vk {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface:
		return v.IsNil()
	case reflect.Ptr:
		if v.IsNil() {
			return true
		} else {
			return IsEmpty(v.Elem())
		}
	default:
		t, isTime := GetRealTime(v)
		if isTime {
			return t.IsZero()
		}
	}

	return v.IsZero()
}

type Time time.Time

// UnmarshalJSON NullTime 类型实现 json marshal 方法
func (nt *Time) UnmarshalJSON(data []byte) error {
	tStr := strings.TrimPrefix(string(data), `"`)
	tStr = strings.TrimSuffix(tStr, `"`)

	if tStr == "null" {
		return nil
	}

	t, err := dateparse.ParseLocal(tStr)
	if err != nil {
		return err
	}
	*nt = Time(t)
	return nil
}

var typeTimes = []reflect.Type{
	reflect.TypeOf(time.Time{}),
	reflect.TypeOf(Time{}),
}

// RegisterTime 注册新的时间类型，例如上面的 Time，底层必须是 time.Time 类型
func RegisterTime(t reflect.Type) error {
	if t.Kind() != reflect.Struct || !t.ConvertibleTo(typeTimes[0]) {
		return fmt.Errorf("the underlying type of registration must be time")
	}

	typeTimes = append(typeTimes, t)
	return nil
}

// GetRealTime 如果是时间类型，则强制转化为 time.Time 返回。
func GetRealTime(vv reflect.Value) (time.Time, bool) {
	v := reflect.Indirect(vv)
	if IsTime(v.Type()) {
		switch iv := v.Interface().(type) {
		case time.Time:
			return iv, true
		case Time:
			return time.Time(iv), true
		default: // 强制转化为 time.Time 类型
			return v.Convert(typeTimes[0]).Interface().(time.Time), true
		}
	}
	return time.Time{}, false
}

// IsTime 传入类型是否时间
func IsTime(v reflect.Type) bool {
	for _, typeTime := range typeTimes {
		if v == typeTime {
			return true
		}
	}
	return false
}
