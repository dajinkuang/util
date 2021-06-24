package timeutil

import "time"

// ClearTimer 清空timer
func ClearTimer(pendingClearTimer *time.Timer) {
	if pendingClearTimer == nil {
		return
	}
	if pendingClearTimer.Stop() || len(pendingClearTimer.C) <= 0 {
		return
	}
	var i = 0
	for ; i < len(pendingClearTimer.C); i++ {
		_ = <-pendingClearTimer.C
		continue
	}
	return
}
