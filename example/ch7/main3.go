package main

import (
	"fmt"

	"github.com/sanrentai/ml"
)

func main3() {
	datMat, classLabels := loadSimpData()
	classifierArray := adaBosstTrainDS(datMat, classLabels, 9)
	fmt.Println(adaClassify([][]float64{
		{0, 0},
	}, classifierArray))

	fmt.Println(adaClassify([][]float64{
		{5, 5},
		{0, 0},
	}, classifierArray))
}

func adaClassify(data [][]float64, classifierArr []*AdaStump) []float64 {
	m := len(data)
	aggClassEst := make([]float64, m)
	for i := range classifierArr {
		classEst := stumpClassify(data, classifierArr[i].Dim, classifierArr[i].Thresh, classifierArr[i].Ineq)
		for j := range aggClassEst {
			aggClassEst[j] += classifierArr[i].Alpha * classEst[j]
		}
		// fmt.Println(aggClassEst)
	}
	return ml.SignVec(aggClassEst)
}
