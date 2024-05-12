package main

import (
	"fmt"

	"github.com/sanrentai/ml"
)

func fn33() {
	myDat, labels := createDataSet()
	fmt.Println("Labels:", labels)
	fmt.Println("Data Set:", myDat)
	fmt.Println("Best Feature:", chooseBestFeatureToSplit(myDat)) // 0
	// 这个结果第0个特征是最好的用于划分数据集的特征
}

func chooseBestFeatureToSplit(dataSet [][]any) int {
	numFeatures := len(dataSet[0]) - 1
	baseEntropy := ml.CalcShannonEnt(dataSet)
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
			subDataSet := splitDataSet(dataSet, i, value)
			prob := float64(len(subDataSet)) / float64(len(dataSet))
			newEntropy += prob * ml.CalcShannonEnt(subDataSet)
		}
		infoGain := baseEntropy - newEntropy
		if infoGain > bestInfoGain {
			bestInfoGain = infoGain
			bestFeature = i
		}
	}
	return bestFeature
}
