package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Img2vector 将二进制图像转换为向量
func Img2vector(filename string) []int {
	var returnVect []int

	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	// 读取文件内容
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		for j := 0; j < 32; j++ {
			num, _ := strconv.Atoi(string(lineStr[j]))
			returnVect = append(returnVect, num)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return returnVect
}