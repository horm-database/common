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
package consts

import (
	"time"
)

type DataType int8

const ( /*  默认为 other，不需要特殊处理，例如 string、float，下面类型需要特殊处理转换，
	比如 int64 整数，json 编码上传到服务端之后，解码会转化为 float64，可能存在精度丢失问题，
	clickhouse 对类型也有非常强的限制。*/
	DataTypeOther  DataType = 0  // 其他类型
	DataTypeBytes  DataType = 1  // 类型是 []byte
	DataTypeTime   DataType = 2  // 类型是 time.Time
	DataTypeInt    DataType = 3  // 类型是 int
	DataTypeInt8   DataType = 4  // 类型是 int8
	DataTypeInt16  DataType = 5  // 类型是 int16
	DataTypeInt32  DataType = 6  // 类型是 int32
	DataTypeInt64  DataType = 7  // 类型是 int64
	DataTypeUint   DataType = 8  // 类型是 uint
	DataTypeUint8  DataType = 9  // 类型是 uint8
	DataTypeUint16 DataType = 10 // 类型是 uint16
	DataTypeUint32 DataType = 11 // 类型是 uint32
	DataTypeUint64 DataType = 12 // 类型是 uint64
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

func GetDataType(v interface{}) DataType {
	switch v.(type) {
	case []byte, *[]byte:
		return DataTypeBytes
	case int, []int, *int, *[]int:
		return DataTypeInt
	case int8, []int8, *int8, *[]int8:
		return DataTypeInt8
	case int16, []int16, *int16, *[]int16:
		return DataTypeInt16
	case int32, []int32, *int32, *[]int32:
		return DataTypeInt32
	case int64, []int64, *int64, *[]int64:
		return DataTypeInt64
	case uint, []uint, *uint, *[]uint:
		return DataTypeUint
	case uint8, *uint8:
		return DataTypeUint8
	case uint16, []uint16, *uint16, *[]uint16:
		return DataTypeUint16
	case uint32, []uint32, *uint32, *[]uint32:
		return DataTypeUint32
	case uint64, []uint64, *uint64, *[]uint64:
		return DataTypeUint64
	case time.Time, []time.Time, *time.Time, *[]time.Time:
		return DataTypeTime
	default:
		return DataTypeOther
	}
}
