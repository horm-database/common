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
package util

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/horm-database/common/consts"
	"github.com/horm-database/common/types"
)

// FormatData 根据类型格式化数据
func FormatData(attributes map[string]interface{}, typeMap map[string]consts.DataType) (map[string]interface{}, error) {
	if len(attributes) == 0 || len(typeMap) == 0 {
		return attributes, nil
	}

	for key, value := range attributes {
		tmp, err := GetDataByType(key, value, typeMap)
		if err != nil {
			return nil, err
		}
		attributes[key] = tmp
	}

	return attributes, nil
}

// FormatDatas 根据类型格式化数据
func FormatDatas(datas []map[string]interface{}, typeMap map[string]consts.DataType) ([]map[string]interface{}, error) {
	if len(datas) == 0 || len(typeMap) == 0 {
		return datas, nil
	}

	for i, data := range datas {
		for key, value := range data {
			tmp, err := GetDataByType(key, value, typeMap)
			if err != nil {
				return nil, err
			}
			datas[i][key] = tmp
		}
	}

	return datas, nil
}

// GetDataByType 获取数据的真实类型
func GetDataByType(key string, value interface{}, typeMap map[string]consts.DataType) (interface{}, error) {
	if len(typeMap) == 0 {
		return value, nil
	}

	typ, ok := typeMap[key]
	if ok {
		var err error

		// 如果是数组，则递归获取每个数据的真实类型
		if arrVal, isArr := value.([]interface{}); isArr {
			for k, v := range arrVal {
				arrVal[k], err = GetDataByType(key, v, typeMap)
				if err != nil {
					return nil, err
				}
			}
			return arrVal, nil
		}

		// 如果是 map，则直接返回，不做类型转换，类型转换主要针对的 clickhouse ，clickhouse 一般不存在 map 类型
		if _, isMap := value.(map[string]interface{}); isMap {
			return value, nil
		}

		return getValueByType(typ, value)
	} else {
		v, is := value.(json.Number)
		if is { // float64 不会被标记
			return v.Float64()
		}
		return value, nil
	}
}

// 获取数据的真实类型
func getValueByType(typ consts.DataType, value interface{}) (interface{}, error) {
	switch typ {
	case consts.DataTypeBytes:
		switch realValue := value.(type) {
		case string:
			bt := types.StringToBytes(realValue)
			ret := make([]byte, base64.StdEncoding.DecodedLen(len(bt)))
			n, err := base64.StdEncoding.Decode(ret, bt)
			if err != nil {
				return nil, err
			}
			return ret[:n], nil
		default:
			return value, nil
		}
	case consts.DataTypeTime:
		switch realValue := value.(type) {
		case string:
			t := time.Time{}
			if realValue == "" {
				return t, nil
			}

			err := t.UnmarshalJSON(types.StringToBytes(`"` + realValue + `"`))
			return t, err
		default:
			return value, nil
		}
	case consts.DataTypeInt:
		return types.InterfaceToInt(value)
	case consts.DataTypeInt8:
		return types.InterfaceToInt8(value)
	case consts.DataTypeInt16:
		return types.InterfaceToInt16(value)
	case consts.DataTypeInt32:
		return types.InterfaceToInt32(value)
	case consts.DataTypeInt64:
		return types.InterfaceToInt64(value)
	case consts.DataTypeUint:
		return types.InterfaceToUint(value)
	case consts.DataTypeUint8:
		return types.InterfaceToUint8(value)
	case consts.DataTypeUint16:
		return types.InterfaceToUint16(value)
	case consts.DataTypeUint32:
		return types.InterfaceToUint32(value)
	case consts.DataTypeUint64:
		return types.InterfaceToUint64(value)
	}

	switch realValue := value.(type) {
	case json.Number:
		return realValue.Float64()
	default:
		return value, nil
	}
}
