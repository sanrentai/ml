package main

import (
	"fmt"
	"math"

	"github.com/sanrentai/ml"
)

func main2() {

	datMat, classLabels := loadSimpData()
	classifierArray := adaBosstTrainDS(datMat, classLabels, 9)
	fmt.Println(classifierArray)
}

// DS 代表 Decision Stump 单层决策树，最流行的弱分类器
func adaBosstTrainDS(dataArr [][]float64, classLabels []float64, numIt int) []*AdaStump {
	weakClassArr := make([]*AdaStump, 0)
	m := len(dataArr)
	D := make([]float64, m)
	for i := range D {
		D[i] = 1.0 / float64(m)
	}
	aggClassEst := make([]float64, m)
	for j := 0; j < numIt; j++ {
		bestStump, err, classEst := buildStump(dataArr, classLabels, D)
		// fmt.Printf("D:%v\n", D)
		alpha := 0.5 * math.Log((1.0-err)/math.Max(err, 1e-16))
		bestStump.Alpha = alpha
		weakClassArr = append(weakClassArr, bestStump)
		// fmt.Printf("classEst:%v\n", classEst)

		expon := ml.VecMul(ml.VecProd(-alpha, classLabels), classEst)
		expon = ml.VecExp(expon)

		for i := range D {
			D[i] = D[i] * expon[i]
		}
		D = ml.VecNormalize(D)
		for i := range aggClassEst {
			aggClassEst[i] += alpha * classEst[i]
		}
		// fmt.Printf("aggClassEst:%v\n", aggClassEst)
		aggErrors := 0.0
		for i := range aggClassEst {
			if ml.Sign(aggClassEst[i]) != classLabels[i] {
				aggErrors += 1.0
			}
		}
		errorRate := aggErrors / float64(m)
		fmt.Printf("total error:%v\n", errorRate)
		if errorRate == 0.0 {
			break
		}
	}

	return weakClassArr
}
