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

func DatingClass() {
	filename := "datingTestSet.txt" // 你的数据文件名

	// 调用 file2matrix 函数读取数据
	datingDataMat, datingLabels := file2matrix(filename)

	// 使用 hoRatio 分割数据集
	hoRatio := 0.10
	m := len(datingDataMat)
	numTestVecs := int(float64(m) * hoRatio)

	// 对数据集进行归一化处理
	normMat, _, _ := ml.AutoNorm(datingDataMat)

	errorCount := 0.0
	for i := 0; i < numTestVecs; i++ {
		classifierResult := knn.Classify(normMat[i], normMat[numTestVecs:m], datingLabels[numTestVecs:m], 3)
		fmt.Printf("the classifier came back with: %s, the real answer is: %s\n", classifierResult, datingLabels[i])
		if classifierResult != datingLabels[i] {
			errorCount += 1.0
		}
	}
	fmt.Printf("the total error rate is: %f\n", errorCount/float64(numTestVecs))
}

// file2matrix 从文件中读取数据并生成特征矩阵和标签向量
func file2matrix(filename string) ([][]float64, []string) {
	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer file.Close()

	// 读取文件内容
	var lines [][]float64
	var labels []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 4 {
			continue // 忽略格式错误的行
		}

		var data []float64
		for _, field := range fields[:3] {
			val, err := strconv.ParseFloat(field, 64)
			if err != nil {
				fmt.Println("Error parsing float:", err)
				return nil, nil
			}
			data = append(data, val)
		}
		lines = append(lines, data)

		// label, err := strconv.Atoi(fields[3])
		// if err != nil {
		// 	fmt.Println("Error parsing label:", err)
		// 	return nil, nil
		// }
		labels = append(labels, fields[3])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, nil
	}

	return lines, labels
}
