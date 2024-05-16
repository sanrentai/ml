package main

import (
	"fmt"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
)

func main4() {
	dataArr, labelMat := loadDataSet()
	wei := stocGradAscent1(dataArr, labelMat, 500)
	plotBestFit(wei, "stocGradAscent1.png")
	wei2 := StochasticGradientAscent1(dataArr, labelMat, 500)
	plotBestFit(wei2.RawVector().Data, "StochasticGradientAscent1.png")
}

func stocGradAscent1(dataMatrix [][]float64, classLabels []float64, numIter int) []float64 {
	t := time.Now()
	m := len(dataMatrix)
	n := len(dataMatrix[0])
	weights := make([]float64, n)
	for i := 0; i < n; i++ {
		weights[i] = 1.0 // Initialize weights to ones
	}
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
			errnum := classLabels[randIndex] - h
			for k := 0; k < n; k++ {
				weights[k] += alpha * errnum * dataMatrix[randIndex][k]
			}
		}
	}
	fmt.Println(time.Since(t))
	return weights
}

func StochasticGradientAscent1(dataMatIn [][]float64, classLabels []float64, numIter int) *mat.VecDense {
	t := time.Now()

	dataMatrix := mat.NewDense(len(dataMatIn), len(dataMatIn[0]), nil)
	for i, row := range dataMatIn {
		for j, val := range row {
			dataMatrix.Set(i, j, val)
		}
	}
	m, n := dataMatrix.Dims()
	weights := mat.NewVecDense(n, nil)
	// Initialize weights to ones
	for i := 0; i < n; i++ {
		weights.SetVec(i, 1.0) // Initialize weights to ones
	}

	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	for j := 0; j < numIter; j++ {
		dataIndex := rand.Perm(m)
		// fmt.Println(dataIndex)
		for i := 0; i < m; i++ {
			alpha := 4.0/(1.0+float64(j+i)) + 0.01
			randIndex := dataIndex[i]
			h := sigmoid(mat.Dot(dataMatrix.RowView(randIndex), weights))
			errnum := classLabels[randIndex] - h
			weights.AddScaledVec(weights, alpha*errnum, dataMatrix.RowView(randIndex))
			// for k := 0; k < n; k++ {
			// 	weights.SetVec(k, weights.AtVec(k)+alpha*errnum*dataMatrix.At(randIndex, k))
			// }

		}
	}
	fmt.Println(time.Since(t))
	return weights
}
