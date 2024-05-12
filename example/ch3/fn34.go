package main

import (
	"fmt"

	"github.com/sanrentai/ml"
)

func fn34() {
	myDat, labels := createDataSet()
	fmt.Println("Labels:", labels)
	fmt.Println("Data Set:", myDat)
	fmt.Println(ml.CreateTree(myDat, labels))
}
