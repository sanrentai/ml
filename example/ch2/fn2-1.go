package main

import (
	"fmt"

	"github.com/sanrentai/ml/knn"
)

func fn21() {
	group, labels := createDataSet()
	fmt.Println(group)
	fmt.Println(labels)

	inX := []float64{0, 0}
	label := knn.Classify(inX, group, labels, 3)
	fmt.Println(label)
}

func createDataSet() ([][]float64, []string) {
	group := [][]float64{
		{1.0, 1.1},
		{1.0, 1.0},
		{0, 0},
		{0, 0.1},
	}

	labels := []string{"A", "A", "B", "B"}

	return group, labels
}
