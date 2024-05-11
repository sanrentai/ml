package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// file2matrix 从文件中读取数据并生成特征矩阵和标签向量
func File2matrix(filename string) ([][]float64, []int) {
	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer file.Close()

	// 读取文件内容
	var lines [][]float64
	var labels []int
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

		label, err := strconv.Atoi(fields[3])
		if err != nil {
			fmt.Println("Error parsing label:", err)
			return nil, nil
		}
		labels = append(labels, label)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, nil
	}

	return lines, labels
}
