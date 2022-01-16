// Package numberutil 数值工具包
package numberutil

import (
	"fmt"
	"strconv"
)

// Float64Round 截取小数位数
func Float64Round(f float64, n int) float64 {
	format := "%." + strconv.Itoa(n) + "f"
	resultNumber, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	if resultNumber > f { // 特殊例子，179.79 格式化为 179.7
		resultNumber = resultNumber - DotPrecision(n)
		resultNumber, _ = strconv.ParseFloat(fmt.Sprintf(format, resultNumber), 64)
	}
	return resultNumber
}

// DotPrecision 根据float的小数精度位数，获取数值精度
// 比如：n=-1--->1 n=1--->0.1 n=2--->0.01
func DotPrecision(n int) float64 {
	if n <= 0 {
		return 1
	}
	var dotBytes []byte
	dotBytes = append(dotBytes, '0', '.')
	for i := 0; i < n; i++ {
		if i != n-1 {
			dotBytes = append(dotBytes, '0')
			continue
		}
		dotBytes = append(dotBytes, '1')
	}
	dotPrecisionNumber, _ := strconv.ParseFloat(string(dotBytes), 64)
	return dotPrecisionNumber
}
