package timeutil

import "time"

// ResetTimer 重置timer
func ResetTimer(pendingResetTimer *time.Timer) {
	if pendingResetTimer == nil {
		return
	}
	if pendingResetTimer.Stop() || len(pendingResetTimer.C) <= 0 {
		return
	}
	var i = 0
	for ; i < len(pendingResetTimer.C); i++ {
		_ = <-pendingResetTimer.C
		continue
	}
	return
}
