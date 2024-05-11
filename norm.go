package ml

import "math"

// AutoNorm 对数据集进行归一化处理
func AutoNorm(dataSet [][]float64) ([][]float64, []float64, []float64) {
	// 获取数据集的最小值和最大值
	minVals, maxVals := MinMax(dataSet)

	// 计算范围
	ranges := make([]float64, len(minVals))
	for i := range minVals {
		ranges[i] = maxVals[i] - minVals[i]
	}

	// 归一化数据集
	normDataSet := make([][]float64, len(dataSet))
	for i, data := range dataSet {
		normDataSet[i] = make([]float64, len(data))
		for j, val := range data {
			normDataSet[i][j] = (val - minVals[j]) / ranges[j]
		}
	}

	return normDataSet, ranges, minVals
}

// MinMax 返回数据集中每个特征的最小值和最大值
func MinMax(dataSet [][]float64) ([]float64, []float64) {
	numFeatures := len(dataSet[0])
	minVals := make([]float64, numFeatures)
	maxVals := make([]float64, numFeatures)
	for i := range minVals {
		minVals[i] = math.Inf(1)
		maxVals[i] = math.Inf(-1)
	}
	for _, data := range dataSet {
		for j, val := range data {
			if val < minVals[j] {
				minVals[j] = val
			}
			if val > maxVals[j] {
				maxVals[j] = val
			}
		}
	}
	return minVals, maxVals
}
