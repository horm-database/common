package structs

type Type int8

const ( /* 我们发送请求到数据统一调度服务的时候，默认情况不需要指定数据类型，但是在某些情况下，比如 clickhouse 对类型有强限制，
	 需要指定具体的类型，又或者一个超大的 uint64 整数，json.Marshal 编码之后请求服务端，在服务端会被转化为 float64，存在精度丢失问题，
	这是因为 json 的基础类型只包含 string、number(当成float64)、bool ，所以这些情况，需要在 Unit.DataType 里面将数据类型上传，
	仅当类型为 time、[]byte、int、int8~int64、uint、uint8~uint64 时需要上传。*/
	TypeTime   Type = 1 // 类型是 time.Time
	TypeBytes  Type = 2 // 类型是 []byte
	TypeInt    Type = 3
	TypeInt8   Type = 4
	TypeInt16  Type = 5
	TypeInt32  Type = 6
	TypeInt64  Type = 7
	TypeUint   Type = 8
	TypeUint8  Type = 9
	TypeUint16 Type = 10
	TypeUint32 Type = 11
	TypeUint64 Type = 12
	TypeFloat  Type = 13
	TypeDouble Type = 14
	TypeString Type = 15
	TypeBool   Type = 16
	TypeJSON   Type = 17
)

var TypeDesc = map[string]Type{
	"time":   TypeTime,
	"bytes":  TypeBytes,
	"int":    TypeInt,
	"int8":   TypeInt8,
	"int16":  TypeInt16,
	"int32":  TypeInt32,
	"int64":  TypeInt64,
	"uint":   TypeUint,
	"uint8":  TypeUint8,
	"uint16": TypeUint16,
	"uint32": TypeUint32,
	"uint64": TypeUint64,
	"float":  TypeFloat,
	"double": TypeDouble,
	"string": TypeString,
	"bool":   TypeBool,
	"json":   TypeJSON,
}
