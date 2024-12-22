// Copyright (c) 2024 The horm-database Authors (such as CaoHao <18500482693@163.com>). All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
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
