package logger

import (
	"sync"
)

var (
	DefaultLogger Logger

	loggers = map[string]Logger{}
	mu      = new(sync.RWMutex)
)

func init() {
	RegisterWriter(WriterConsole, &ConsoleWriter{})
	RegisterWriter(WriterFile, &FileWriter{})
	DefaultLogger = newZapLog([]*Config{{Writer: "console", Level: "debug"}}) // 初始化为 console ，会被配置的日志所覆盖
}

// Set 设置自定义日志打印器
func Set(name string, logger Logger) {
	mu.Lock()
	loggers[name] = logger
	mu.Unlock()
}

// Get 获取自定义日志打印器
func Get(name string) Logger {
	mu.RLock()
	l := loggers[name]
	mu.RUnlock()
	return l
}

// CreateDefaultLogger 根据配置创建默认 DefaultLogger
func CreateDefaultLogger(cfg []*Config) {
	if len(cfg) == 0 {
		panic("log config empty")
	}

	logger := newZapLog(cfg)
	if logger == nil {
		panic("new zap log fail")
	}

	DefaultLogger = logger
}

// Sync 同步所有日志
func Sync() {
	mu.RLock()
	defer mu.RUnlock()

	_ = DefaultLogger.Sync()

	for _, logger := range loggers {
		_ = logger.Sync()
	}
}
