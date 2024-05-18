package ml

func Mat(m, n int) [][]float64 {
	r := make([][]float64, m)
	for i := range r {
		r[i] = make([]float64, n)
	}
	return r
}

func OnesMat(m, n int) [][]float64 {
	r := Mat(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			r[i][j] = 1
		}
	}
	return r
}

func ValProd(a, b [][]float64) [][]float64 {
	// a[m][n] * b[m][n] => c[m][n]
	m, n := len(a), len(a[0])
	c := Mat(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			c[i][j] = a[i][j] * b[i][j]
		}
	}
	return c
}

// 矩阵点乘
func MatDot(a, b [][]float64) [][]float64 {
	// a[m][n] * b[n][p] => c[m][p]
	m, n, p := len(a), len(a[0]), len(b[0])
	c := Mat(m, p)
	for i := 0; i < m; i++ {
		for j := 0; j < p; j++ {
			for k := 0; k < n; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}

func Col(mat [][]float64, col int) []float64 {
	r := make([]float64, len(mat))
	for i := 0; i < len(mat); i++ {
		r[i] = mat[i][col]
	}
	return r
}

func Row(mat [][]float64, row int) []float64 {
	return mat[row]
}
