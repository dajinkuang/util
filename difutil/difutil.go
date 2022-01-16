// Package difutil dif计算
package difutil

// ComputeDifSecond 计算第2根K线的dif
func ComputeDifSecond(c1, c2 float64) (ema12, ema26, dif float64) {
	ema12 = c1 + (c2-c1)*2/13
	ema26 = c1 + (c2-c1)*2/27
	dif = ema12 - ema26
	return
}

// ComputeEma12Ema26Dif 计算第2根K线之后的K线的dif
func ComputeEma12Ema26Dif(frontEma12, frontEma26, c float64) (ema12, ema26, dif float64) {
	ema12 = frontEma12*11/13 + c*2/13
	ema26 = frontEma26*25/27 + c*2/27
	dif = ema12 - ema26
	return
}
