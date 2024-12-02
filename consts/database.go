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
