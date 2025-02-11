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

package plugin

import (
	"github.com/horm-database/common/proto"
	"github.com/horm-database/common/proto/sql"
	"github.com/horm-database/common/types"
	"github.com/horm-database/common/util"
)

// Header 请求头部，通过 extend["request_header"] 传入插件
type Header struct {
	RequestId uint64 `json:"request_id,omitempty"` // 请求唯一id
	TraceId   string `json:"trace_id,omitempty"`   // trace_id
	Timestamp uint64 `json:"timestamp,omitempty"`  // 请求时间戳（精确到毫秒）
	Timeout   uint32 `json:"timeout,omitempty"`    // 请求超时时间，单位ms
	Caller    string `json:"caller,omitempty"`     // 主调服务的名称 app.server.service
	Appid     uint64 `json:"appid,omitempty"`      // appid
	Ip        string `json:"ip,omitempty"`         // ip 地址
}

// Request 插件请求信息
type Request struct {
	// query base info
	Op     string    `json:"op,omitempty"`     // operation
	Tables []string  `json:"tables,omitempty"` // tables or elastic index or ...
	Column []string  `json:"column,omitempty"` // columns
	Where  types.Map `json:"where,omitempty"`  // query condition
	Order  []string  `json:"order,omitempty"`  // order by
	Page   int       `json:"page,omitempty"`   // request pages. when page > 0, the request is returned in pagination.
	Size   int       `json:"size,omitempty"`   // size per page
	From   uint64    `json:"from,omitempty"`   // offset

	// data maintain
	Val   interface{}              `json:"val,omitempty"`   // 单条记录 val (not map/[]map)
	Data  map[string]interface{}   `json:"data,omitempty"`  // maintain one map data
	Datas []map[string]interface{} `json:"datas,omitempty"` // maintain multiple map data
	Args  []interface{}            `json:"args,omitempty"`  // multiple args, 还可用于 query 语句的参数，或者 redis 协议，如 MGET、HMGET、HDEL 等

	// group by
	Group  []string  `json:"group,omitempty"`  // group by
	Having types.Map `json:"having,omitempty"` // group by condition

	// for some other databases such as mysql ...
	Join []*sql.Join `json:"join,omitempty"`

	// for some other databases such as elastic ...
	Type   string        `json:"type,omitempty"`   // type, such as elastic`s type, it can be customized before v7, and unified as _doc after v7
	Scroll *proto.Scroll `json:"scroll,omitempty"` // scroll info

	// for some other databases such as redis ...
	Prefix string `json:"prefix,omitempty"` // prefix, it is strongly recommended to bring it to facilitate finer-grained summary statistics, otherwise the statistical granularity can only be cmd ，such as GET、SET、HGET ...
	Key    string `json:"key,omitempty"`    // key
	Field  string `json:"field,omitempty"`  // redis hash field

	// bytes 字节流
	Bytes []byte `json:"bytes,omitempty"`

	// params 与数据库特性相关的附加参数，例如 redis 的 withscores、EX、NX、等，以及 elastic 的 refresh、collapse、runtime_mappings、track_total_hits 等等。
	Params types.Map `json:"params,omitempty"`

	// query
	Query string `json:"query,omitempty"` // 直接送 query 语句，需要拥有库的表操作权限、或 root 权限。具体参数为 args

	// db address will be changing if Addr is set by plugin
	Addr *util.DBAddress
}

// Response 插件返回信息
type Response struct {
	Error  error         // 返回错误
	IsNil  bool          // 结果是否为空数据
	Detail *proto.Detail // 查询细节信息
	Result interface{}   // 返回结果
}
