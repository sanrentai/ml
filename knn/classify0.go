package knn

import (
	"github.com/sanrentai/ml"
)

// k-近邻算法
func Classify0(inX []float64, dataSet [][]float64, labels []string, k int) string {
	distances := make([]float64, len(dataSet))

	// 计算输入向量与每个样本的欧式距离
	for i, data := range dataSet {
		distances[i] = ml.EuclideanDistance(inX, data)
	}

	// 将距离排序并返回最近的 k 个点的索引
	sortedDistIndicies := ml.Argsort(distances)
	// 统计最近的 k 个点的类别标签
	classCount := make(map[string]int)
	for i := 0; i < k; i++ {
		voteLabel := labels[sortedDistIndicies[i]]
		classCount[voteLabel]++
	}

	// 找出票数最多的类别标签
	var maxCount int
	var maxLabel string
	for label, count := range classCount {
		if count > maxCount {
			maxCount = count
			maxLabel = label
		}
	}

	return maxLabel
}
