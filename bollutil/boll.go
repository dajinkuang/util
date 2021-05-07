package bollutil

import (
	"github.com/dajinkuang/errors"
	"math"
)

// ComputeBoll 计算boll线指标上轨，中轨，下轨，N=20
func ComputeBoll(close20List []float64) (bollUp, bollMiddle, bollDown float64, err error) {
	if len(close20List) != 20 {
		err = errors.Newf("收盘价必须是连续20根K线")
		return
	}
	// 求标准差
	avg, S := computeAvgAndStdev(close20List)
	bollMiddle = avg
	bollUp = bollMiddle + 2*S
	bollDown = bollMiddle - 2*S
	return
}

// 计算平均数和标准差
func computeAvgAndStdev(closeList []float64) (avg, S float64) {
	// 计算平均数
	var sum float64
	for _, v := range closeList {
		sum = sum + v
	}
	avg = sum / float64(len(closeList))
	// 标准差
	var sumPow float64
	for _, v := range closeList {
		sumPow = sumPow + math.Pow(v-avg, 2)
	}
	D := sumPow / float64(len(closeList))
	S = math.Sqrt(D)
	return
}
