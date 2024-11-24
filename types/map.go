package types

import (
	"errors"
	"reflect"
	"time"
)

type Map map[string]interface{}

func (m Map) Set(key string, value interface{}) {
	m[key] = value
}

func (m Map) GetString(key string) (ret string, exist bool) {
	if len(m) == 0 {
		return "", false
	}

	value, ok := m[key]
	if !ok {
		return "", false
	}

	if value == nil {
		return "", false
	}

	return InterfaceToString(value), true
}

func (m Map) GetBytes(key string) (ret []byte, exist bool) {
	if len(m) == 0 {
		return nil, false
	}

	value, ok := m[key]
	if !ok {
		return nil, false
	}

	return InterfaceToBytes(value), true
}

func (m Map) GetBool(key string) (ret bool, exist bool) {
	if len(m) == 0 {
		return false, false
	}

	value, ok := m[key]
	if !ok {
		return false, false
	}

	if value == nil {
		return false, false
	}

	return InterfaceToBool(value), true
}

func (m Map) GetInt64(key string) (ret int64, exist bool, err error) {
	if len(m) == 0 {
		return 0, false, nil
	}

	value, ok := m[key]
	if !ok {
		return 0, false, nil
	}

	if value == nil {
		return 0, false, nil
	}

	tmp, err := InterfaceToInt64(value)
	return tmp, true, err
}

func (m Map) GetUint64(key string) (ret uint64, exist bool, err error) {
	if len(m) == 0 {
		return 0, false, nil
	}

	value, ok := m[key]
	if !ok {
		return 0, false, nil
	}

	if value == nil {
		return 0, false, nil
	}

	tmp, err := InterfaceToUint64(value)
	return tmp, true, err
}

func (m Map) GetFloat64(key string) (ret float64, exist bool, err error) {
	if len(m) == 0 {
		return 0, false, nil
	}

	value, ok := m[key]
	if !ok {
		return 0, false, nil
	}

	if value == nil {
		return 0, false, nil
	}

	ret, err = InterfaceToFloat64(value)
	return ret, true, err
}

func (m Map) GetInt(key string) (int, bool, error) {
	i64, exist, err := m.GetInt64(key)
	return int(i64), exist, err
}

func (m Map) GetInt8(key string) (int8, bool, error) {
	i64, exist, err := m.GetInt64(key)
	return int8(i64), exist, err
}

func (m Map) GetInt16(key string) (int16, bool, error) {
	i64, exist, err := m.GetInt64(key)
	return int16(i64), exist, err
}

func (m Map) GetInt32(key string) (int32, bool, error) {
	i64, exist, err := m.GetInt64(key)
	return int32(i64), exist, err
}

func (m Map) GetUint(key string) (uint, bool, error) {
	ui64, exist, err := m.GetUint64(key)
	return uint(ui64), exist, err
}

func (m Map) GetUint8(key string) (uint8, bool, error) {
	ui64, exist, err := m.GetUint64(key)
	return uint8(ui64), exist, err
}

func (m Map) GetUint16(key string) (uint16, bool, error) {
	ui64, exist, err := m.GetUint64(key)
	return uint16(ui64), exist, err
}

func (m Map) GetUint32(key string) (uint32, bool, error) {
	ui64, exist, err := m.GetUint64(key)
	return uint32(ui64), exist, err
}

func (m Map) GetTime(key, layout string, loc ...*time.Location) (ret time.Time, exist bool, err error) {
	if len(m) == 0 {
		return time.Time{}, false, nil
	}

	value, ok := m[key]
	if !ok {
		return time.Time{}, false, nil
	}

	if value == nil {
		return time.Time{}, false, nil
	}

	ret, err = InterfaceToTime(value, layout, loc...)
	return ret, true, err
}

func (m Map) GetMap(key string) (ret Map, exist bool, err error) {
	if len(m) == 0 {
		return nil, false, nil
	}

	value, ok := m[key]
	if !ok {
		return nil, false, nil
	}

	if value == nil {
		return nil, true, nil
	}

	ret, err = InterfaceToMap(value)
	return ret, true, err
}

func (m Map) GetStringArray(key string) (ret []string, exist bool, err error) {
	if len(m) == 0 {
		return nil, false, nil
	}

	value, ok := m[key]
	if !ok {
		return nil, false, nil
	}

	if value == nil {
		return nil, false, nil
	}

	ret, err = InterfaceToStringArray(value)
	return ret, true, err
}

func (m Map) GetInt64Array(key string) (ret []int64, exist bool, err error) {
	if len(m) == 0 {
		return nil, false, nil
	}

	value, ok := m[key]
	if !ok {
		return nil, false, nil
	}

	if value == nil {
		return nil, false, nil
	}

	ret, err = InterfaceToInt64Array(value)
	return ret, true, err
}

