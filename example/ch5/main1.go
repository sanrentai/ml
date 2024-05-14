package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sanrentai/ml/matrix"
	"gonum.org/v1/gonum/mat"
)

func main1() {
	dataArr, labelMat := loadDataSet()

	fmt.Println(gradAscent(dataArr, labelMat))
	fmt.Println(gradAscent2(dataArr, labelMat))
	fmt.Println(stocGradAscent0(dataArr, labelMat))
}

// loadDataSet 函数用于加载数据集
func loadDataSet() ([][]float64, []float64) {
	var dataMat [][]float64
	var labelMat []float64

	file, err := os.Open("testSet.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineArr := strings.Fields(line)
		x0, _ := strconv.ParseFloat(lineArr[0], 64)
		x1, _ := strconv.ParseFloat(lineArr[1], 64)
		label, _ := strconv.ParseFloat(lineArr[2], 64)
		dataMat = append(dataMat, []float64{1.0, x0, x1})
		labelMat = append(labelMat, label)
	}

	return dataMat, labelMat
}

// sigmoid 函数用于计算 sigmoid 函数值
func sigmoid(inX float64) float64 {
	return 1.0 / (1 + math.Exp(-inX))
}

// gradAscent 函数用于梯度上升法训练逻辑回归模型
func gradAscent(dataMatIn [][]float64, classLabels []float64) []float64 {
	t := time.Now()
	dataMatrix := mat.NewDense(len(dataMatIn), len(dataMatIn[0]), nil)
	for i, row := range dataMatIn {
		for j, val := range row {
			dataMatrix.Set(i, j, val)
		}
	}
	labelMat := mat.NewDense(len(classLabels), 1, nil)
	for i, label := range classLabels {
		labelMat.Set(i, 0, label)
	}
	_, n := dataMatrix.Dims()
	alpha := 0.001
	maxCycles := 500
	weights := mat.NewDense(n, 1, nil)
	weights.Apply(func(_, _ int, _ float64) float64 { return 1 }, weights) // 初始化为全 1 向量

	for k := 0; k < maxCycles; k++ {
		var h mat.Dense
		h.Mul(dataMatrix, weights)
		h.Apply(func(_, _ int, v float64) float64 {
			return sigmoid(v)
		}, &h)

		var errorMat mat.Dense
		errorMat.Sub(labelMat, &h)

		var deltaWeights mat.Dense
		deltaWeights.Mul(dataMatrix.T(), &errorMat)
		deltaWeights.Scale(alpha, &deltaWeights)

		weights.Add(weights, &deltaWeights)
	}

	weightsSlice := make([]float64, n)
	for i := 0; i < n; i++ {
		weightsSlice[i] = weights.At(i, 0)
	}
	fmt.Println(time.Since(t))
	return weightsSlice

}

// vecDotProduct 函数用于计算向量点乘
func vecDotProduct(vec1 []float64, vec2 []float64) float64 {
	result := 0.0
	for i := 0; i < len(vec1); i++ {
		result += vec1[i] * vec2[i]
	}
	return result
}

// gradAscent 函数用于梯度上升法训练逻辑回归模型
func gradAscent2(dataMatIn [][]float64, classLabels []float64) matrix.Matrix {
	t := time.Now()
	dataMatrix := matrix.Matrix(dataMatIn)
	labelMat := matrix.Matrix{classLabels}.Transpose()
	_, n := dataMatrix.Shape()
	alpha := 0.001
	maxCycles := 500
	weights := matrix.New1(n, 1)
	for cycle := 0; cycle < maxCycles; cycle++ {

		h := dataMatrix.Mul(weights).Sigmoid()
		errorTerm := matrix.Sub(labelMat, h)
		// fmt.Println(errorTerm)
		weights = matrix.Add(weights, dataMatrix.Transpose().MulWithAlpha(errorTerm, alpha))
		// fmt.Println(weights)
	}
	fmt.Println(time.Since(t))
	return weights
}
