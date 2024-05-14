package main

import (
	"math/rand"
	"time"
)

func main4() {
	dataArr, labelMat := loadDataSet()
	wei := stocGradAscent1(dataArr, labelMat, 500)
	plotBestFit(wei)
}

func stocGradAscent1(dataMatrix [][]float64, classLabels []float64, numIter int) []float64 {
	m := len(dataMatrix)
	n := len(dataMatrix[0])
	weights := make([]float64, n)
	rand.Seed(time.Now().UnixNano())

	for j := 0; j < numIter; j++ {
		dataIndex := rand.Perm(m)
		for i := 0; i < m; i++ {
			alpha := 4.0/(1.0+float64(j+i)) + 0.01
			randIndex := dataIndex[i]
			var h float64
			for k := 0; k < n; k++ {
				h += dataMatrix[randIndex][k] * weights[k]
			}
			h = sigmoid(h)
			error := classLabels[randIndex] - h
			for k := 0; k < n; k++ {
				weights[k] += alpha * error * dataMatrix[randIndex][k]
			}
		}
	}
	return weights
}
