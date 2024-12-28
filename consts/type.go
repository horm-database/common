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

package consts

import (
	"reflect"

	"github.com/horm-database/common/structs"
	"github.com/horm-database/common/types"
)

type RetType int8

const (
	RedisRetTypeNil         RetType = 1 // 无返回
	RedisRetTypeString      RetType = 2 // 字符串
	RedisRetTypeBool        RetType = 3 // bool
	RedisRetTypeInt64       RetType = 4 // int64
	RedisRetTypeFloat64     RetType = 5 // float64
	RedisRetTypeStrings     RetType = 6 // 字符串数组
	RedisRetTypeMapString   RetType = 7 // map[string]string
	RedisRetTypeMemberScore RetType = 8 // 有序成员
)

func GetRedisRetType(op string, withScore bool) RetType {
	switch op {
	case OpGet, OpGetSet, OpHGet, OpLPop, OpRPop:
		return RedisRetTypeString
	case OpSet, OpExists, OpSetNX, OpSetBit, OpGetBit, OpHSetNx, OpHExists, OpSIsMember:
		return RedisRetTypeBool
	case OpTTL, OpDel, OpIncr, OpDecr, OpIncrBy, OpBitCount, OpHIncrBy, OpHDel, OpHLen,
		OpHStrLen, OpLPush, OpRPush, OpLLen, OpSAdd, OpSRem, OpSCard, OpSMove, OpZAdd,
		OpZRem, OpZRemRangeByScore, OpZRemRangeByRank, OpZCard, OpZRank, OpZRevRank, OpZCount:
		return RedisRetTypeInt64
	case OpHIncrByFloat, OpZIncrBy, OpZScore:
		return RedisRetTypeFloat64
	case OpMGet, OpHKeys, OpHVals, OpSMembers, OpSRandMember, OpSPop:
		return RedisRetTypeStrings
	case OpHGetAll, OpHMGet:
		return RedisRetTypeMapString
	case OpZRange, OpZRangeByScore, OpZRevRange, OpZRevRangeByScore:
		if withScore {
			return RedisRetTypeMemberScore
		} else {
			return RedisRetTypeStrings
		}
	case OpZPopMin, OpZPopMax:
		return RedisRetTypeMemberScore
	}

	return RedisRetTypeNil
}

func GetDataType(v interface{}) structs.Type {
	switch v.(type) {
	case []byte, *[]byte:
		return structs.TypeBytes
	case int, []int, *int, *[]int:
		return structs.TypeInt
	case int8, []int8, *int8, *[]int8:
		return structs.TypeInt8
	case int16, []int16, *int16, *[]int16:
		return structs.TypeInt16
	case int32, []int32, *int32, *[]int32:
		return structs.TypeInt32
	case int64, []int64, *int64, *[]int64:
		return structs.TypeInt64
	case uint, []uint, *uint, *[]uint:
		return structs.TypeUint
	case uint8, *uint8:
		return structs.TypeUint8
	case uint16, []uint16, *uint16, *[]uint16:
		return structs.TypeUint16
	case uint32, []uint32, *uint32, *[]uint32:
		return structs.TypeUint32
	case uint64, []uint64, *uint64, *[]uint64:
		return structs.TypeUint64
	default:
		if types.IsTime(reflect.TypeOf(v)) {
			return structs.TypeTime
		}
		return 0
	}
}
