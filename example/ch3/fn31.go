package main

import (
	"fmt"

	"github.com/sanrentai/ml"
)

func fn31() {
	myDat, labels := createDataSet()
	fmt.Println("Labels:", labels)
	fmt.Println("Data Set:", myDat)

	shannonEnt := ml.CalcShannonEnt(myDat)
	fmt.Println("Shannon Entropy:", shannonEnt)
}

func createDataSet() ([][]interface{}, []string) {
	return [][]interface{}{
		{1, 1, "yes"},
		{1, 1, "yes"},
		{1, 0, "no"},
		{0, 1, "no"},
		{0, 1, "no"},
	}, []string{"no surfacing", "flippers"}
}
