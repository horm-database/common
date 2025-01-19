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
	"encoding/json"
	"go/types"
	"reflect"
	"time"
)

// IsArray 判断是否 Array 或 Slice
func IsArray(v reflect.Value) bool {
	k := v.Kind()
	if k == reflect.Interface {
		return IsArray(v.Elem())
	}

	return k == reflect.Slice || k == reflect.Array
}

func IsMap(v reflect.Value) bool {
	k := v.Kind()
	if k == reflect.Interface {
		return IsMap(v.Elem())
	}

	return k == reflect.Map
}

func IsStruct(typ reflect.Type) bool {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	return typ.Kind() == reflect.Struct
}

func IsStructArray(v reflect.Value) bool {
	if !IsArray(v) {
		return false
	}

	return IsStruct(v.Type().Elem())
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

	if !v.CanInterface() {
		return nil
	}

	return v.Interface()
}

// Indirect 获取指针的值
func Indirect(data interface{}) interface{} {
	if data == nil {
		return nil
	}

	switch v := data.(type) {
	case *string:
		return *v
	case *[]byte:
		return *v
	case *bool:
		return *v
	case *int:
		return *v
	case *int8:
		return *v
	case *int16:
		return *v
	case *int32:
		return *v
	case *int64:
		return *v
	case *uint:
		return *v
	case *uint8:
		return *v
	case *uint16:
		return *v
	case *uint32:
		return *v
	case *uint64:
		return *v
	case *float32:
		return *v
	case *float64:
		return *v
	case *json.Number:
		return *v
	case *types.Map:
		return *v
	case *[]types.Map:
		return *v
	case *map[string]interface{}:
		return *v
	case *[]map[string]interface{}:
		return *v
	case *time.Time:
		return *v
	case *interface{}:
		return Indirect(*v)
	}

	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Pointer {
		return v.Elem().Interface()
	}

	return data
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
		t, isTime := GetRealTime(v.Interface())
		if isTime {
			return t.IsZero()
		}
	}

	return v.IsZero()
}
