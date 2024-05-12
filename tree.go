package ml

// TreeNode 表示决策树节点
type TreeNode struct {
	SplitFeature string
	Children     map[interface{}]*TreeNode
	LeafValue    interface{}
}

func CreateTree(dataSet [][]interface{}, labels []string) interface{} {
	classList := make([]interface{}, len(dataSet))
	for i, example := range dataSet {
		classList[i] = example[len(example)-1]
	}
	// 如果类别完全相同，则停止划分，返回该类别
	if countOccurrences(classList, classList[0]) == len(classList) {
		return classList[0]
	}
	if len(dataSet[0]) == 1 {
		return MajorityCnt(classList)
	}

	bestFeat := ChooseBestFeatureToSplit(dataSet)
	bestFeatLabel := labels[bestFeat]
	myTree := map[string][]interface{}{
		bestFeatLabel: {},
	}

	// 清空labels[bestFeat]
	labels = append(labels[:bestFeat], labels[bestFeat+1:]...)

	featValues := make([]interface{}, len(dataSet))
	for i, example := range dataSet {
		featValues[i] = example[bestFeat]
	}

	uniqueVals := set(featValues)

	for value := range uniqueVals {
		subLabels := make([]string, len(labels))
		copy(subLabels, labels)
		subDataSet := SplitDataSet(dataSet, bestFeat, value)
		myTree[bestFeatLabel] = append(myTree[bestFeatLabel], CreateTree(subDataSet, subLabels))
	}

	return myTree
}

// countOccurrences 计算列表中指定元素出现的次数
func countOccurrences(list []interface{}, elem interface{}) int {
	count := 0
	for _, e := range list {
		if e == elem {
			count++
		}
	}
	return count
}
