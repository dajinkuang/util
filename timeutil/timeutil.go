// Package timeutil 时间工具包
package timeutil

import "time"

const (
	// StdTimeFmt 标准的时间格式
	StdTimeFmt = "2006-01-02 15:04:05"
)

// ZeroTime 0点时间
func ZeroTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}
