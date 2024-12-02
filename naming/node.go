package naming

import (
	"time"
)

// Node is the information of a node.
type Node struct {
	ServiceName string        // service name
	Address     string        // target address ip:port
	Network     string        // network protocol tcp/udp
	Weight      int           // weight
	CostTime    time.Duration // request cost time
	Metadata    map[string]interface{}
}
