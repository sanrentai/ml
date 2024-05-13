package main

import (
	"fmt"

	"github.com/sanrentai/ml/tree"
)

func fn38() {
	myDat, labels := createDataSet()
	fmt.Println(myDat)
	fmt.Println(labels)
	myTree := retrieveTree(0)
	fmt.Println(myTree)
	fmt.Println(classify(myTree, labels, []interface{}{1, 0}))
	fmt.Println(classify(myTree, labels, []interface{}{1, 1}))
}

func classify(inputTree *tree.Tree, featureLabels []string, testVec []interface{}) interface{} {
	firstStr := inputTree.Value.(string)

	secondDict := inputTree.Children
	featIndex := 0
	for i, label := range featureLabels {
		if label == firstStr {
			featIndex = i
			break
		}
	}
	var classLabel interface{}
	for i, val := range secondDict {
		if testVec[featIndex] == i {
			if len(val.Children) == 0 {
				classLabel = val.Value
			} else {
				// 递归调用
				classLabel = classify(val, featureLabels[1:], testVec[1:])
			}
		}
	}
	return classLabel
}
