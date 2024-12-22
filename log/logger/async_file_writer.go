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
	"bytes"
	"errors"
	"io"
	"time"

	"github.com/hashicorp/go-multierror"
)

// AsyncOptions 异步写文件日志配置
type AsyncOptions struct {
	QueueSize     int  // 日志队列大小，默认 10000
	WriteSize     int  // 日志写入阈值， 默认 16K
	WriteInterval int  // 异步写日志时间间隔，（单位 ms）默认 100ms
	DropLog       bool // 队列满的时候是否丢弃日志，默认 false
}

// AsyncFileWriter 异步写文件，实现 zapcore.WriteSyncer 接口。
type AsyncFileWriter struct {
	logger io.WriteCloser
	opts   *AsyncOptions

	logQueue chan []byte
	sync     chan struct{}
	syncErr  chan error
	close    chan struct{}
	closeErr chan error
}

// NewAsyncFileWriter create a new AsyncFileWriter.
func NewAsyncFileWriter(logger io.WriteCloser, dropLog bool) *AsyncFileWriter {
	opts := &AsyncOptions{
		QueueSize:     10000,
		WriteSize:     16 * 1024,
		WriteInterval: 100,
		DropLog:       dropLog,
	}

	w := &AsyncFileWriter{}
	w.logger = logger
	w.opts = opts
	w.logQueue = make(chan []byte, opts.QueueSize)
	w.sync = make(chan struct{})
	w.syncErr = make(chan error)
	w.close = make(chan struct{})
	w.closeErr = make(chan error)

	go w.batchWriteLog() // 开启协程异步批量写日志
	return w
}

// Write 实现 io.Writer 接口。
func (w *AsyncFileWriter) Write(data []byte) (int, error) {
	log := make([]byte, len(data))
	copy(log, data)
	if w.opts.DropLog {
		select {
		case w.logQueue <- log:
		default:
			return 0, errors.New("log queue is full")
		}
	} else {
		w.logQueue <- log
	}
	return len(data), nil
}

// Sync 日志同步，实现 zapcore.WriteSyncer 接口。
func (w *AsyncFileWriter) Sync() error {
	w.sync <- struct{}{}
	return <-w.syncErr
}

// Close 关闭当前写的日志，实现 io.Closer 接口。
func (w *AsyncFileWriter) Close() error {
	err := w.Sync()
	close(w.close)
	return multierror.Append(err, <-w.closeErr).ErrorOrNil()
}

// batchWriteLog 异步批量写日志
func (w *AsyncFileWriter) batchWriteLog() {
	buffer := bytes.NewBuffer(make([]byte, 0, w.opts.WriteSize*2))
	ticker := time.NewTicker(time.Millisecond * time.Duration(w.opts.WriteInterval))
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if buffer.Len() > 0 {
				_, _ = w.logger.Write(buffer.Bytes())
				buffer.Reset()
			}
		case data := <-w.logQueue:
			buffer.Write(data)
			if buffer.Len() >= w.opts.WriteSize {
				_, _ = w.logger.Write(buffer.Bytes())
				buffer.Reset()
			}
		case <-w.sync:
			var err error
			if buffer.Len() > 0 {
				_, e := w.logger.Write(buffer.Bytes())
				err = multierror.Append(err, e).ErrorOrNil()
				buffer.Reset()
			}
			size := len(w.logQueue)
			for i := 0; i < size; i++ {
				v := <-w.logQueue
				_, e := w.logger.Write(v)
				err = multierror.Append(err, e).ErrorOrNil()
			}
			w.syncErr <- err
		case <-w.close:
			w.closeErr <- w.logger.Close()
			return
		}
	}
}
