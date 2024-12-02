package proto

const (
	HeaderVersion      = "head-version"    // 客户端版本
	HeaderQueryMode    = "head-query-mode" // 查询模式 0-单执行单元（默认）1-多执行单元并行（不含嵌套子查询） 2-复合查询（包含嵌套子查询）
	HeaderRequestID    = "head-request-id" // 请求唯一id
	HeaderTraceID      = "head-trace-id"   // trace-id
	HeaderTimestamp    = "head-timestamp"  // 请求时间戳（精确到毫秒）
	HeaderTimeout      = "head-timeout"    // 请求超时时间，单位ms
	HeaderCaller       = "head-caller"     // 主调服务的名称 app.server.service
	HeaderAppid        = "head-appid"      // appid
	HeaderAuthRand     = "head-auth-rand"  // 随机生成 0-9999999 的数字，相同 timestamp 不允许出现同样的 ip、auth_rand。为了避免碰撞，0-9999999，单机理论最大支持 100 亿/秒的并发。
	HeaderSign         = "head-sign"       // sign 签名，为 md5(appid+secret+version+request_type+query_mode+request_id+trace_id+timestamp+timeout+caller+compress+ip+auth_rand)
	HeaderIsNil        = "head-is-nil"     // 返回是否为空（针对单执行单元）
	HeaderErrorType    = "head-error-type" // 错误类型
	HeaderErrorCode    = "head-error-code" // 错误码
	HeaderErrorMessage = "head-error-msg"  // 错误消息
	HeaderRspNils      = "head-rsp-nils"   // 是否为空返回，是一个 json 串，内容为：map[string]bool（针对多执行单元并发）
	HeaderRspErrs      = "head-rsp-errs"   // 错误返回，是一个json 串，内容为：map[string]Error（针对多执行单元并发）
)
