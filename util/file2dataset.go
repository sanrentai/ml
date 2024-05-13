package util

import (
	"bufio"
	"os"
	"strings"
)

func File2Dataset(filename string) ([][]interface{}, error) {
	var data [][]interface{}
	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "\t")

		var row []interface{}
		for _, field := range fields {
			row = append(row, field)
		}

		data = append(data, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
