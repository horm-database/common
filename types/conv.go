package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
	"unsafe"

	"github.com/horm/common/structspec"
)

// InterfaceToString 接口转字符串
func InterfaceToString(value interface{}) string {
	if value == nil {
		return ""
	}

	switch v := value.(type) {
	case string:
		return v
	case *string:
		return *v
	case []byte:
		return string(v)
	case *[]byte:
		return string(*v)
	case bool:
		return fmt.Sprintf("%v", v)
	case *bool:
		return fmt.Sprintf("%v", *v)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v)
	case *int:
		return fmt.Sprintf("%d", *v)
	case *int8:
		return fmt.Sprintf("%d", *v)
	case *int16:
		return fmt.Sprintf("%d", *v)
	case *int32:
		return fmt.Sprintf("%d", *v)
	case *int64:
		return fmt.Sprintf("%d", *v)
	case *uint:
		return fmt.Sprintf("%d", *v)
	case *uint8:
		return fmt.Sprintf("%d", *v)
	case *uint16:
		return fmt.Sprintf("%d", *v)
	case *uint32:
		return fmt.Sprintf("%d", *v)
	case *uint64:
		return fmt.Sprintf("%d", *v)
	case float32, float64:
		return fmt.Sprintf("%f", v)
	case *float32:
		return fmt.Sprintf("%f", *v)
	case *float64:
		return fmt.Sprintf("%f", *v)
	case json.Number:
		return v.String()
	case *json.Number:
		return v.String()
	case *interface{}:
		return InterfaceToString(*v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// InterfaceToBytes 接口转字节码
func InterfaceToBytes(value interface{}) []byte {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case string:
		return []byte(v)
	case *string:
		return []byte(*v)
	case []byte:
		return v
	case *[]byte:
		return *v
	case bool:
		return []byte(fmt.Sprintf("%v", v))
	case *bool:
		return []byte(fmt.Sprintf("%v", *v))
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return []byte(fmt.Sprintf("%d", v))
	case *int:
		return []byte(fmt.Sprintf("%d", v))
	case *int8:
		return []byte(fmt.Sprintf("%d", v))
	case *int16:
		return []byte(fmt.Sprintf("%d", v))
	case *int32:
		return []byte(fmt.Sprintf("%d", v))
	case *int64:
		return []byte(fmt.Sprintf("%d", v))
	case *uint:
		return []byte(fmt.Sprintf("%d", v))
	case *uint8:
		return []byte(fmt.Sprintf("%d", v))
	case *uint16:
		return []byte(fmt.Sprintf("%d", v))
	case *uint32:
		return []byte(fmt.Sprintf("%d", v))
	case *uint64:
		return []byte(fmt.Sprintf("%d", v))
	case float32, float64:
		return []byte(fmt.Sprintf("%f", v))
	case *float32:
		return []byte(fmt.Sprintf("%f", *v))
	case *float64:
		return []byte(fmt.Sprintf("%f", *v))
	case json.Number:
		return []byte(v.String())
	case *json.Number:
		return []byte(v.String())
	case *interface{}:
		return InterfaceToBytes(*v)
	default:
		return []byte(fmt.Sprintf("%v", v))
	}
}

func InterfaceToBool(value interface{}) bool {
	if value == nil {
		return false
	}

	switch v := value.(type) {
	case string:
		b, _ := strconv.ParseBool(v)
		return b
	case *string:
		b, _ := strconv.ParseBool(*v)
		return b
	case []byte:
		b, _ := strconv.ParseBool(BytesToString(v))
		return b
	case *[]byte:
		b, _ := strconv.ParseBool(BytesToString(*v))
		return b
	case int64:
		if v >= 1 {
			return true
		} else {
			return false
		}
	case uint64:
		if v >= 1 {
			return true
		} else {
			return false
		}
	case *int64:
		if *v >= 1 {
			return true
		} else {
			return false
		}
	case *uint64:
		if *v >= 1 {
			return true
		} else {
			return false
		}
	case json.Number:
		iv, _ := v.Int64()
		if iv >= 1 {
			return true
		} else {
			return false
		}
	case *json.Number:
		iv, _ := v.Int64()
		if iv >= 1 {
			return true
		} else {
			return false
		}
	default:
		str := fmt.Sprintf("%v", v)
		b, _ := strconv.ParseBool(str)
		return b
	}
}

func InterfaceToInt64(value interface{}) (int64, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case string:
		return strconv.ParseInt(v, 10, 64)
	case *string:
		return strconv.ParseInt(*v, 10, 64)
	case []byte:
		return strconv.ParseInt(BytesToString(v), 10, 64)
	case *[]byte:
		return strconv.ParseInt(BytesToString(*v), 10, 64)
	case int:
		return int64(v), nil
	case *int:
		return int64(*v), nil
	case int8:
		return int64(v), nil
	case *int8:
		return int64(*v), nil
	case int16:
		return int64(v), nil
	case *int16:
		return int64(*v), nil
	case int32:
		return int64(v), nil
	case *int32:
		return int64(*v), nil
	case uint:
		return int64(v), nil
	case *uint:
		return int64(*v), nil
	case uint8:
		return int64(v), nil
	case *uint8:
		return int64(*v), nil
	case uint16:
		return int64(v), nil
	case *uint16:
		return int64(*v), nil
	case uint32:
		return int64(v), nil
	case *uint32:
		return int64(*v), nil
	case int64:
		return v, nil
	case *int64:
		return *v, nil
	case uint64:
		return int64(v), nil
	case *uint64:
		return int64(*v), nil
	case float64:
		return int64(v), nil
	case *float64:
		return int64(*v), nil
	case float32:
		return int64(v), nil
	case *float32:
		return int64(*v), nil
	case json.Number:
		return v.Int64()
	case *json.Number:
		return v.Int64()
	default:
		return strconv.ParseInt(fmt.Sprintf("%v", v), 10, 64)
	}
}

