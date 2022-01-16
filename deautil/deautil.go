// Package deautil dea计算工具包
package deautil

// ComputeDea 计算K线的dea
func ComputeDea(frontDea, curDif float64) (dea float64) {
	dea = (frontDea*8)/10 + (curDif*2)/10
	return dea
}
