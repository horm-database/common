package logger

import (
	"errors"
	"path/filepath"

	"go.uber.org/zap/zapcore"
)

type Writer interface {
	Setup(cfg *Config) (zapcore.Core, error) // 根据配置setup writer
}

var (
	writers = make(map[string]Writer)
)

// RegisterWriter registers log output writer. Writer may have multiple implementations.
func RegisterWriter(name string, writer Writer) {
	writers[name] = writer
}

// GetWriter gets log output writer, returns nil if not exist.
func GetWriter(name string) Writer {
	return writers[name]
}

type ConsoleWriter struct{} // console writer
type FileWriter struct{}    // file writer

// Setup 加载注册 ConsoleWriter
func (f *ConsoleWriter) Setup(cfg *Config) (zapcore.Core, error) {
	if cfg == nil {
		return nil, errors.New("console writer output config empty")
	}

	return newConsoleCore(cfg), nil
}

// Setup 加载注册 FileWriter
func (f *FileWriter) Setup(cfg *Config) (zapcore.Core, error) {
	if cfg == nil {
		return nil, errors.New("file writer output config empty")
	}

	return newFileCore(fixFileConfig(cfg)), nil
}

func fixFileConfig(cfg *Config) *Config {
	if cfg.FileConfig.LogPath != "" {
		cfg.FileConfig.Filename = filepath.Join(cfg.FileConfig.LogPath, cfg.FileConfig.Filename)
	}

	if cfg.FileConfig.MaxSize != 0 {
		cfg.FileConfig.MaxSize = defaultMaxSize
	}

	if cfg.FileConfig.MaxDay != 0 {
		cfg.FileConfig.MaxDay = defaultMaxDay
	}

	if cfg.FileConfig.WriteMode == 0 {
		cfg.FileConfig.WriteMode = WriteFast // WriteFast 性能更好，但是会丢弃已满的日志并避免阻塞服务。
	}

	return cfg
}
