package main

import (
	"fmt"
	"math"

	"github.com/sanrentai/ml"
)

func main1() {

	datMat, classLabels := loadSimpData()
	m := len(datMat)
	D := make([]float64, m)
	for i := range D {
		D[i] = 1.0 / float64(m)
	}
	bestStump, minError, bestClasEst := buildStump(datMat, classLabels, D)
	fmt.Println(bestStump)
	fmt.Println(minError)
	fmt.Println(bestClasEst)
}

func loadSimpData() ([][]float64, []float64) {
	return [][]float64{
			{1.0, 2.1},
			{2.0, 1.1},
			{1.3, 1.0},
			{1.0, 1.0},
			{2.0, 1.0},
		},
		[]float64{1.0, 1.0, -1.0, -1.0, 1.0}
}

func stumpClassify(dataMatrix [][]float64, dimen int, threshVal float64, threshIneq string) []float64 {
	retArray := make([]float64, len(dataMatrix))
	for i := range dataMatrix {

		if threshIneq == "lt" {
			if dataMatrix[i][dimen] <= threshVal {
				retArray[i] = -1
			} else {
				retArray[i] = 1
			}
		} else {
			if dataMatrix[i][dimen] > threshVal {
				retArray[i] = -1
			} else {
				retArray[i] = 1
			}
		}
	}
	return retArray
}

// 找到数据集上最佳的单层决策树
// dataArr - 数据矩阵
// classLabels - 数据标签
// D - 样本权重
func buildStump(dataArr [][]float64, classLabels []float64, D []float64) (*AdaStump, float64, []float64) {
	m := len(dataArr)
	n := len(dataArr[0])
	numSteps := 10.0
	bestStump := &AdaStump{}
	bestClasEst := make([]float64, m)
	minError := math.Inf(1)
	for i := 0; i < n; i++ {
		c := ml.Col(dataArr, i)
		rangeMin := ml.VecMin(c)
		rangeMax := ml.VecMax(c)
		stepSize := (rangeMax - rangeMin) / numSteps

		for j := -1; j <= int(numSteps); j++ {
			for _, inequal := range []string{"lt", "gt"} {
				threshVal := (rangeMin + float64(j)*stepSize)
				predictedVals := stumpClassify(dataArr, i, threshVal, inequal)
				errArr := make([]float64, m)
				for k, v := range classLabels {
					if v != predictedVals[k] {
						errArr[k] = 1.0
					}
				}
				weightedError := ml.Dot(D, errArr)
				// fmt.Printf("split: dim %d, thresh %.2f, thresh ineqal: %s, the weighted error is %.3f\n", i, threshVal, inequal, weightedError)
				if weightedError < minError {
					minError = weightedError
					for k := range bestClasEst {
						bestClasEst[k] = predictedVals[k]
					}
					bestStump.Dim = i
					bestStump.Thresh = threshVal
					bestStump.Ineq = inequal
				}
			}
		}
	}
	return bestStump, minError, bestClasEst
}

type AdaStump struct {
	Dim    int
	Thresh float64
	Ineq   string
	Alpha  float64
}

func (a AdaStump) String() string {
	return fmt.Sprintf(`{"dim": %d, "thresh": %v, "ineq": "%s", "alpha": %v}
 `, a.Dim, a.Thresh, a.Ineq, a.Alpha)
}
