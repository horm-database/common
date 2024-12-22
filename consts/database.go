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

const ( // 数据库类型
	DBTypeNil        = 0  // 空操作，仅走插件
	DBTypeElastic    = 1  // elastic search
	DBTypeMongo      = 2  // mongo 暂未支持
	DBTypeRedis      = 3  // redis
	DBTypeMySQL      = 10 // mysql
	DBTypePostgreSQL = 11 // postgresql
	DBTypeClickHouse = 12 // clickhouse
	DBTypeOracle     = 13 // oracle 暂未支持
	DBTypeDB2        = 14 // DB2 暂未支持
	DBTypeSQLite     = 15 // sqlite 暂未支持
	DBTypeRPC        = 40 // rpc 协议，暂未支持，spring cloud 协议可以选 grpc、thrift、tars、dubbo 协议
	DBTypeHTTP       = 50 // http 请求
)

var DBTypes = []int{
	DBTypeNil,
	DBTypeElastic,
	DBTypeMongo,
	DBTypeRedis,
	DBTypeMySQL,
	DBTypePostgreSQL,
	DBTypeClickHouse,
	DBTypeOracle,
	DBTypeDB2,
	DBTypeSQLite,
	DBTypeRPC,
	DBTypeHTTP,
}

var DBTypeMap = map[string]int{
	"nil":        DBTypeNil,
	"elastic":    DBTypeElastic,
	"redis":      DBTypeRedis,
	"mysql":      DBTypeMySQL,
	"postgresql": DBTypePostgreSQL,
	"clickhouse": DBTypeClickHouse,
	"oracle":     DBTypeOracle,
	"db2":        DBTypeDB2,
	"sqlite":     DBTypeSQLite,
	"mongo":      DBTypeMongo,
}

var DBTypeDesc = map[int]string{
	DBTypeNil:        "nil",
	DBTypeElastic:    "elastic",
	DBTypeRedis:      "redis",
	DBTypeMySQL:      "mysql",
	DBTypePostgreSQL: "postgresql",
	DBTypeClickHouse: "clickhouse",
	DBTypeOracle:     "oracle",
	DBTypeDB2:        "db2",
	DBTypeSQLite:     "sqlite",
	DBTypeMongo:      "mongo",
}
