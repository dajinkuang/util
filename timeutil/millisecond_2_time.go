package timeutil

import "time"

func Millisecond2Time(millisecondToTime int64) time.Time {
	return time.Unix(millisecondToTime/1000, 0)
}

func Time2Millisecond(t time.Time) int64 {
	return t.UnixNano() / 1000000
}
