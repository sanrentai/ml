package main

import (
	"fmt"

	"github.com/sanrentai/ml/tree"
)

func fn34() {
	myDat, labels := createDataSet()
	fmt.Println("Labels:", labels)
	fmt.Println("Data Set:", myDat)
	fmt.Println(tree.CreateTree(myDat, labels))
}
