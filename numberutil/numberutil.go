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
	return resultNumber
}
