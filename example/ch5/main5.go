package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main5() {
	multiTest()
}

// dotProduct 计算两个向量的点积
func dotProduct(vec1 []float64, vec2 []float64) float64 {
	result := 0.0
	for i := 0; i < len(vec1); i++ {
		result += vec1[i] * vec2[i]
	}
	return result
}

// classifyVector 分类函数
func classifyVector(inX []float64, weights []float64) float64 {
	// 计算输入向量与权重的点积，并通过sigmoid函数计算概率
	prob := sigmoid(dotProduct(inX, weights))
	// 如果概率大于0.5，则返回1.0，否则返回0.0
	if prob > 0.5 {
		return 1.0
	}
	return 0.0
}

func colicTest() float64 {
	frTrain, err := os.Open("horseColicTraining.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return 0.0
	}
	defer frTrain.Close()

	frTest, err := os.Open("horseColicTest.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return 0.0
	}
	defer frTest.Close()

	trainingSet := make([][]float64, 0)
	trainingLabels := make([]float64, 0)

	// 读取训练数据
	scanner := bufio.NewScanner(frTrain)
	for scanner.Scan() {
		line := scanner.Text()
		lineArr := strings.Split(line, "\t")
		dataArr := make([]float64, len(lineArr)-1)
		for i := 0; i < len(lineArr)-1; i++ {
			val, _ := strconv.ParseFloat(lineArr[i], 64)
			dataArr[i] = val
		}
		label, _ := strconv.ParseFloat(lineArr[len(lineArr)-1], 64)
		trainingSet = append(trainingSet, dataArr)
		trainingLabels = append(trainingLabels, label)
	}

	// 使用随机梯度上升算法进行训练
	trainWeights := stocGradAscent1(trainingSet, trainingLabels, 500)

	// 测试数据
	errorCount := 0
	numTestVec := 0.0
	scanner = bufio.NewScanner(frTest)
	for scanner.Scan() {
		numTestVec += 1.0
		line := scanner.Text()
		lineArr := strings.Split(line, "\t")
		dataArr := make([]float64, len(lineArr)-1)
		for i := 0; i < len(lineArr)-1; i++ {
			val, _ := strconv.ParseFloat(lineArr[i], 64)
			dataArr[i] = val
		}
		label, _ := strconv.ParseFloat(lineArr[len(lineArr)-1], 64)
		if int(classifyVector(dataArr, trainWeights)) != int(label) {
			errorCount++
		}
	}

	// 计算错误率
	errorRate := float64(errorCount) / numTestVec
	fmt.Printf("the error rate of this test is: %f\n", errorRate)
	return errorRate
}

// multiTest 函数
func multiTest() {
	numTests := 10
	errorSum := 0.0
	for k := 0; k < numTests; k++ {
		errorSum += colicTest()
	}
	avgErrorRate := errorSum / float64(numTests)
	fmt.Printf("after %d iterations the average error rate is: %f\n", numTests, avgErrorRate)
}
