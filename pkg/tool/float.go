package tool

import (
	"math"
	"strconv"
)

// FormatFloat 指定浮点数精度
func FormatFloat(value float64, decimal int) string {
	// 默认乘1
	d := float64(1)
	if decimal > 0 {
		// 10的N次方
		d = math.Pow10(decimal)
	}
	// math.trunc作用就是返回浮点数的整数部分
	// 再除回去，小数点后无效的0也就不存在了
	return strconv.FormatFloat(math.Trunc(value*d)/d, 'f', -1, 64)
}
