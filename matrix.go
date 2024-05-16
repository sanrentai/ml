package ml

func Mat(m, n int) [][]float64 {
	r := make([][]float64, m)
	for i := range r {
		r[i] = make([]float64, n)
	}
	return r
}

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
