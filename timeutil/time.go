package timeutil

import "time"

const (
	StdTimeFmt = "2006-01-02 15:04:05"
)

func ZeroTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}
