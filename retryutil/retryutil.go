// Package retryutil 函数或者方法调用重试框架
package retryutil

// Retry 函数调用重试小框架
func Retry(limit int, do func() bool) int {
	var i int
	for i = 0; i < limit; i++ {
		continueRetry := do()
		if !continueRetry {
			return i + 1
		}
	}
	return i
}
