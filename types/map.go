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
		return "", true
	}

	return ToString(value), true
}

func (m Map) GetBytes(key string) (ret []byte, exist bool) {
	if len(m) == 0 {
		return nil, false
	}

	value, ok := m[key]
	if !ok {
		return nil, false
	}

	return ToBytes(value), true
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
		return false, true
	}

	return ToBool(value), true
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
		return 0, true, nil
	}

	tmp, err := ToInt64(value)
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
		return 0, true, nil
	}

	tmp, err := ToUint64(value)
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
		return 0, true, nil
	}

	ret, err = ToFloat64(value)
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

func (m Map) GetTime(key string, loc *time.Location, layout ...string) (ret time.Time, exist bool, err error) {
	if len(m) == 0 {
		return time.Time{}, false, nil
	}

	value, ok := m[key]
	if !ok {
		return time.Time{}, false, nil
	}

	if value == nil {
		return time.Time{}, true, nil
	}

	ret, err = ParseTime(value, loc, layout...)
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

	ret, err = ToMap(value, "")
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
		return nil, true, nil
	}

	ret, err = ToStringArray(value)
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
		return nil, true, nil
	}

	ret, err = ToInt64Array(value)
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
		return nil, true, nil
	}

	ret, err = ToUint64Array(value)
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
		return nil, true, nil
	}

	ret, err = ToFloat64Array(value)
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

	mapArr, err := ToMapArray(value, "")
	return mapArr, true, err
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

func GetTime(v map[string]interface{}, key string,
	loc *time.Location, layout ...string) (ret time.Time, exist bool, err error) {
	return Map(v).GetTime(key, loc, layout...)
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
