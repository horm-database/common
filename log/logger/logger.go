package logger

// Enums log level constants.
const (
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"
	LevelFatal = "fatal"
)

// Logger is the underlying logging work for framework.
type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field) // 所有 Fatal 日志会调用 os.Exit(1) 函数退出系统

	// 系统日志，无 field 信息
	Debugf(msg string, args ...interface{})
	Infof(msg string, args ...interface{})
	Warnf(msg string, args ...interface{})
	Errorf(msg string, args ...interface{})
	Fatalf(msg string, args ...interface{})

	Sync() error //刷新所有缓冲的日志，系统必须在退出之前调用 Sync

	With(fields ...Field) Logger
}

type Field struct {
	Key   string
	Value interface{}
}
