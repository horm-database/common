package plugin

import (
	"github.com/horm-database/common/proto"
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

	// data update
	Data  types.Map                `json:"data,omitempty"`  // add/update one data
	Datas []map[string]interface{} `json:"datas,omitempty"` // batch add/update data

	// group by
	Group  []string  `json:"group,omitempty"`  // group by
	Having types.Map `json:"having,omitempty"` // group by condition

	// for some other databases such as elastic ...
	Type   string        `json:"type,omitempty"`   // type, such as elastic`s type, it can be customized before v7, and unified as _doc after v7
	Scroll *proto.Scroll `json:"scroll,omitempty"` // scroll info

	// for some other databases such as redis ...
	Prefix string        `json:"prefix,omitempty"` // prefix, it is strongly recommended to bring it to facilitate finer-grained summary statistics, otherwise the statistical granularity can only be cmd ，such as GET、SET、HGET ...
	Key    string        `json:"key,omitempty"`    // key
	Args   []interface{} `json:"args,omitempty"`   // args

	// bytes 字节流
	Bytes []byte `json:"bytes,omitempty"`

	// params 与数据库特性相关的附加参数，例如 mysql 的join，redis 的 WITHSCORES，以及 elastic 的 refresh、collapse、runtime_mappings、track_total_hits 等等。
	Params types.Map `json:"params,omitempty"`

	// query
	Query string `json:"query,omitempty"` // 直接送 query 语句，需要拥有库的 表权限、或 root 权限。具体参数为 args

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
