package util

// min 函数用于返回两个整数中的最小值
func Min[T int | float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// max 函数用于返回两个整数中的最大值
func Max[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}
