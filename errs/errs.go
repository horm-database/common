// Package errs provides error code type, which contains errcode errmsg.
// These definitions are multi-language universal.
package errs

import (
	"fmt"
	"io"
)

const (
	Success    = 0    // 成功
	RetSystem  = 1    // 服务端系统异常
	RetPanic   = 8888 // panic
	RetUnknown = 9999 // 未知错误

	// 客户端错误
	RetClientReadFrameFail     = 11 // 客户端帧读取失败
	RetClientTimeout           = 12 // 请求在客户端调用超时
	RetClientConnectFail       = 13 // 客户端连接错误
	RetClientEncodeFail        = 14 // 客户端编码错误
	RetClientDecodeFail        = 15 // 客户端解码错误
	RetClientRouteErr          = 16 // 客户端选ip路由错误
	RetClientNetErr            = 17 // 客户端网络错误
	RetClientCanceled          = 18 // 上游调用方提前取消请求
	RetClientParamInvalid      = 19 // 请求参数非法
	RetClientParamTypeInvalid  = 20 // 请求参数类型非法
	RetClientUnitNameEmpty     = 21 // 执行单元名不能为空
	RetClientNotInit           = 22 // server 未被初始化
	RetClientRequestIDNotMatch = 23 // 请求与返回 id 不匹配
	RetClientQueryModeNotMatch = 24 // 请求与返回 query mode 不匹配

	// 网关错误网关错误
	RetAppCallLimited = 60 // 限制时间内调用失败次数 网关错误

	// 服务器错误
	RetServerReadFrameFail  = 101 // 服务端帧读取失败
	RetServerDecompressFail = 102 // 服务端解压错误
	RetServerDecodeFail     = 103 // 服务端解码错误
	RetServerEncodeFail     = 104 // 服务端编码错误
	RetServerNoService      = 105 // 服务端没有调用相应的service实现
	RetServerNoFunc         = 106 // 服务端没有调用相应的接口实现
	RetServerTimeout        = 107 // 请求在服务端队列超时
	RetServerOverload       = 108 // 请求在服务端过载

	// 参数错误
	RetParamInvalid         = 301 // 请求参数非法
	RetParamEmpty           = 302 // 请求参数不得为空
	RetParamMiss            = 303 // 请求参数未上传
	RetParamType            = 304 // 请求参数类型错误
	RetParamValue           = 305 // 请求参数取值错误
	RetNotFindName          = 306 // 未找到 name 对应表配置
	RetUnitNameEmpty        = 307 // 执行单元名不能为空
	RetRepeatNameAlias      = 308 // 在同一层级有重复的 name 或 alias
	RetNotFindReferer       = 309 // 未找到被引用的执行单元
	RetRefererUnitFailed    = 310 // 被引用的执行单元查询失败
	RetRefererResultType    = 311 // 被引用的执行单元结果类型不符
	RetRefererFieldNotExist = 312 // 被引用的执行单元结果中不包含引用字段
	RetFormatDataError      = 313 // 数据格式化失败
	RetSameTransaction      = 314 // 事务重复定义

	// 权限错误
	RetServerAuthFail    = 401 // 鉴权失败
	RetHasNoTableRight   = 402 // 无该表访问权限
	RetHasNoDBRight      = 403 // 无数据库访问权限
	RetNotFindAppid      = 404 // 未找到 appid
	RetTableVerifyFailed = 405 // 表校验失败

	// 数据库错误
	RetQuery           = 500 // query error
	RetNotFindQueryImp = 501 // 未找到数据库的 Query 实现
	RetTransaction     = 502 // 事务错误
	RetDBReqParams     = 503 // database request params error

	RetSQLQuery           = 510 // mysql/postgresql/clickhouse query error
	RetAffectResultFailed = 512 // 获取影响行数信息失败

	RetClickhouseInsert = 530 // clickhouse insert error
	RetClickhouseCreate = 530 // clickhouse create error

	RetElastic = 550 // new elastic client error

	RetRedisDo           = 570 //redis do error
	RetRedisDecodeFailed = 571 //redis 结果解码 失败

	// 插件错误
	RetFilterConfigDecode    = 601 // 插件配置解压失败
	RetFilterNotFind         = 602 // 未找到插件
	RetFilterFuncNotRegister = 603 // 插件函数未注册
	RetFilterHandle          = 604 // 插件执行异常
	RetFilterConfigInvalid   = 605 // 插件配置异常
	RetFilterParamCopy       = 606 // 异步插件参数备份异常
	RetFilterFrontNotFind    = 607 // 插件先决执行插件未找到

	// 其他错误
	RetOpNotSupport  = 801 // 该数据库不支持该操作
	RetNameAmbiguity = 802 // 表有歧义，需要加 namespace

	RetNotFindDBConfig     = 851 // 未找到数据库配置
	RetDBConfigTypeInvalid = 852 // 数据库类型错误
	RetDBAddressParseError = 853 // 数据库地址解析错误
)

