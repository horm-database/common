package proto

import (
	"strconv"

	"github.com/horm-database/common/errs"
)

// QueryResp 返回请求
type QueryResp struct {
	IsNil   bool              `orm:"is_nil,omitempty" json:"is_nil,omitempty"`     // 是否为空，仅针对单查询单元
	RspNils map[string]bool   `orm:"rsp_nils,omitempty" json:"rsp_nils,omitempty"` // 是否为空，针对并行查询
	RspErrs map[string]*Error `orm:"rsp_errs,omitempty" json:"rsp_errs,omitempty"` // 返回错误码，针对并行查询
	RspData interface{}       `orm:"rsp_data,omitempty" json:"rsp_data,omitempty"` // 返回数据
}

// PageResult 当 page > 1 时会返回分页结果
type PageResult struct {
	Detail *Detail       `orm:"detail,omitempty" json:"detail,omitempty"` // 查询细节信息
	Data   []interface{} `orm:"data,omitempty" json:"data,omitempty"`     // 分页结果
}

// CompResult 混合查询返回结果
type CompResult struct {
	RetBase             // 返回基础信息
	Data    interface{} `json:"data"` // 返回数据
}

// RetBase 混合查询返回结果基础信息
type RetBase struct {
	Error  *Error  `json:"error,omitempty"`  // 错误返回
	IsNil  bool    `json:"is_nil,omitempty"` // 是否为空
	Detail *Detail `json:"detail,omitempty"` // 查询细节信息
}

// Detail 其他查询细节信息，例如 分页信息、滚动翻页信息、其他信息等。
type Detail struct {
	Total     uint64                 `orm:"total" json:"total"`                               // 总数
	TotalPage uint32                 `orm:"total_page,omitempty" json:"total_page,omitempty"` // 总页数
	Page      int                    `orm:"page,omitempty" json:"page,omitempty"`             // 当前分页
	Size      int                    `orm:"size,omitempty" json:"size,omitempty"`             // 每页大小
	Scroll    *Scroll                `orm:"scroll,omitempty" json:"scroll,omitempty"`         // 滚动翻页信息
	Extras    map[string]interface{} `orm:"extras,omitempty" json:"extras,omitempty"`         // 更多详细信息
}

// ModRet 新增/更新返回信息
type ModRet struct {
	ID          ID                     `orm:"id,omitempty" json:"id,omitempty"`                       // id 主键，可能是 mysql 的最后自增id，last_insert_id 或 elastic 的 _id 等，类型可能是 int64、string
	RowAffected int64                  `orm:"rows_affected,omitempty" json:"rows_affected,omitempty"` // 影响行数
	Version     int64                  `orm:"version,omitempty" json:"version,omitempty"`             // 数据版本
	Status      int                    `orm:"status,omitempty" json:"status,omitempty"`               // 返回状态码
	Reason      string                 `orm:"reason,omitempty" json:"reason,omitempty"`               // mod 失败原因
	Extras      map[string]interface{} `orm:"extras,omitempty" json:"extras,omitempty"`               // 更多详细信息
}

// MemberScore redis 集合成员及其分数信息。
type MemberScore struct {
	Member []string  `orm:"member,omitempty" json:"member,omitempty"`
	Score  []float64 `orm:"score,omitempty" json:"score,omitempty"`
}

type ID string

func (id ID) String() string {
	return string(id)
}

func (id ID) Float64() float64 {
	f, _ := strconv.ParseFloat(string(id), 64)
	return f
}

func (id ID) Int() int {
	return int(id.Int64())
}

func (id ID) Int64() int64 {
	i, _ := strconv.ParseInt(string(id), 10, 64)
	return i
}

func (id ID) Uint() uint {
	return uint(id.Uint64())
}

func (id ID) Uint64() uint64 {
	ui, _ := strconv.ParseUint(string(id), 10, 64)
	return ui
}

func (x *Error) ToError() *errs.Error {
	return &errs.Error{
		Type: int8(x.Type),
		Code: int(x.Code),
		Msg:  x.Msg,
	}
}
