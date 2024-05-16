package main

import (
	"fmt"
	"math"
	"time"

	"github.com/sanrentai/ml"
)

func main3() {
	t := time.Now()
	// Example usage
	dataMat, labelMat := loadDataSet("testSet.txt")
	// fmt.Println("Data Matrix:", dataMat)
	// fmt.Println("Label Matrix:", labelMat)

	b, alphas := smoP(dataMat, labelMat, 0.6, 0.001, 40)
	fmt.Println("b:", b)
	showVec(alphas, 0.0)

	for i := 0; i < len(alphas); i++ {
		if alphas[i] > 0.0 {
			fmt.Printf("%v, %v\n", dataMat[i], labelMat[i])
		}
	}

	w := getW(dataMat, labelMat, alphas)
	fmt.Println("w:", w)
	showClassifier(dataMat, w, b, alphas, labelMat)
	fmt.Println("time:", time.Since(t))
}

type optStruct struct {
	X        [][]float64 // 数据矩阵
	labelMat []float64   // 数据标签
	C        float64     // 松弛变量
	tol      float64     // 容错率
	m        int         // 数据矩阵行数
	alphas   []float64   // alpha参数
	b        float64     // b参数
	eCache   [][]float64 // 误差缓存
	K        [][]float64
}

func (os *optStruct) calcEk(k int) float64 {
	// 计算预测值
	fk := float64(0)
	for i := 0; i < os.m; i++ {
		fk += os.alphas[i] * os.labelMat[i] * kernelTrans(os.X[i], os.X[k])
	}
	fk += os.b

	// 计算误差
	Ek := fk - os.labelMat[k]

	// 更新误差缓存
	// os.eCache[k] = []float64{1, Ek}

	return Ek
}

func smoP(dataMatIn [][]float64, classLabelsIn []float64, C, toler float64, maxIter int) (float64, []float64) {

	oS := newOptStruct(dataMatIn, classLabelsIn, C, toler, ml.Dot)
	iter := 0
	entireSet := true
	alphaPairsChanged := 0

	for iter < maxIter && (alphaPairsChanged > 0 || entireSet) {
		alphaPairsChanged = 0
		if entireSet {
			// 遍历所有样本
			for i := 0; i < oS.m; i++ {
				alphaPairsChanged += oS.innerL(i)
				fmt.Printf("fullSet, iter: %d i:%d, pairs changed %d\n", iter, i, alphaPairsChanged)
			}
			iter += 1
		} else {
			// 遍历非边界样本
			nonBoundIs := []int{}
			for i := 0; i < oS.m; i++ {
				if oS.alphas[i] > 0 && oS.alphas[i] < C {
					nonBoundIs = append(nonBoundIs, i)
				}
			}
			for i := 0; i < len(nonBoundIs); i++ {
				alphaPairsChanged += oS.innerL(nonBoundIs[i])
				fmt.Printf("non-bound, iter: %d i:%d, pairs changed %d\n", iter, nonBoundIs[i], alphaPairsChanged)
			}
			iter += 1
		}
		if entireSet {
			entireSet = false
		} else if alphaPairsChanged == 0 {
			entireSet = true
		}
		fmt.Printf("iteration number: %d\n", iter)
	}
	return oS.b, oS.alphas
}