func (m Map) GetUint64Array(key string) (ret []uint64, exist bool, err error) {
	if len(m) == 0 {
		return nil, false, nil
	}

	value, ok := m[key]
	if !ok {
		return nil, false, nil
	}

	if value == nil {
		return nil, false, nil
	}

	ret, err = InterfaceToUint64Array(value)
	return ret, true, err
}

func (m Map) GetFloat64Array(key string) (ret []float64, exist bool, err error) {
	if len(m) == 0 {
		return nil, false, nil
	}

	value, ok := m[key]
	if !ok {
		return nil, false, nil
	}

	if value == nil {
		return nil, false, nil
	}

	ret, err = InterfaceToFloat64Array(value)
	return ret, true, err
}

func (m Map) GetMapArray(key string) (ret []Map, exist bool, err error) {
	value, ok := m[key]
	if !ok {
		return nil, false, nil
	}

	if value == nil {
		return nil, true, nil
	}

	switch arrVal := value.(type) {
	case []Map:
		return arrVal, true, nil
	case []interface{}:
		ret = make([]Map, len(arrVal))
		for k, arrItem := range arrVal {
			im, e := InterfaceToMap(arrItem)
			if e != nil {
				return nil, true, e
			}
			ret[k] = im
		}
	case []map[string]interface{}:
		ret = make([]Map, len(arrVal))
		for k, arrItem := range arrVal {
			ret[k] = arrItem
		}
	default:
		v := reflect.ValueOf(value)
		if IsNil(v) {
			return nil, true, nil
		}

		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		if !IsArray(v) {
			return nil, true, errors.New("value is not array")
		}

		l := v.Len()
		ret = make([]Map, l)

		for i := 0; i < l; i++ {
			im, e := InterfaceToMap(Interface(v.Index(i)))
			if e != nil {
				return nil, true, e
			}
			ret[i] = im
		}
	}

	return ret, true, nil
}

func GetString(v map[string]interface{}, key string) (ret string, exist bool) {
	return Map(v).GetString(key)
}

func GetBytes(v map[string]interface{}, key string) (ret []byte, exist bool) {
	return Map(v).GetBytes(key)
}

func GetBool(v map[string]interface{}, key string) (ret bool, exist bool) {
	return Map(v).GetBool(key)
}

func GetInt64(v map[string]interface{}, key string) (ret int64, exist bool, err error) {
	return Map(v).GetInt64(key)
}

func GetUint64(v map[string]interface{}, key string) (ret uint64, exist bool, err error) {
	return Map(v).GetUint64(key)
}

func GetFloat64(v map[string]interface{}, key string) (ret float64, exist bool, err error) {
	return Map(v).GetFloat64(key)
}

func GetInt(v map[string]interface{}, key string) (ret int, exist bool, err error) {
	return Map(v).GetInt(key)
}

func GetInt8(v map[string]interface{}, key string) (ret int8, exist bool, err error) {
	return Map(v).GetInt8(key)
}

func GetInt16(v map[string]interface{}, key string) (ret int16, exist bool, err error) {
	return Map(v).GetInt16(key)
}

func GetInt32(v map[string]interface{}, key string) (ret int32, exist bool, err error) {
	return Map(v).GetInt32(key)
}

func GetUint(v map[string]interface{}, key string) (ret uint, exist bool, err error) {
	return Map(v).GetUint(key)
}

func GetUint8(v map[string]interface{}, key string) (ret uint8, exist bool, err error) {
	return Map(v).GetUint8(key)
}

func GetUint16(v map[string]interface{}, key string) (ret uint16, exist bool, err error) {
	return Map(v).GetUint16(key)
}

func GetUint32(v map[string]interface{}, key string) (ret uint32, exist bool, err error) {
	return Map(v).GetUint32(key)
}

func GetTime(v map[string]interface{}, key, layout string,
	loc ...*time.Location) (ret time.Time, exist bool, err error) {
	return Map(v).GetTime(key, layout, loc...)
}

func GetMap(v map[string]interface{}, key string) (Map, bool, error) {
	return Map(v).GetMap(key)
}

func GetStringArray(v map[string]interface{}, key string) ([]string, bool, error) {
	return Map(v).GetStringArray(key)
}

func GetInt64Array(v map[string]interface{}, key string) ([]int64, bool, error) {
	return Map(v).GetInt64Array(key)
}

func GetUint64Array(v map[string]interface{}, key string) ([]uint64, bool, error) {
	return Map(v).GetUint64Array(key)
}

func GetFloat64Array(v map[string]interface{}, key string) ([]float64, bool, error) {
	return Map(v).GetFloat64Array(key)
}

func GetMapArray(v map[string]interface{}, key string) ([]Map, bool, error) {
	return Map(v).GetMapArray(key)
}
