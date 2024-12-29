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
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/araddon/dateparse"
)

type Time time.Time

// UnmarshalJSON NullTime 类型实现 json marshal 方法
func (nt *Time) UnmarshalJSON(data []byte) error {
	tStr := strings.TrimPrefix(string(data), `"`)
	tStr = strings.TrimSuffix(tStr, `"`)

	if tStr == "null" {
		return nil
	}

	t, err := ParseTime(tStr, "", time.Local)
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
func GetRealTime(data interface{}) (time.Time, bool) {
	switch t := data.(type) {
	case time.Time:
		return t, true
	case Time:
		return time.Time(t), true
	case *time.Time:
		return *t, true
	case *Time:
		return time.Time(*t), true
	}

	v := reflect.Indirect(reflect.ValueOf(data))
	if IsTime(v.Type()) {
		// 强制转化为 time.Time 类型
		return v.Convert(typeTimes[0]).Interface().(time.Time), true
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

var dateTimeLen = len(time.DateTime)
var dateOnlyLen = len(time.DateOnly)

// ParseTime 解析任何时间
func ParseTime(src interface{}, layout string, loc *time.Location) (time.Time, error) {
	if loc == nil {
		loc = time.Local
	}

	switch v := src.(type) {
	case []byte:
		return ParseTime(BytesToString(v), layout, loc)
	case *[]byte:
		return ParseTime(BytesToString(*v), layout, loc)
	case string:
		l := len(v)
		if l == 0 {
			return time.Time{}, nil
		}

		if layout != "" {
			return time.ParseInLocation(layout, v, loc)
		}

		i, err := strconv.ParseUint(v, 10, 64)
		if err == nil && i > 631123200 { //时间戳，且大于 1990-01-01 00:00:00
			return ParseTime(i, layout, loc)
		}

		switch l {
		case dateTimeLen:
			t, e := time.ParseInLocation(time.DateTime, v, loc) // 首先校验是否最常用的 DateTime 格式。
			if e == nil {
				return t, nil
			}
		case dateOnlyLen:
			t, e := time.ParseInLocation(time.DateOnly, v, loc) // 首先校验是否最常用的 DateOnly 格式，日期一般这个格式。
			if e == nil {
				return t, nil
			}
		default:
			t, e := time.ParseInLocation(time.RFC3339Nano, v, loc) // 首先校验是否最常用的 RFC3339Nano 格式，json.Marshal 一般是这个格式。
			if e == nil {
				return t, nil
			}
		}

		return dateparse.ParseIn(v, loc)
	case json.Number:
		s, e := v.Int64()
		if e != nil {
			return time.Time{}, fmt.Errorf("unable to cast %#v of type %T to Time", src, src)
		}

		if s > 17356608000000 { // 微妙
			return time.Unix(s/1e6, (s%1e6)*1e3), nil
		} else if s > 17356608000 { // 毫秒
			return time.Unix(s/1e3, (s%1e3)*1e6), nil
		} else {
			return time.Unix(s, 0), nil
		}
	case int:
		if v > 17356608000 { // 毫秒
			return time.Unix(int64(v/1e3), int64((v%1e3)*1e6)), nil
		} else {
			return time.Unix(int64(v), 0), nil
		}
	case int64:
		if v > 17356608000000 { // 微妙
			return time.Unix(v/1e6, (v%1e6)*1e3), nil
		} else if v > 17356608000 { // 毫秒
			return time.Unix(v/1e3, (v%1e3)*1e6), nil
		} else {
			return time.Unix(v, 0), nil
		}
	case int32:
		return time.Unix(int64(v), 0), nil
	case uint:
		if v > 17356608000 { // 毫秒
			return time.Unix(int64(v/1e3), int64((v%1e3)*1e6)), nil
		} else {
			return time.Unix(int64(v), 0), nil
		}
	case uint64:
		if v > 17356608000000 { // 微妙
			return time.Unix(int64(v/1e6), int64((v%1e6)*1e3)), nil
		} else if v > 17356608000 { // 毫秒
			return time.Unix(int64(v/1e3), int64((v%1e3)*1e6)), nil
		} else {
			return time.Unix(int64(v), 0), nil
		}
	case uint32:
		return time.Unix(int64(v), 0), nil
	default:
		t, ok := GetRealTime(src)
		if ok {
			return t, nil
		}

		return time.Time{}, fmt.Errorf("unable to cast %#v of type %T to Time", src, src)
	}
}
