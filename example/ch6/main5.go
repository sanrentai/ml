package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/sanrentai/ml"
)

func main5() {
	// testDigits(ml.Rbf(0.1))
	// 支持向量个数:1934
	// 训练集错误率：0.00%
	// 测试集错误率：9.41%
	// 耗时：9m14.6070452s
	// testDigits(ml.Rbf(10))
	// 支持向量个数:299
	// 训练集错误率：0.00%
	// 测试集错误率：0.63%
	// 耗时：3m17.3254017s
	// testDigits(ml.Rbf(20))
	// 支持向量个数:157
	// 训练集错误率：0.00%
	// 测试集错误率：1.27%
	// 耗时：4m58.950615s
	// testDigits(ml.Rbf(50))
	// 支持向量个数:127
	// 训练集错误率：1.40%
	// 测试集错误率：2.43%
	// 耗时：2m15.9800277s
	// testDigits(ml.Rbf(100))
	// 支持向量个数:149
	// 训练集错误率：1.29%
	// 测试集错误率：3.28%
	// 耗时：4m24.800389s
	testDigits(ml.Dot)
	// 支持向量个数:136
	// 训练集错误率：5.48%
	// 测试集错误率：4.97%
	// 耗时：3m15.4057284s
}

// rbf 20

func testDigits(kernel func([]float64, []float64) float64) {
	t := time.Now()
	dataArr, labelArr := loadImage("trainingDigits")
	b, alphas := smoP2(dataArr, labelArr, 200, 0.001, 10000, kernel)

	svInd := make([]int, 0)
	svAlphas := make([]float64, 0)
	for i := range alphas {
		if alphas[i] > 0 {
			svInd = append(svInd, i)
			svAlphas = append(svAlphas, alphas[i])
		}
	}

	sVs := make([][]float64, len(svInd))
	labelSV := make([]float64, len(svInd))
	for i := range svInd {
		sVs[i] = dataArr[svInd[i]]
		labelSV[i] = labelArr[svInd[i]]
	}

	fmt.Printf("支持向量个数:%d\n", len(svInd))

	m := len(dataArr)
	errCount := 0

	for i := 0; i < m; i++ {
		kernelEval := ml.KernelTrans(sVs, dataArr[i], kernel)
		predict := 0.0
		for j := 0; j < len(svInd); j++ {
			predict += kernelEval[j] * labelSV[j] * svAlphas[j]
		}
		predict += b
		if math.Signbit(predict) != math.Signbit(labelArr[i]) {
			errCount++
		}
	}

	fmt.Printf("训练集错误率：%.2f%%\n", float64(errCount)/float64(m)*100)

	dataArr, labelArr = loadImage("testDigits")
	errorCount := 0
	m = len(dataArr)
	for i := 0; i < m; i++ {
		kernelEval := ml.KernelTrans(sVs, dataArr[i], kernel)
		predict := 0.0
		for j := 0; j < len(svInd); j++ {
			predict += kernelEval[j] * labelSV[j] * svAlphas[j]
		}
		predict += b
		if math.Signbit(predict) != math.Signbit(labelArr[i]) {
			errorCount++
		}
	}
	fmt.Printf("测试集错误率：%.2f%%\n", float64(errorCount)/float64(m)*100)
	fmt.Printf("耗时：%v\n", time.Since(t))
}

// img2vector converts a 32x32 binary image to a 1x1024 vector.
func img2vector(filename string) []float64 {
	// Create a new 1x1024 slice filled with zeros
	returnVect := make([]float64, 1024)

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	// Initialize a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	for i := 0; scanner.Scan(); i++ {
		// Read a line from the file
		lineStr := scanner.Text()
		// Convert the line string to an integer slice
		lineInt := make([]int, 32)
		for j := range lineInt {
			lineInt[j], _ = strconv.Atoi(lineStr[j : j+1]) // Convert each character to an integer
		}
		// Add the integers to the return vector
		for j := range lineInt {
			returnVect[32*i+j] = float64(lineInt[j])
		}
	}

	// Close the scanner and return the vector
	return returnVect
}

func loadImage(dirName string) ([][]float64, []float64) {
	hwLabels := []float64{}

	// List the files in the directory
	fileList, err := filepath.Glob(filepath.Join(dirName, "*.txt"))
	if err != nil {
		fmt.Println("Error listing files:", err)
		return nil, nil
	}
	m := len(fileList)

	// Initialize a matrix to store the vectors
	trainingMat := make([][]float64, m)

	// Process each file
	for i, filePath := range fileList {
		fileName := filepath.Base(filePath)
		fileStr := strings.TrimSuffix(fileName, filepath.Ext(fileName))
		classNumStr := strings.Split(fileStr, "_")[0]
		classNum, err := strconv.Atoi(classNumStr)
		if err != nil {
			fmt.Println("Error parsing class number:", err)
			return nil, nil
		}
		if classNum == 9 {
			hwLabels = append(hwLabels, -1)
		} else {
			hwLabels = append(hwLabels, 1)
		}
		trainingMat[i] = img2vector(filePath)
	}

	// Return the matrix and labels
	return trainingMat, hwLabels
}