// innerL 执行SMO算法中的内循环
func (oS *optStruct) innerL(i int) int {
	Ei := oS.calcEk(i)
	if ((oS.labelMat[i]*Ei < -oS.tol) && (oS.alphas[i] < oS.C)) ||
		((oS.labelMat[i]*Ei > oS.tol) && (oS.alphas[i] > 0)) {
		j, Ej := oS.selectJ(i, Ei)
		alphaIold := oS.alphas[i]
		alphaJold := oS.alphas[j]
		var L, H float64
		if oS.labelMat[i] != oS.labelMat[j] {
			L = math.Max(0, oS.alphas[j]-oS.alphas[i])
			H = math.Min(oS.C, oS.C+oS.alphas[j]-oS.alphas[i])
		} else {
			L = math.Max(0, oS.alphas[j]+oS.alphas[i]-oS.C)
			H = math.Min(oS.C, oS.alphas[j]+oS.alphas[i])
		}

		if L == H {
			fmt.Println("L==H")
			return 0
		}

		eta := 2.0*kernelTrans(oS.X[i], oS.X[j]) - kernelTrans(oS.X[i], oS.X[i]) - kernelTrans(oS.X[j], oS.X[j])
		if eta >= 0 {
			fmt.Println("eta>=0")
			return 0
		}

		oS.alphas[j] -= oS.labelMat[j] * (Ei - Ej) / eta
		oS.alphas[j] = clipAlpha(oS.alphas[j], H, L)
		oS.updateEk(j)
		if math.Abs(oS.alphas[j]-alphaJold) < 0.00001 {
			fmt.Println("j not moving enough")
			return 0
		}

		oS.alphas[i] += oS.labelMat[j] * oS.labelMat[i] * (alphaJold - oS.alphas[j])
		oS.updateEk(i)

		b1 := oS.b - Ei - oS.labelMat[i]*(oS.alphas[i]-alphaIold)*kernelTrans(oS.X[i], oS.X[i]) - oS.labelMat[j]*(oS.alphas[j]-alphaJold)*kernelTrans(oS.X[i], oS.X[j])
		b2 := oS.b - Ej - oS.labelMat[i]*(oS.alphas[i]-alphaIold)*kernelTrans(oS.X[i], oS.X[j]) - oS.labelMat[j]*(oS.alphas[j]-alphaJold)*kernelTrans(oS.X[j], oS.X[j])
		if 0 < oS.alphas[i] && oS.alphas[i] < oS.C {
			oS.b = b1
		} else if 0 < oS.alphas[j] && oS.alphas[j] < oS.C {
			oS.b = b2
		} else {
			oS.b = (b1 + b2) / 2
		}
		return 1
	}
	return 0
}

// selectJ 选择第二个变量alpha_j
func (oS *optStruct) selectJ(i int, Ei float64) (int, float64) {
	// rand.Seed(time.Now().UnixNano())
	maxK := len(oS.alphas) - 1 // python 这里是-1
	maxDeltaE := 0.0
	Ej := 0.0
	oS.eCache[i] = []float64{1, Ei} // 更新误差缓存

	// 获取所有非零误差的索引
	validEcacheList := make([]int, 0)
	for k, entry := range oS.eCache {
		if entry[0] == 1 {
			validEcacheList = append(validEcacheList, k)
		}
	}

	if len(validEcacheList) > 1 {
		for _, k := range validEcacheList {
			if k == i {
				continue
			}
			Ek := oS.calcEk(k)
			deltaE := math.Abs(Ei - Ek)
			if deltaE > maxDeltaE { // 选择具有最大|Ei - Ek|的j
				maxK = k
				maxDeltaE = deltaE
				Ej = Ek
			}
		}
		return maxK, Ej
	} else {

		j := selectJrand(i, oS.m)
		Ej = oS.calcEk(j)
		return j, Ej
	}
}

func (oS *optStruct) updateEk(k int) {
	Ek := oS.calcEk(k)
	oS.eCache[k] = []float64{1, Ek}
}

func newOptStruct(dataMatIn [][]float64, classLabels []float64, C, toler float64, kernel func([]float64, []float64) float64) *optStruct {
	m := len(dataMatIn)
	alphas := make([]float64, m)
	eCache := make([][]float64, m)
	for i := range eCache {
		eCache[i] = make([]float64, 2)
	}
	K := make([][]float64, m)
	for i := range dataMatIn {
		K[i] = make([]float64, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			K[j][i] = ml.KernelTrans(dataMatIn, dataMatIn[i], kernel)[j]
		}
	}

	return &optStruct{
		X:        dataMatIn,
		labelMat: classLabels,
		C:        C,
		tol:      toler,
		m:        m,
		alphas:   alphas,
		b:        0,
		eCache:   eCache,
		K:        K,
	}
}