// ErrorType 错误类型
const (
	ErrorTypeSystem   = 0 //系统错误
	ErrorTypeFilter   = 1 //插件错误
	ErrorTypeDatabase = 2 //数据库错误
)

func typeDesc(t int8) string {
	switch t {
	case ErrorTypeFilter:
		return "filter"
	case ErrorTypeDatabase:
		return "database"
	default:
		return "system"
	}
}

// New 创建一个系统错误
func New(code int, msg string) error {
	return &Error{
		Type: ErrorTypeSystem,
		Code: code,
		Msg:  msg,
	}
}

// Newf 创建一个格式化系统错误
func Newf(code int, format string, params ...interface{}) error {
	return &Error{
		Type: ErrorTypeSystem,
		Code: code,
		Msg:  fmt.Sprintf(format, params...),
	}
}

// NewDBError 创建一个数据库错误
func NewDBError(code int, msg string) error {
	return &Error{
		Type: ErrorTypeDatabase,
		Code: code,
		Msg:  msg,
	}
}

// NewDBErrorf 创建一个格式化数据库错误
func NewDBErrorf(code int, format string, params ...interface{}) error {
	return &Error{
		Type: ErrorTypeDatabase,
		Code: code,
		Msg:  fmt.Sprintf(format, params...),
	}
}

// Type 获取错误类型
func Type(e error) int8 {
	if e == nil {
		return ErrorTypeSystem
	}

	err, ok := e.(*Error)
	if !ok {
		return ErrorTypeSystem
	}

	return err.Type
}

// Code 获取错误码
func Code(e error) int {
	if e == nil {
		return 0
	}

	err, ok := e.(*Error)
	if !ok {
		return RetUnknown
	}

	return err.Code
}

// Msg 获取错误信息
func Msg(e error) string {
	if e == nil {
		return "success"
	}

	err, ok := e.(*Error)
	if !ok {
		return e.Error()
	}

	return err.Msg
}

// SetErrorType 设置类型
func SetErrorType(err error, errorType int8) error {
	if err == nil {
		return nil
	}

	if Type(err) != errorType {
		return &Error{
			Type: errorType,
			Code: Code(err),
			Msg:  Msg(err),
		}
	}

	return err
}

// SetErrorCode 设置默认 code
func SetErrorCode(err error, defaultCode int) error {
	if err == nil {
		return nil
	}

	code := Code(err)

	if code == 0 || code == RetUnknown {
		return &Error{
			Type: Type(err),
			Code: defaultCode,
			Msg:  Msg(err),
		}
	}

	return err
}

// SetErrorMsg 设置错误消息
func SetErrorMsg(err error, msg string) error {
	if err == nil {
		return nil
	}

	e, ok := err.(*Error)
	if !ok {
		return &Error{
			Type: ErrorTypeSystem,
			Code: RetUnknown,
			Msg:  msg,
		}
	}

	e.Msg = msg
	return e
}

// SetErrorSQL 设置
func SetErrorSQL(err error, sql string) error {
	if err == nil {
		return nil
	}

	e, ok := err.(*Error)
	if !ok {
		return &Error{
			Type: ErrorTypeSystem,
			Code: RetUnknown,
			Msg:  "",
			SQL:  sql,
		}
	}

	e.SQL = sql
	return e
}

// Error error 结构体
type Error struct {
	Type int8
	Code int
	Msg  string
	SQL  string //发生 error 时的 sql 语句
}

// Error error 信息
func (e *Error) Error() string {
	if e == nil {
		return "success"
	}

	if e.SQL != "" {
		return fmt.Sprintf("type:%s, code:%d, msg:%s, sql=[%s]", typeDesc(e.Type), e.Code, e.Msg, e.SQL)
	} else {
		return fmt.Sprintf("type:%s, code:%d, msg:%s", typeDesc(e.Type), e.Code, e.Msg)
	}
}

// Format 实现 fmt.Formatter 接口
func (e *Error) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			if e.Msg != "" {
				msg := fmt.Sprintf("type:%s, code:%d, msg:%s", typeDesc(e.Type), e.Code, e.Msg)
				_, _ = io.WriteString(s, msg)
			}
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, e.Error())
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", e.Error())
	default:
		_, _ = fmt.Fprintf(s, "%%!%c(errs.Error=%s)", verb, e.Error())
	}
}
