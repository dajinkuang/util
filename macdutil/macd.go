package macdutil

// ComputeMACD 计算K线的macd
func ComputeMACD(dif, dea float64) (macd float64) {
	macd = (dif - dea) * 2
	return
}
