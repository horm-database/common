package log

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/horm-database/common/codec"
	"github.com/horm-database/common/errs"
	"github.com/horm-database/common/log/logger"
)

// GetLogger 从上下文中获取logger
func GetLogger(msg *codec.Msg) logger.Logger {
	l := msg.Logger()
	if l == nil {
		return logger.DefaultLogger
	}

	return l
}

// Debug 调试日志
func Debug(ctx context.Context, args ...interface{}) {
	msg := codec.Message(ctx)

	GetLogger(msg).Debug(fmt.Sprint(args...),
		logger.Field{"files", GetTraceback()},
		logger.Field{"seq", msg.LogSeq()})
}

// Debugf 调试日志
func Debugf(ctx context.Context, format string, args ...interface{}) {
	msg := codec.Message(ctx)

	GetLogger(msg).Debug(fmt.Sprintf(format, args...),
		logger.Field{"files", GetTraceback()},
		logger.Field{"seq", msg.LogSeq()})
}

// Info 消息日志
func Info(ctx context.Context, args ...interface{}) {
	msg := codec.Message(ctx)

	GetLogger(msg).Info(fmt.Sprint(args...), logger.Field{"seq", msg.LogSeq()})
}

// Infof 消息日志
func Infof(ctx context.Context, format string, args ...interface{}) {
	msg := codec.Message(ctx)

	GetLogger(msg).Info(fmt.Sprintf(format, args...), logger.Field{"seq", msg.LogSeq()})
}

// Warn 警告日志
func Warn(ctx context.Context, args ...interface{}) {
	msg := codec.Message(ctx)

	GetLogger(msg).Warn(fmt.Sprint(args...),
		logger.Field{"files", GetTraceback()},
		logger.Field{"seq", msg.LogSeq()})
}

// Warnf 警告日志
func Warnf(ctx context.Context, format string, args ...interface{}) {
	msg := codec.Message(ctx)

	GetLogger(msg).Warn(fmt.Sprintf(format, args...),
		logger.Field{"files", GetTraceback()},
		logger.Field{"seq", msg.LogSeq()})
}

// Error 错误日志
func Error(ctx context.Context, code int, args ...interface{}) {
	msg := codec.Message(ctx)

	GetLogger(msg).Error(fmt.Sprint(args...),
		logger.Field{"files", GetTraceback()},
		logger.Field{"seq", msg.LogSeq()},
		logger.Field{"code", code})
}

// Errorf 错误日志
func Errorf(ctx context.Context, code int, format string, args ...interface{}) {
	msg := codec.Message(ctx)

	GetLogger(msg).Error(fmt.Sprintf(format, args...),
		logger.Field{"files", GetTraceback()},
		logger.Field{"seq", msg.LogSeq()},
		logger.Field{"code", code})
}

// Fatal fatal 日志
func Fatal(ctx context.Context, args ...interface{}) {
	msg := codec.Message(ctx)

	// 不调用 FatalContext 打日志，以免系统异常退出，崩溃日志错误码默认为 8888
	GetLogger(msg).Error(fmt.Sprint(args...),
		logger.Field{"files", GetTraceback()},
		logger.Field{"seq", msg.LogSeq()},
		logger.Field{"code", errs.RetPanic})
}

// Fatalf fatal 日志
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	msg := codec.Message(ctx)

	// 不调用 FatalContext 打日志，以免系统异常退出，崩溃日志错误码默认为 8888
	GetLogger(msg).Error(fmt.Sprintf(format, args...),
		logger.Field{"files", GetTraceback()},
		logger.Field{"seq", msg.LogSeq()},
		logger.Field{"code", errs.RetPanic})
}

// DebugWith 调试日志，带用户自定义上报字段 addFields
func DebugWith(ctx context.Context, addFields []logger.Field, args ...interface{}) {
	msg := codec.Message(ctx)

	fields := []logger.Field{
		{"files", GetTraceback()},
		{"seq", msg.LogSeq()},
	}

	fields = append(fields, addFields...)

	GetLogger(msg).Debug(fmt.Sprint(args...), fields...)
}

// DebugWithf 调试日志，带用户自定义上报字段 addFields
func DebugWithf(ctx context.Context, addFields []logger.Field, format string, args ...interface{}) {
	msg := codec.Message(ctx)

	fields := []logger.Field{
		{"files", GetTraceback()},
		{"seq", msg.LogSeq()},
	}

	fields = append(fields, addFields...)

	GetLogger(msg).Debug(fmt.Sprintf(format, args...), fields...)
}

// InfoWith 消息日志，带用户自定义上报字段 addFields
func InfoWith(ctx context.Context, addFields []logger.Field, args ...interface{}) {
	msg := codec.Message(ctx)

	fields := []logger.Field{
		{"seq", msg.LogSeq()},
	}

	fields = append(fields, addFields...)

	GetLogger(msg).Info(fmt.Sprint(args...), fields...)
}