func InterfaceToUint64(value interface{}) (uint64, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case string:
		return strconv.ParseUint(v, 10, 64)
	case *string:
		return strconv.ParseUint(*v, 10, 64)
	case []byte:
		return strconv.ParseUint(BytesToString(v), 10, 64)
	case *[]byte:
		return strconv.ParseUint(BytesToString(*v), 10, 64)
	case int:
		return uint64(v), nil
	case *int:
		return uint64(*v), nil
	case int8:
		return uint64(v), nil
	case *int8:
		return uint64(*v), nil
	case int16:
		return uint64(v), nil
	case *int16:
		return uint64(*v), nil
	case int32:
		return uint64(v), nil
	case *int32:
		return uint64(*v), nil
	case uint:
		return uint64(v), nil
	case *uint:
		return uint64(*v), nil
	case uint8:
		return uint64(v), nil
	case *uint8:
		return uint64(*v), nil
	case uint16:
		return uint64(v), nil
	case *uint16:
		return uint64(*v), nil
	case uint32:
		return uint64(v), nil
	case *uint32:
		return uint64(*v), nil
	case int64:
		return uint64(v), nil
	case *int64:
		return uint64(*v), nil
	case uint64:
		return v, nil
	case *uint64:
		return *v, nil
	case float64:
		return uint64(v), nil
	case *float64:
		return uint64(*v), nil
	case float32:
		return uint64(v), nil
	case *float32:
		return uint64(*v), nil
	case json.Number:
		return strconv.ParseUint(v.String(), 10, 64)
	case *json.Number:
		return strconv.ParseUint(v.String(), 10, 64)
	default:
		return strconv.ParseUint(fmt.Sprintf("%v", v), 10, 64)
	}
}

func InterfaceToFloat64(value interface{}) (float64, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case string:
		return strconv.ParseFloat(v, 64)
	case *string:
		return strconv.ParseFloat(*v, 64)
	case []byte:
		return strconv.ParseFloat(BytesToString(v), 64)
	case *[]byte:
		return strconv.ParseFloat(BytesToString(*v), 64)
	case int:
		return float64(v), nil
	case *int:
		return float64(*v), nil
	case int8:
		return float64(v), nil
	case *int8:
		return float64(*v), nil
	case int16:
		return float64(v), nil
	case *int16:
		return float64(*v), nil
	case int32:
		return float64(v), nil
	case *int32:
		return float64(*v), nil
	case uint:
		return float64(v), nil
	case *uint:
		return float64(*v), nil
	case uint8:
		return float64(v), nil
	case *uint8:
		return float64(*v), nil
	case uint16:
		return float64(v), nil
	case *uint16:
		return float64(*v), nil
	case uint32:
		return float64(v), nil
	case *uint32:
		return float64(*v), nil
	case int64:
		return float64(v), nil
	case *int64:
		return float64(*v), nil
	case uint64:
		return float64(v), nil
	case *uint64:
		return float64(*v), nil
	case float64:
		return v, nil
	case *float64:
		return *v, nil
	case float32:
		return float64(v), nil
	case *float32:
		return float64(*v), nil
	case json.Number:
		return v.Float64()
	case *json.Number:
		return v.Float64()
	default:
		return strconv.ParseFloat(fmt.Sprintf("%v", v), 64)
	}
}

func InterfaceToInt(value interface{}) (int, error) {
	tmp, err := InterfaceToInt64(value)
	return int(tmp), err
}

func InterfaceToInt8(value interface{}) (int8, error) {
	tmp, err := InterfaceToInt64(value)
	return int8(tmp), err
}

func InterfaceToInt16(value interface{}) (int16, error) {
	tmp, err := InterfaceToInt64(value)
	return int16(tmp), err
}

func InterfaceToInt32(value interface{}) (int32, error) {
	tmp, err := InterfaceToInt64(value)
	return int32(tmp), err
}

func InterfaceToUint(value interface{}) (uint, error) {
	tmp, err := InterfaceToInt64(value)
	return uint(tmp), err
}

func InterfaceToUint8(value interface{}) (uint8, error) {
	tmp, err := InterfaceToInt64(value)
	return uint8(tmp), err
}

func InterfaceToUint16(value interface{}) (uint16, error) {
	tmp, err := InterfaceToInt64(value)
	return uint16(tmp), err
}

func InterfaceToUint32(value interface{}) (uint32, error) {
	tmp, err := InterfaceToInt64(value)
	return uint32(tmp), err
}

func InterfaceToTime(value interface{}, layout string, loc ...*time.Location) (time.Time, error) {
	if value == nil {
		return time.Time{}, nil
	}

	switch val := value.(type) {
	case time.Time:
		return val, nil
	case *time.Time:
		if val == nil {
			return time.Time{}, nil
		} else {
			return *val, nil
		}
	default:
		str := InterfaceToString(value)

		l := time.Local
		if len(loc) > 0 {
			l = loc[0]
		}

		return time.ParseInLocation(layout, str, l)
	}
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
		ss := structspec.GetStructSpec("", rv.Type())
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
	}

	return map[string]interface{}{}, errors.New("value is not map")
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
