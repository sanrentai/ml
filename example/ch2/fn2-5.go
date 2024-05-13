package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sanrentai/ml"
	"github.com/sanrentai/ml/knn"
)

func classifyPerson() {
	resultList := []string{"not at all", "in small doses", "in large doses"}

	// 获取用户输入
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("percentage of time spent playing video games? ")
	percentTatsStr, _ := reader.ReadString('\n')
	percentTats, _ := strconv.ParseFloat(strings.TrimSpace(percentTatsStr), 64)

	fmt.Print("frequent flier miles earned per year? ")
	ffMilesStr, _ := reader.ReadString('\n')
	ffMiles, _ := strconv.ParseFloat(strings.TrimSpace(ffMilesStr), 64)

	fmt.Print("liters of ice cream consumed per year? ")
	iceCreamStr, _ := reader.ReadString('\n')
	iceCream, _ := strconv.ParseFloat(strings.TrimSpace(iceCreamStr), 64)

	// 读取数据集并进行归一化处理
	datingDataMat, datingLabels := file2matrix("datingTestSet.txt")
	normMat, ranges, minVals := ml.AutoNorm(datingDataMat)

	// 使用用户输入的数据进行分类
	inArr := []float64{ffMiles, percentTats, iceCream}
	classifierResult := knn.Classify([]float64{(inArr[0] - minVals[0]) / ranges[0], (inArr[1] - minVals[1]) / ranges[1], (inArr[2] - minVals[2]) / ranges[2]}, normMat, datingLabels, 3)
	classifierResultInt, err := strconv.Atoi(classifierResult)
	if err != nil {
		panic(err)
	}
	fmt.Println("You will probably like this person:", resultList[classifierResultInt-1])
}
