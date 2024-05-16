package ml

import "math"

// Dot product of two vectors
func Dot(a, b []float64) float64 {
	result := float64(0)
	for i := range a {
		result += a[i] * b[i]
	}
	return result
}

func Sub(a, b []float64) []float64 {
	result := make([]float64, len(a))
	for i := range a {
		result[i] = a[i] - b[i]
	}
	return result
}

func Rbf(gamma float64) func([]float64, []float64) float64 {
	return func(a, b []float64) float64 {
		deltaRow := Sub(a, b)
		deltaRowSquared := Dot(deltaRow, deltaRow)
		return math.Exp(deltaRowSquared / (-1.0 * math.Pow(gamma, 2)))
	}
}

func KernelTrans(matrix [][]float64, vec []float64, kernel func(x, y []float64) float64) []float64 {
	m := len(matrix)
	K := make([]float64, m)
	for i := range matrix {
		K[i] = kernel(matrix[i], vec)
	}
	return K
}
