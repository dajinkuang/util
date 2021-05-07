package retryutil

// Retry 函数调用重试小框架
func Retry(limit int, do func() bool) {
	for i := 0; i < limit; i++ {
		continueRetry := do()
		if continueRetry {
			continue
		}
	}
}
