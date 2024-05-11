package matrix

func Sub(a, b Matrix) Matrix {
	return Subtraction(a, b)
}

// 矩阵减法函数
func Subtraction(a, b Matrix) Matrix {
	// 检查矩阵维度是否相同
	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		panic("矩阵维度不匹配")
	}

	// 创建结果矩阵
	result := make(Matrix, len(a))
	for i := range result {
		result[i] = make([]float64, len(a[0]))
	}

	// 执行矩阵减法
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			result[i][j] = a[i][j] - b[i][j]
		}
	}

	return result
}
