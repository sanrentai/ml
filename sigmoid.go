package ml

import "math"

// Sigmoid 函数用于计算 sigmoid 函数值
func Sigmoids(inX []float64) []float64 {
	result := make([]float64, len(inX))
	for i, v := range inX {
		result[i] = Sigmoid(v)
	}
	return result
}

func Sigmoid(inX float64) float64 {
	return 1.0 / (1 + math.Exp(-inX))
}
