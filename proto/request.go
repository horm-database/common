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

package proto

import (
	"github.com/horm-database/common/proto/sql"
	"github.com/horm-database/common/types"
)

// Unit 查询单元（执行单元）
type Unit struct {
	// query base info
	Name  string   `json:"name,omitempty"`  // name
	Op    string   `json:"op,omitempty"`    // operation
	Shard []string `json:"shard,omitempty"` // 分片、分表、分库

	// 结构化查询共有
	Column []string               `json:"column,omitempty"` // columns
	Where  map[string]interface{} `json:"where,omitempty"`  // query condition
	Order  []string               `json:"order,omitempty"`  // order by
	Page   int                    `json:"page,omitempty"`   // request pages. when page > 0, the request is returned in pagination.
	Size   int                    `json:"size,omitempty"`   // size per page
	From   uint64                 `json:"from,omitempty"`   // offset

	// data maintain
	Val      interface{}              `json:"val,omitempty"`       // 单条记录 val (not map/[]map)
	Data     map[string]interface{}   `json:"data,omitempty"`      // add/update one map data
	Datas    []map[string]interface{} `json:"datas,omitempty"`     // batch add/update map data
	Args     []interface{}            `json:"args,omitempty"`      // multiple args, 还可用于 query 语句的参数，或者 redis 协议，如 MGET、HMGET、HDEL 等
	DataType map[string]types.Type    `json:"data_type,omitempty"` // 数据类型（主要用于 clickhouse，对于数据类型有强依赖），请求 json 不区分 int8、int16、int32、int64 等，只有 Number 类型，bytes 也会被当成 string 处理。

	// group by
	Group  []string               `json:"group,omitempty"`  // group by
	Having map[string]interface{} `json:"having,omitempty"` // group by condition

	// for databases such as mysql ...
	Join []*sql.Join `json:"join,omitempty"`

	// for databases such as elastic ...
	Type   string  `json:"type,omitempty"`   // type, such as elastic`s type, it can be customized before v7, and unified as _doc after v7
	Scroll *Scroll `json:"scroll,omitempty"` // scroll info

	// for databases such as redis ...
	Prefix string `json:"prefix,omitempty"` // prefix, It is strongly recommended to bring it to facilitate finer-grained summary statistics, otherwise the statistical granularity can only be cmd ，such as GET、SET、HGET ...
	Key    string `json:"key,omitempty"`    // key
	Field  string `json:"field,omitempty"`  // redis hash field

	// bytes 字节流
	Bytes []byte `json:"bytes,omitempty"`

	// params 与数据库特性相关的附加参数，例如 redis 的 withscores、EX、NX、等，以及 elastic 的 refresh、collapse、runtime_mappings、track_total_hits 等等。
	Params map[string]interface{} `json:"params,omitempty"`

	// 直接送 Query 语句，需要拥有库的 表权限、或 root 权限。具体参数为 args
	Query string `json:"query,omitempty"`

	// Extend 扩展信息，作用于插件
	Extend map[string]interface{} `json:"extend,omitempty"`

	Sub   []*Unit `json:"sub,omitempty"`   // 子查询
	Trans []*Unit `json:"trans,omitempty"` // 事务，该事务下的所有 Unit 必须同时成功或失败（注意：仅适合支持事务的数据库回滚，如果数据库不支持事务，则操作不会回滚）
}

// Scroll 滚动查询
type Scroll struct {
	ID   string `json:"id,omitempty"`   // 滚动 id
	Info string `json:"info,omitempty"` // 滚动查询信息，如时间
}
