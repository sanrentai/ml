package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main4() {
	datArr, labelArr := loadDataSet("horseColicTraining2.txt")
	calssifierArray, _ := adaBosstTrainDS(datArr, labelArr, 10)
	// fmt.Println(calssifierArray)
	testArr, testLabelArr := loadDataSet("horseColicTest2.txt")
	prediction10 := adaClassify(testArr, calssifierArray)
	errNum := 0.0
	for i := range testLabelArr {
		if prediction10[i] != testLabelArr[i] {
			errNum++
		}
	}
	fmt.Println(errNum / float64(len(testLabelArr)))
}

func loadDataSet(filename string) (dataMat [][]float64, labelMat []float64) {
	dataMat = make([][]float64, 0)
	labelMat = make([]float64, 0)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "\t")
		var length = len(fields)
		var lineArr = make([]float64, 0, length-1)
		for i, field := range fields {
			v, err := strconv.ParseFloat(field, 64)
			if err != nil {
				fmt.Println("Error ParseFloat :", err)
				return
			}
			if i < length-1 {
				lineArr = append(lineArr, v)
			} else {
				labelMat = append(labelMat, v)
			}
		}
		dataMat = append(dataMat, lineArr)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	return
}
