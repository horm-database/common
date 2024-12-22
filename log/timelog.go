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

package log

import (
	"context"
	"fmt"
	"time"

	"github.com/horm-database/common/codec"
	"github.com/horm-database/common/log/logger"
)

// TimeLog 耗时日志
type TimeLog struct {
	ctx       context.Context
	start     time.Time     // 开始时间
	threshold time.Duration // 告警类日志耗时阈值
}

func NewTimeLog(ctx context.Context, threshold ...time.Duration) *TimeLog {
	tl := TimeLog{
		ctx:   ctx,
		start: time.Now(),
	}

	if len(threshold) > 0 {
		tl.threshold = threshold[0]
	}

	return &tl
}

// Start 新开启一个 TimeLog 重新记时
func (t *TimeLog) Start(ctx context.Context) *TimeLog {
	return &TimeLog{
		ctx:       ctx,
		start:     time.Now(),
		threshold: t.threshold,
	}
}

// During 获取 during
func (t *TimeLog) During() time.Duration {
	return time.Since(t.start)
}

// SetThreshold 告警阈值，接口化
func (t *TimeLog) SetThreshold(i time.Duration) {
	t.threshold = i
}

// OverThreshold 判断是否超过阈值
func (t *TimeLog) OverThreshold() bool {
	return t.During() > t.threshold
}

// Debug debug 带耗时的调试日志
func (t *TimeLog) Debug(args ...interface{}) {
	msg := codec.Message(t.ctx)

	fields := []logger.Field{}
	fields = append(fields, logger.Field{"files", GetTraceback()})
	fields = append(fields, logger.Field{"seq", msg.LogSeq()})
	fields = append(fields, logger.Field{"during", fmt.Sprintf("%d", t.During().Milliseconds())})

	GetLogger(msg).Debug(fmt.Sprint(args...), fields...)
}

// Debugf debug 带耗时的调试日志
func (t *TimeLog) Debugf(format string, args ...interface{}) {
	msg := codec.Message(t.ctx)

	fields := []logger.Field{}
	fields = append(fields, logger.Field{"files", GetTraceback()})
	fields = append(fields, logger.Field{"seq", msg.LogSeq()})
	fields = append(fields, logger.Field{"during", fmt.Sprintf("%d", t.During().Milliseconds())})

	GetLogger(msg).Debug(fmt.Sprintf(format, args...), fields...)
}

// Info 带耗时的消息日志
func (t *TimeLog) Info(args ...interface{}) {
	msg := codec.Message(t.ctx)

	fields := []logger.Field{}
	fields = append(fields, logger.Field{"seq", msg.LogSeq()})
	fields = append(fields, logger.Field{"during", fmt.Sprintf("%d", t.During().Milliseconds())})

	GetLogger(msg).Info(fmt.Sprint(args...), fields...)
}

// Infof 带耗时的消息日志
func (t *TimeLog) Infof(format string, args ...interface{}) {
	msg := codec.Message(t.ctx)

	fields := []logger.Field{}
	fields = append(fields, logger.Field{"seq", msg.LogSeq()})
	fields = append(fields, logger.Field{"during", fmt.Sprintf("%d", t.During().Milliseconds())})

	GetLogger(msg).Info(fmt.Sprintf(format, args...), fields...)
}

// Warn 带耗时的警告日志
func (t *TimeLog) Warn(args ...interface{}) {
	msg := codec.Message(t.ctx)

	fields := []logger.Field{}
	fields = append(fields, logger.Field{"files", GetTraceback()})
	fields = append(fields, logger.Field{"seq", msg.LogSeq()})
	fields = append(fields, logger.Field{"during", fmt.Sprintf("%d", t.During().Milliseconds())})

	GetLogger(msg).Warn(fmt.Sprint(args...), fields...)

}

// Warnf 带耗时的警告日志
func (t *TimeLog) Warnf(format string, args ...interface{}) {
	msg := codec.Message(t.ctx)

	fields := []logger.Field{}
	fields = append(fields, logger.Field{"files", GetTraceback()})
	fields = append(fields, logger.Field{"seq", msg.LogSeq()})
	fields = append(fields, logger.Field{"during", fmt.Sprintf("%d", t.During().Milliseconds())})

	GetLogger(msg).Warn(fmt.Sprintf(format, args...), fields...)
}

// Error 带耗时的错误日志
func (t *TimeLog) Error(code int, args ...interface{}) {
	msg := codec.Message(t.ctx)

	fields := []logger.Field{}
	fields = append(fields, logger.Field{"files", GetTraceback()})
	fields = append(fields, logger.Field{"seq", msg.LogSeq()})
	fields = append(fields, logger.Field{"code", code})
	fields = append(fields, logger.Field{"during", fmt.Sprintf("%d", t.During().Milliseconds())})

	GetLogger(msg).Error(fmt.Sprint(args...), fields...)
}

// Errorf 带耗时的错误日志
func (t *TimeLog) Errorf(code int, format string, args ...interface{}) {
	msg := codec.Message(t.ctx)

	fields := []logger.Field{}
	fields = append(fields, logger.Field{"files", GetTraceback()})
	fields = append(fields, logger.Field{"seq", msg.LogSeq()})
	fields = append(fields, logger.Field{"code", code})
	fields = append(fields, logger.Field{"during", fmt.Sprintf("%d", t.During().Milliseconds())})

	GetLogger(msg).Error(fmt.Sprintf(format, args...), fields...)
}
