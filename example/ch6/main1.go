package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func loadDataSet(filename string) ([][]float64, []float64) {
	dataMat := make([][]float64, 0)
	labelMat := make([]float64, 0)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")
		if len(parts) != 3 {
			continue
		}
		x1, err1 := strconv.ParseFloat(parts[0], 64)
		x2, err2 := strconv.ParseFloat(parts[1], 64)
		y, err3 := strconv.ParseFloat(parts[2], 64)
		if err1 != nil || err2 != nil || err3 != nil {
			continue
		}
		dataMat = append(dataMat, []float64{x1, x2})
		labelMat = append(labelMat, y)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return dataMat, labelMat
}

// 辅助函数，用于剪辑alpha的值
func clipAlpha(aj float64, H, L float64) float64 {
	if aj > H {
		return H
	} else if aj < L {
		return L
	} else {
		return aj
	}
}

// 辅助函数，用于生成一个不等于i的随机数
func selectJrand(i, m int) int {
	j := i
	for j == i {
		j = rand.Intn(m)
	}
	return j
}

func main1() {
	// Example usage
	dataMat, labelMat := loadDataSet("testSet.txt")
	fmt.Println("Data Matrix:", dataMat)
	fmt.Println("Label Matrix:", labelMat)

	// Example of selecting a random index
	i := 5
	m := len(dataMat)
	j := selectJrand(i, m)
	fmt.Println("Random index not equal to", i, ":", j)

	// Example of clipping a value
	aj := 2.5
	H := 1.0
	L := 0.1
	clippedAj := clipAlpha(aj, H, L)
	fmt.Println("Clipped value:", clippedAj)
}
