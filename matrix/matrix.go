package matrix

import (
	"math"
	"math/rand"
	"time"
)

type Matrix [][]float64

func Mat(a []float64) Matrix {
	return Matrix{
		a,
	}
}

func New1(rows, cols int) Matrix {
	matrix := make([][]float64, rows)
	for i := range matrix {
		matrix[i] = make([]float64, cols)
	}
	// 填充矩阵
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = 1
		}
	}
	return matrix
}

// 生成一个随机的数组
func Rand(rows, cols int) Matrix {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 创建一个rows x cols的切片
	matrix := make([][]float64, rows)
	for i := range matrix {
		matrix[i] = make([]float64, cols)
	}

	// 填充矩阵
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = rand.Float64()
		}
	}

	return matrix
}

// Transpose 返回矩阵的转置
// 矩阵的转置是指将矩阵的行和列互换的操作。
// 如果一个矩阵是由 m 行 n 列的元素组成的，
// 那么它的转置矩阵就是由 n 行 m 列的元素组成的，
// 其中原矩阵的第 i 行第 j 列的元素会变成转置矩阵的第 j 行第 i 列的元素。
func (m Matrix) Transpose() Matrix {
	rows := len(m)
	cols := len(m[0])

	transposed := make(Matrix, cols)
	for i := 0; i < cols; i++ {
		transposed[i] = make([]float64, rows)
		for j := 0; j < rows; j++ {
			transposed[i][j] = m[j][i]
		}
	}

	return transposed
}

// Determinant 计算矩阵的行列式值
// 矩阵的行列式值是一个与矩阵相关的标量值,它在线性代数中具有重要的意义
// 行列式值可以用来衡量矩阵的性质，例如矩阵是否可逆、矩阵表示的线性变换是否保持了区域的面积（或体积）等。
// 对于一个 n×n 的方阵（即行数等于列数的矩阵），其行列式值的计算涉及到矩阵中元素的排列组合。
// 一种常见的方法是利用代数余子式的定义，递归地计算行列式值。
// 行列式值通常用符号 "det" 表示，例如对于一个矩阵 A，其行列式值表示为 det(A) 或 |A|。
// 行列式值的性质包括：
// 1 如果行列式值等于零，则矩阵为奇异矩阵（不可逆）；如果行列式值不等于零，则矩阵可逆。
// 2 行列式值与矩阵的转置无关，即 det(A) = det(A^T)。
// 3 行列式值的计算方法可以是通过行变换、列变换、或是对角线上的元素相乘再相减等方法。
func (m Matrix) Determinant() float64 {
	n := len(m)
	if n == 0 || len(m[0]) != n {
		panic("Invalid matrix dimensions")
	}

	if n == 1 {
		return m[0][0]
	}

	if n == 2 {
		return m[0][0]*m[1][1] - m[0][1]*m[1][0]
	}

	var det float64
	for i := 0; i < n; i++ {
		subMatrix := make(Matrix, n-1)
		for j := 0; j < n-1; j++ {
			subMatrix[j] = make([]float64, n-1)
			copy(subMatrix[j][:i], m[j+1][:i])
			copy(subMatrix[j][i:], m[j+1][i+1:])
		}
		sign := 1.0
		if i%2 == 1 {
			sign = -1.0
		}
		det += sign * m[0][i] * subMatrix.Determinant()
	}
	return det
}

func (m Matrix) I() Matrix {
	return m.Inverse()
}

func (mat Matrix) Shape() (int, int) {
	m := len(mat)
	if m == 0 {
		panic("empty array")
	}
	n := len(mat[0])
	return m, n
}

// Inverse 返回矩阵的逆矩阵
func (m Matrix) Inverse() Matrix {
	det := m.Determinant()
	if math.Abs(det) < 1e-10 {
		panic("Matrix is singular, cannot compute inverse")
	}

	n := len(m)
	adjugate := make(Matrix, n)
	for i := range adjugate {
		adjugate[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			subMatrix := make(Matrix, n-1)
			for k := 0; k < n-1; k++ {
				subMatrix[k] = make([]float64, n-1)
				for l := 0; l < n-1; l++ {
					row := k
					if k >= i {
						row++
					}
					col := l
					if l >= j {
						col++
					}
					subMatrix[k][l] = m[row][col]
				}
			}
			sign := 1.0
			if (i+j)%2 == 1 {
				sign = -1.0
			}
			adjugate[j][i] = sign * subMatrix.Determinant()
		}
	}

	for i := range adjugate {
		for j := range adjugate[i] {
			adjugate[i][j] /= det
		}
	}

	return adjugate
}
