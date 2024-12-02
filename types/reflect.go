package types

import (
	"reflect"
	"time"
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

// IsEmptyValue 是否空值
func IsEmptyValue(v reflect.Value) bool {
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
			return IsEmptyValue(v.Elem())
		}
	default:
		if v.Type().String() == "time.Time" {
			if !v.CanInterface() {
				return true
			}

			t, ok := v.Interface().(time.Time)
			if ok && !t.IsZero() {
				return false
			}
			return true
		} else if vk == reflect.Struct && !v.IsZero() {
			return true
		}
	}

	return false
}