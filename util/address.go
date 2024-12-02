package util

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/horm-database/common/consts"
	"github.com/horm-database/common/errs"
	"github.com/horm-database/common/log"
	"github.com/horm-database/common/types"
	"github.com/horm-database/common/url"
)

type DBAddress struct {
	Type    int    `json:"type,omitempty"`    // 数据库类型 0-nil（仅执行插件） 1-elastic 2-mongo 3-redis 10-mysql 11-postgresql 12-clickhouse 13-oracle 14-DB2 15-sqlite
	Version string `json:"version,omitempty"` // 数据库版本，比如elastic v6，v7
	Network string `json:"network,omitempty"` // network TCP/UDP
	Address string `json:"address,omitempty"` // address

	Conn *DBConnInfo `json:"conn,omitempty"` // connect info

	WriteTimeout int `json:"write_timeout,omitempty"` // 写超时（毫秒）
	ReadTimeout  int `json:"read_timeout,omitempty"`  // 读超时（毫秒）

	WarnTimeout int  `json:"warn_timeout,omitempty"` // 告警超时（ms），如果请求耗时超过这个时间，就会打 warning 日志
	OmitError   int8 `json:"omit_error,omitempty"`   // 是否忽略 error 日志，0-否 1-是
	Debug       int8 `json:"debug,omitempty"`        // 是否开启 debug 日志，正常的数据库请求也会被打印到日志，0-否 1-是，会造成海量日志，慎重开启
}

type DBConnInfo struct {
	Schema   string `json:"schema,omitempty"`   // schema
	Target   string `json:"target,omitempty"`   // target
	DB       string `json:"db,omitempty"`       // db name
	Password string `json:"password,omitempty"` // password
	Params   string `json:"params,omitempty"`   // params
	DSN      string `json:"dsn,omitempty"`      // dsn
	DSN2     string `json:"dsn2,omitempty"`     // dsn2
	DSN3     string `json:"dsn3,omitempty"`     // dsn3
}

var (
	DBConnMap     = map[int]map[string]*DBConnInfo{}
	DBConnMapLock = new(sync.RWMutex)
)

// ParseConnFromAddress 解析数据库网络信息
func ParseConnFromAddress(addr *DBAddress) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("parse address failed: %v", e)
		}
	}()

	if addr == nil {
		return errors.New("address is nil")
	}

	if addr.Type == consts.DBTypeNil {
		return nil
	}

	if addr.Address == "" {
		return errors.New("address is empty")
	}

	DBConnMapLock.RLock()
	dbConnMap, ok := DBConnMap[addr.Type]

	if ok {
		conn, ok := dbConnMap[addr.Address]
		DBConnMapLock.RUnlock()

		if ok {
			addr.Conn = conn
			return
		}
	} else {
		DBConnMapLock.RUnlock()

		DBConnMapLock.Lock()
		DBConnMap[addr.Type] = map[string]*DBConnInfo{}
		DBConnMapLock.Unlock()
	}

	addr.Conn, err = getConnFromAddress(addr.Type, addr.Address)
	if err != nil {
		return err
	}

	DBConnMapLock.Lock()
	DBConnMap[addr.Type][addr.Address] = addr.Conn
	DBConnMapLock.Unlock()

	return nil
}

func getConnFromAddress(typ int, address string) (*DBConnInfo, error) {
	conn := DBConnInfo{}
	conn.Schema = "ip"

	switch typ {
	case consts.DBTypeRedis, consts.DBTypeElastic:

		_, conn.Schema, conn.Target = types.CutString(address, "://")

		target, _, params, e := ParseTarget(conn.Target)
		if e != nil {
			return nil, e
		}

		if target == "" {
			return nil, errors.New("target is empty")
		}

		conn.DSN = target
		conn.Target = target
		conn.DB, _ = params["db"]
		conn.Password, _ = params["password"]

		delete(params, "password")
		conn.Params = url.ParamEncode(params)
	case consts.DBTypeMySQL, consts.DBTypePostgreSQL, consts.DBTypeClickHouse, consts.DBTypeOracle, consts.DBTypeDB2:
		var network = "tcp"

		found, password, target := types.CutString(address, "@tcp(")
		if !found {
			found, password, target = types.CutString(address, "@udp(")
			network = "udp"
		}

		if !found {
			return nil, errors.New("address need include @tcp or @udp")
		}

		_, conn.Schema, conn.Password = types.CutString(password, "://")
		if conn.Password == "" {
			return nil, errors.New("password is empty")
		}

		_, conn.Target, conn.DB = types.CutString(target, ")/")
		if conn.Target == "" {
			return nil, errors.New("target is empty")
		}

		db, _, params, _ := ParseTarget(conn.DB)
		if db == "" {
			return nil, errors.New("db is empty")
		}

		conn.DB = db

		if typ == consts.DBTypeClickHouse {
			conn.DSN = fmt.Sprintf("dsn://%s@%s(%s)/%s?%s",
				conn.Password, network, conn.Target, conn.DB, url.ParamEncode(params))
		} else {
			conn.DSN = fmt.Sprintf("%s@%s(%s)/%s?%s",
				conn.Password, network, conn.Target, conn.DB, url.ParamEncode(params))
		}
	default:
		_, conn.Schema, conn.Target = types.CutString(address, "://")
		target, params, _, e := ParseTarget(conn.Target)
		if e != nil {
			return nil, e
		}

		if target == "" {
			return nil, errors.New("target is empty")
		}

		conn.Target = target
		conn.Params = params
		conn.DSN = address
	}

	return &conn, nil
}

// ParseTarget address should be like: target?param1=value1&param2=value2&param3=value3 ...
func ParseTarget(address string) (target, params string, paramsMap map[string]string, err error) {
	if address == "" {
		return
	}

	found, s1, s2 := types.CutString(address, "?")
	if found {
		target = s1
		params = s2
	} else {
		target = s2
	}

	paramsMap, err = url.ParseQuery(params)
	return
}

var localIP string
var localIPOnce = new(sync.Once)

// GetLocalIP 获取本机 ip 地址
func GetLocalIP() string {
	localIPOnce.Do(
		func() {
			addrs, err := net.InterfaceAddrs()
			if err != nil {
				log.Error(context.Background(), errs.RetSystem, "get local address error")
				return
			}

			for _, address := range addrs {
				if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
					if ipNet.IP.To4() != nil {
						localIP = ipNet.IP.String()
					}
				}
			}

			return
		})

	return localIP
}

func GetIpFromAddr(addr net.Addr) string {
	ipPort := addr.String()
	cutIpPort := strings.Index(ipPort, ":")

	switch cutIpPort {
	case -1:
		return ipPort
	case 0:
		return ipPort[1:]
	default:
		return ipPort[0:cutIpPort]
	}
}
