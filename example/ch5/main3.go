package main

func main3() {
	dataArr, labelMat := loadDataSet()
	wei := stocGradAscent0(dataArr, labelMat)
	plotBestFit(wei)
}

func stocGradAscent0(dataMatrix [][]float64, classLabels []float64) []float64 {
	m := len(dataMatrix)
	n := len(dataMatrix[0])
	alpha := 0.01
	weights := make([]float64, n)
	// rand.Seed(time.Now().UnixNano()) // 使用当前时间作为随机数种子
	for j := 0; j < n; j++ {
		weights[j] = 1
	}
	maxCycles := 200 // 这里为什么是200
	for a := 0; a < maxCycles; a++ {

		for i := 0; i < m; i++ {
			h := 0.0
			for j := 0; j < n; j++ {
				h += dataMatrix[i][j] * weights[j]
			}
			h = sigmoid(h)

			errorMat := classLabels[i] - h
			for j := 0; j < n; j++ {
				weights[j] = weights[j] + alpha*errorMat*dataMatrix[i][j]
			}
		}
	}
	return weights
}
