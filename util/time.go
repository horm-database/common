package util

import (
	"time"
)

// GetMillisecond 返回毫秒 time.Duration
func GetMillisecond(sec int) time.Duration {
	return time.Millisecond * time.Duration(sec)
}
