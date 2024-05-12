package ml

// 按给定特征划分数据集 数据集 特征 特征值
func SplitDataSet(dataSet [][]any, axis int, value any) [][]any {
	retDataSet := [][]any{}
	for _, featVec := range dataSet {
		if featVec[axis] == value {
			reducedFeatVec := append([]any{}, featVec[:axis]...) // chop out axis used for splitting
			reducedFeatVec = append(reducedFeatVec, featVec[axis+1:]...)
			retDataSet = append(retDataSet, reducedFeatVec)
		}
	}
	return retDataSet
}

// 选择最好的数据集划分方式
func ChooseBestFeatureToSplit(dataSet [][]any) int {
	numFeatures := len(dataSet[0]) - 1
	baseEntropy := CalcShannonEnt(dataSet)
	bestInfoGain := 0.0
	bestFeature := -1
	for i := 0; i < numFeatures; i++ {
		featList := make([]any, len(dataSet))
		for j, example := range dataSet {
			featList[j] = example[i]
		}
		uniqueVals := make(map[any]bool)
		for _, val := range featList {
			uniqueVals[val] = true
		}
		newEntropy := 0.0
		for value := range uniqueVals {
			subDataSet := SplitDataSet(dataSet, i, value)
			prob := float64(len(subDataSet)) / float64(len(dataSet))
			newEntropy += prob * CalcShannonEnt(subDataSet)
		}
		infoGain := baseEntropy - newEntropy
		if infoGain > bestInfoGain {
			bestInfoGain = infoGain
			bestFeature = i
		}
	}
	return bestFeature
}
