package matrix

import "math"

func Mul(a, b Matrix) Matrix {
	return Multiplication(a, b, 1.0)
}

func (a Matrix) MulWithAlpha(b Matrix, alpha float64) Matrix {
	return Multiplication(a, b, alpha)
}

func (a Matrix) Mul(b Matrix) Matrix {
	return Multiplication(a, b, 1.0)
}

func (a Matrix) Sigmoid() Matrix {
	m, n := a.Shape()
	result := make(Matrix, m)
	for i := range result {
		result[i] = make([]float64, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = 1.0 / (1.0 + math.Exp(-a[i][j]))
		}
	}
	return result
}

// 矩阵乘法是线性代数中的一个基本操作，用于将两个矩阵相乘得到一个新的矩阵。
// 在矩阵乘法中，两个矩阵的乘法不像普通的数乘，而是需要满足一定的条件才能进行。
// 具体来说，如果一个矩阵 A 的列数等于另一个矩阵 B 的行数，则可以对这两个矩阵进行乘法运算。

// 考虑两个矩阵 A 和 B，它们的维度分别为 m×n 和 n×p，那么它们的乘积 C 的维度将是 m×p。
// 矩阵 C 中的每个元素 c[i][j] 是由矩阵 A 的第 i 行和矩阵 B 的第 j 列对应元素相乘后再求和得到的。
func Multiplication(a, b Matrix, alpha float64) Matrix {
	rowsA := len(a)
	colsA := len(a[0])
	rowsB := len(b)
	colsB := len(b[0])

	if colsA != rowsB {
		panic("矩阵尺寸不匹配，无法进行乘法运算。")
	}

	result := make(Matrix, rowsA)
	for i := range result {
		result[i] = make([]float64, colsB)
	}

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += alpha * a[i][k] * b[k][j]
			}
		}
	}

	return result
}
