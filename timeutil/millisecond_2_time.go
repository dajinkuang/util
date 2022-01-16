package timeutil

import "time"

// Millisecond2Time 毫秒转为时间
func Millisecond2Time(millisecondToTime int64) time.Time {
	return time.Unix(millisecondToTime/1000, 0)
}

// Time2Millisecond 时间转为毫秒
func Time2Millisecond(t time.Time) int64 {
	return t.UnixNano() / 1000000
}
