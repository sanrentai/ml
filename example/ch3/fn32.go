package main

import (
	"fmt"
)

func fn32() {
	myDat, labels := createDataSet()
	fmt.Println("Labels:", labels)
	fmt.Println("Data Set:", myDat)

	fmt.Println(splitDataSet(myDat, 0, 1))

	fmt.Println(splitDataSet(myDat, 0, 0))
}

// 按给定特征划分数据集 数据集 特征 特征值
func splitDataSet(dataSet [][]any, axis int, value any) [][]any {
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
