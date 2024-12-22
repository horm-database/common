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

package json

import (
	"encoding/json"
	"fmt"

	"github.com/horm-database/common/types"

	"github.com/json-iterator/go"
)

const (
	EncodeTypeNormal = 0
	EncodeTypeFast   = 1
	EncodeTypeSort   = 2
)

var Api = jsoniter.Config{
	EscapeHTML:             true,
	ValidateJsonRawMessage: true,
	UseNumber:              true,
}.Froze()

var EscapeApi = jsoniter.Config{
	EscapeHTML: true,
}.Froze()

var FastApi = jsoniter.Config{}.Froze()

var SortApi = jsoniter.Config{
	EscapeHTML:             true,
	SortMapKeys:            true,
	ValidateJsonRawMessage: true,
}.Froze()

// MarshalToString json encode，return string
func MarshalToString(data interface{}, encodeType ...int8) string {
	return types.BytesToString(Marshal(data, encodeType...))
}

// Marshal json encode，return []byte
func Marshal(data interface{}, encodeType ...int8) []byte {
	var result []byte

	if len(encodeType) > 0 {
		switch encodeType[0] {
		case EncodeTypeFast:
			result, _ = FastApi.Marshal(data)
		case EncodeTypeSort:
			result, _ = SortApi.Marshal(data)
		default:
			result, _ = Api.Marshal(data)
		}
	} else {
		result, _ = Api.Marshal(data)
	}

	return result
}

// MarshalBase marshal data other than base structure
func MarshalBase(value interface{}, encodeType ...int8) ([]byte, error) {
	if value == nil {
		return types.StringToBytes(""), nil
	}

	switch v := value.(type) {
	case string:
		return types.StringToBytes(v), nil
	case *string:
		return types.StringToBytes(*v), nil
	case []byte:
		return v, nil
	case *[]byte:
		return *v, nil
	case bool:
		return types.StringToBytes(fmt.Sprintf("%v", v)), nil
	case *bool:
		return types.StringToBytes(fmt.Sprintf("%v", *v)), nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return types.StringToBytes(fmt.Sprintf("%d", v)), nil
	case *int:
		return types.StringToBytes(fmt.Sprintf("%d", *v)), nil
	case *int8:
		return types.StringToBytes(fmt.Sprintf("%d", *v)), nil
	case *int16:
		return types.StringToBytes(fmt.Sprintf("%d", *v)), nil
	case *int32:
		return types.StringToBytes(fmt.Sprintf("%d", *v)), nil
	case *int64:
		return types.StringToBytes(fmt.Sprintf("%d", *v)), nil
	case *uint:
		return types.StringToBytes(fmt.Sprintf("%d", *v)), nil
	case *uint8:
		return types.StringToBytes(fmt.Sprintf("%d", *v)), nil
	case *uint16:
		return types.StringToBytes(fmt.Sprintf("%d", *v)), nil
	case *uint32:
		return types.StringToBytes(fmt.Sprintf("%d", *v)), nil
	case *uint64:
		return types.StringToBytes(fmt.Sprintf("%d", *v)), nil
	case float32, float64:
		return types.StringToBytes(fmt.Sprintf("%f", v)), nil
	case *float32:
		return types.StringToBytes(fmt.Sprintf("%f", *v)), nil
	case *float64:
		return types.StringToBytes(fmt.Sprintf("%f", *v)), nil
	case json.Number:
		return types.StringToBytes(v.String()), nil
	case *json.Number:
		return types.StringToBytes(v.String()), nil
	case *interface{}:
		return MarshalBase(*v, encodeType...)
	default:
		var result []byte
		var err error

		if len(encodeType) > 0 {
			switch encodeType[0] {
			case EncodeTypeFast:
				result, err = FastApi.Marshal(v)
			case EncodeTypeSort:
				result, err = SortApi.Marshal(v)
			default:
				result, err = Api.Marshal(v)
			}
		} else {
			result, err = Api.Marshal(v)
		}

		return result, err
	}
}

// MarshalBaseToString marshal data to string other than base structure
func MarshalBaseToString(value interface{}, encodeType ...int8) (string, error) {
	if value == nil {
		return "", nil
	}

	switch v := value.(type) {
	case string:
		return v, nil
	case *string:
		return *v, nil
	case []byte:
		return types.BytesToString(v), nil
	case *[]byte:
		return types.BytesToString(*v), nil
	case bool:
		return fmt.Sprintf("%v", v), nil
	case *bool:
		return fmt.Sprintf("%v", *v), nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v), nil
	case *int:
		return fmt.Sprintf("%d", *v), nil
	case *int8:
		return fmt.Sprintf("%d", *v), nil
	case *int16:
		return fmt.Sprintf("%d", *v), nil
	case *int32:
		return fmt.Sprintf("%d", *v), nil
	case *int64:
		return fmt.Sprintf("%d", *v), nil
	case *uint:
		return fmt.Sprintf("%d", *v), nil
	case *uint8:
		return fmt.Sprintf("%d", *v), nil
	case *uint16:
		return fmt.Sprintf("%d", *v), nil
	case *uint32:
		return fmt.Sprintf("%d", *v), nil
	case *uint64:
		return fmt.Sprintf("%d", *v), nil
	case float32, float64:
		return fmt.Sprintf("%f", v), nil
	case *float32:
		return fmt.Sprintf("%f", *v), nil
	case *float64:
		return fmt.Sprintf("%f", *v), nil
	case json.Number:
		return v.String(), nil
	case *json.Number:
		return v.String(), nil
	case *interface{}:
		return MarshalBaseToString(*v, encodeType...)
	default:
		var result []byte
		var err error

		if len(encodeType) > 0 {
			switch encodeType[0] {
			case EncodeTypeFast:
				result, err = FastApi.Marshal(v)
			case EncodeTypeSort:
				result, err = SortApi.Marshal(v)
			default:
				result, err = Api.Marshal(v)
			}
		} else {
			result, err = Api.Marshal(v)
		}

		return types.BytesToString(result), err
	}
}
