package ml

import "math"

// 计算给定数据集的香农熵  dataSet 每组数据长度一致，最后一个元素是分类标签
func CalcShannonEnt(dataSet [][]any) float64 {
	numEntries := len(dataSet)
	labelCounts := make(map[any]int)
	for _, featVec := range dataSet {
		currentLabel := featVec[len(featVec)-1]
		if _, ok := labelCounts[currentLabel]; !ok {
			labelCounts[currentLabel] = 0
		}
		labelCounts[currentLabel]++
	}
	shannonEnt := 0.0
	for _, count := range labelCounts {
		prob := float64(count) / float64(numEntries)
		shannonEnt -= prob * math.Log2(prob)
	}
	return shannonEnt
}