// InfoWithf 消息日志，带用户自定义上报字段 addFields
func InfoWithf(ctx context.Context, addFields []logger.Field, format string, args ...interface{}) {
	msg := codec.Message(ctx)

	fields := []logger.Field{
		{"seq", msg.LogSeq()},
	}

	fields = append(fields, addFields...)

	GetLogger(msg).Info(fmt.Sprintf(format, args...), fields...)
}

// WarnWith 警告日志，带用户自定义上报字段 addFields
func WarnWith(ctx context.Context, addFields []logger.Field, args ...interface{}) {
	msg := codec.Message(ctx)

	fields := []logger.Field{
		{"files", GetTraceback()},
		{"seq", msg.LogSeq()},
	}

	fields = append(fields, addFields...)

	GetLogger(msg).Warn(fmt.Sprint(args...), fields...)

}

// WarnWithf 警告日志，带用户自定义上报字段 addFields
func WarnWithf(ctx context.Context, addFields []logger.Field, format string, args ...interface{}) {
	msg := codec.Message(ctx)

	fields := []logger.Field{
		{"files", GetTraceback()},
		{"seq", msg.LogSeq()},
	}

	fields = append(fields, addFields...)

	GetLogger(msg).Warn(fmt.Sprintf(format, args...), fields...)
}

// ErrorWith 错误日志，带用户自定义上报字段 addFields
func ErrorWith(ctx context.Context, addFields []logger.Field, code int, args ...interface{}) {
	msg := codec.Message(ctx)

	fields := []logger.Field{
		{"files", GetTraceback()},
		{"seq", msg.LogSeq()},
		{"code", code},
	}

	fields = append(fields, addFields...)

	GetLogger(msg).Error(fmt.Sprint(args...), fields...)
}

// ErrorWithf 错误日志，带用户自定义上报字段 addFields
func ErrorWithf(ctx context.Context, addFields []logger.Field, code int, format string, args ...interface{}) {
	msg := codec.Message(ctx)

	fields := []logger.Field{
		{"files", GetTraceback()},
		{"seq", msg.LogSeq()},
		{"code", code},
	}

	fields = append(fields, addFields...)

	GetLogger(msg).Error(fmt.Sprintf(format, args...), fields...)
}

// FatalWith panic 等崩溃日志，recover时打日志，带用户自定义上报字段 addFields
func FatalWith(ctx context.Context, addFields []logger.Field, args ...interface{}) {
	msg := codec.Message(ctx)

	fields := []logger.Field{
		{"files", GetTraceback()},
		{"seq", msg.LogSeq()},
		{"code", errs.RetPanic},
	}

	fields = append(fields, addFields...)

	// 不调用 FatalContext 打日志，以免系统异常退出，崩溃日志错误码默认为8888
	GetLogger(msg).Error(fmt.Sprint(args...), fields...)
}

// FatalWithf panic 等崩溃日志，带用户自定义上报字段 addFields
func FatalWithf(ctx context.Context, addFields []logger.Field, format string, args ...interface{}) {
	msg := codec.Message(ctx)

	fields := []logger.Field{
		{"files", GetTraceback()},
		{"seq", msg.LogSeq()},
		{"code", errs.RetPanic},
	}

	fields = append(fields, addFields...)

	GetLogger(msg).Error(fmt.Sprintf(format, args...), fields...)
}

// GetTraceback 调用栈
func GetTraceback() (ret string) {
	var stacks []string

	for i := 2; i < 15; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		// 获取计数器所在函数名，文件名
		methodName := runtime.FuncForPC(pc).Name()
		fileName := filepath.Base(file)

		if fileName == "service.go" && strings.HasSuffix(methodName, "srv.apiHandle") {
			break
		} else if strings.HasPrefix(methodName, "main.(") {
			stacks = append(stacks, getStackInfo(fileName, methodName, line))
			break
		} else if fileName == "" {
			break
		}

		stacks = append(stacks, getStackInfo(fileName, methodName, line))
	}

	var start = len(stacks)

	retBuilder := strings.Builder{}
	if start > 0 {
		retBuilder.WriteString(stacks[start-1])

		for i := start - 2; i >= 0; i-- {
			retBuilder.WriteString(" => ")
			retBuilder.WriteString(stacks[i])
		}
	}

	return retBuilder.String()
}

func getStackInfo(fileName, methodName string, line int) string {
	filePath := fileName
	method := methodName

	lastIndex := strings.LastIndex(methodName, "/")
	if lastIndex != -1 {
		method = methodName[lastIndex+1:]

		firstPoint := strings.Index(method, ".")

		var pkgName string

		if firstPoint != -1 {
			pkgName = method[:firstPoint] + "/"
			method = method[firstPoint+1:]
		}

		filePath = methodName[:lastIndex+1] + pkgName + fileName
	}
	return fmt.Sprintf("%s:%d  %s()", filePath, line, method)
}
