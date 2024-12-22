// Copyright (c) 2024 The horm-database Authors. All rights reserved.
// This file Author:  CaoHao <18500482693@163.com> .
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
	"gopkg.in/yaml.v3"
)

const ( //日志输出名，默认支持 console、file、online
	WriterConsole = "console"
	WriterFile    = "file"
)

// Config 日志输出配置，包括 console, file 和 third_party.
type Config struct {
	Writer string `yaml:"writer"` // 日志输出，例如 console 、file、esfile
	Level  string `yaml:"level"`  // 日志级别，例如 debug、info、warn、error、fatal

	Encoder       string        `yaml:"encoder"`        // 日志编码格式，比如 console、json
	EncoderConfig EncoderConfig `yaml:"encoder_config"` // 格式配置
	Field         []string      `yaml:"field"`          // 当采用 separator 的时候，fields 是按顺序提取的字段数据，各个数据用分隔符隔开。
	Escape        bool          `yaml:"escape"`         // 内容是否转义,性能原因默认关闭,true开启

	FileConfig       FileConfig `yaml:"file_config"`        // 文件日志配置
	ThirdPartyConfig yaml.Node  `yaml:"third_party_config"` // 第三方日志组件配置。它是由业务定义的，应该由第三方模块注册。
}

// FileConfig 文件输出配置
type FileConfig struct {
	LogPath    string `yaml:"log_path"`    // 日志路径，如 /usr/local/server/log/
	Filename   string `yaml:"filename"`    // 日志文件名，如 server.log
	WriteMode  int    `yaml:"write_mode"`  // 日志写入模式，1: sync, 2: async, 3: fast(队列满会丢弃日志).
	Compress   bool   `yaml:"compress"`    // 是否压缩
	LocalTime  bool   `yaml:"local_time"`  // 是否本地时间
	MaxDay     int    `yaml:"max_day"`     // 日志最大过期天数
	MaxBackups int    `yaml:"max_backups"` // 最大日志文件数
	MaxSize    int    `yaml:"max_size"`    // 本地文件滚动日志的大小 单位 MB
}

// EncoderConfig 编码配置
type EncoderConfig struct {
	TimeFmt       string `yaml:"time_fmt"`       // 日志输出的时间格式，默认 "2006-01-02 15:04:05.000"
	TimeKey       string `yaml:"time_key"`       // 日志输出的 时间 key，默认 "time"
	LevelKey      string `yaml:"level_key"`      // 日志输出的 级别 key，默认 "level"
	MessageKey    string `yaml:"message_key"`    // 日志输出的 消息 key，默认 "msg"
	NameKey       string `yaml:"name_key"`       // 日志输出的 名字 key，默认 ""
	FunctionKey   string `yaml:"function_key"`   // 日志输出的 函数 key，默认""，表示不打印函数名
	CallerKey     string `yaml:"caller_key"`     // 日志输出的 caller key，默认 ""
	StacktraceKey string `yaml:"stacktrace_key"` // 日志输出的 stack trace key，默认 ""
}

const ( //日志写入模式
	WriteSync  = 1 // 同步写入
	WriteAsync = 2 // 异步写入
	WriteFast  = 3 // 快速写入（队列满的时候会丢弃日志）
)

const ( // 本地文件默认配置
	defaultMaxDay  = 3   // 默认本地文件最大保存天数
	defaultMaxSize = 100 // 默认本地文件滚动日志的大小
)
