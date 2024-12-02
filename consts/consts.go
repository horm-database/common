package consts

const (
	FALSE = 0
	TRUE  = 1
)

const (
	QueryModeSingle   = 0 //单个执行单元
	QueryModeParallel = 1 //多个执行单元并发
	QueryModeCompound = 2 //复合查询
)

var QueryModeDesc = map[int8]string{
	QueryModeSingle:   "single",
	QueryModeParallel: "parallel",
	QueryModeCompound: "compound",
}

const (
	NoCompression = 0
	Compression   = 1
)

const ( //request type
	RequestTypeRPC  = 0
	RequestTypeHTTP = 1
	RequestTypeWeb  = 2
)

const (
	StatusOnline  = 1 // 正常
	StatusOffline = 2 // 下线
)
