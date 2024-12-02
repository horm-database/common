// Package metrics reports the statistics of the framework.
package metrics

var (
	ServiceHandleFail           = Counter("ServiceHandleFail")
	ServiceCodecDecodeFail      = Counter("ServiceCodecDecodeFail")
	ServiceCodecEncodeFail      = Counter("ServiceCodecEncodeFail")
	ServiceHandleRPCNameInvalid = Counter("ServiceHandleRpcNameInvalid")
	ServiceCodecMarshalFail     = Counter("ServiceCodecMarshalFail")

	TCPServerTransportHandleFail = Counter("TcpServerTransportHandleFail")
	TCPServerTransportWriteFail  = Counter("TcpServerTransportWriteFail")

	SelectNodeFail   = Counter("SelectNodeFail")
	ClientCodecEmpty = Counter("ClientCodecEmpty")

	ConnectionPoolGetNewConnection = Counter("ConnectionPoolGetNewConnection")
	ConnectionPoolGetConnectionErr = Counter("ConnectionPoolGetConnectionErr")
	ConnectionPoolRemoteErr        = Counter("ConnectionPoolRemoteErr")
	ConnectionPoolRemoteEOF        = Counter("ConnectionPoolRemoteEOF")
	ConnectionPoolIdleTimeout      = Counter("ConnectionPoolIdleTimeout")
	ConnectionPoolLifetimeExceed   = Counter("ConnectionPoolLifetimeExceed")
	ConnectionPoolOverLimit        = Counter("ConnectionPoolOverLimit")
)
