package matrix

func Eye(n int) Matrix {
	eyeMatrix := make(Matrix, n)
	for i := range eyeMatrix {
		eyeMatrix[i] = make([]float64, n)
		eyeMatrix[i][i] = 1
	}
	return eyeMatrix
}
