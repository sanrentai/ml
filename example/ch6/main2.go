package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main2() {
	t := time.Now()
	// Example usage
	dataMat, labelMat := loadDataSet("testSet.txt")
	// fmt.Println("Data Matrix:", dataMat)
	// fmt.Println("Label Matrix:", labelMat)

	b, alphas := smoSimple(dataMat, labelMat, 0.6, 0.001, 40)
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
	// 看书里写这个用时需要14.5秒，本地运行测试需要1.04秒
	// var xys1, xys2 plotter.XYs

	// for i, data := range dataMat {
	// 	if labelMat[i] > 1 {
	// 		xys1 = append(xys1, plotter.XY{
	// 			X: data[0],
	// 			Y: data[1],
	// 		})
	// 	} else {
	// 		xys2 = append(xys2, plotter.XY{
	// 			X: data[0],
	// 			Y: data[1],
	// 		})
	// 	}
	// }
	// // 创建绘图
	// p := plot.New()
	// s1, err := plotter.NewScatter(xys1)
	// if err != nil {
	// 	panic(err)
	// }
	// s1.GlyphStyle.Color = plotutil.Color(0)

	// s2, err := plotter.NewScatter(xys2)
	// if err != nil {
	// 	panic(err)
	// }
	// s1.GlyphStyle.Color = plotutil.Color(1)
	// p.Add(s1, s2)

	// w := getW(dataMat, labelMat, alphas)

	// // 绘制拟合直线
	// lineFunc := func(x float64) float64 {
	// 	return (-b - w[0]*x) / w[1]
	// }
	// ff := plotter.NewFunction(lineFunc)
	// // ff.XMin = -3
	// // ff.XMax = 3
	// p.Add(ff)
	// // p.X.Label.Text = "X1"
	// // p.Y.Label.Text = "X2"
	// // p.X.Max =
	// // p.X.Min = -4

	// // 保存图像
	// if err := p.Save(4*vg.Inch, 4*vg.Inch, "6-4.png"); err != nil {
	// 	panic(err)
	// }
}

func Max(matrix [][]float64) float64 {
	max := matrix[0][0]
	for _, row := range matrix {
		if len(row) == 0 {
			return 0
		}
		if row[0] > max {
			max = row[0]
		}
	}
	return max
}

func Min(matrix [][]float64) float64 {
	min := matrix[0][0]
	for _, row := range matrix {
		if len(row) == 0 {
			return 0
		}
		if row[0] < min {
			min = row[0]
		}
	}
	return min
}

// showClassifier 绘制分类器和支持向量
func showClassifier(dataMat [][]float64, w []float64, b float64, alphas []float64, labelMat []float64) {
	// 创建一个新的图表
	p := plot.New()

	p.Title.Text = "SVM Classifier"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// 分离正负样本
	dataPlus := make(plotter.XYs, 0)
	dataMinus := make(plotter.XYs, 0)
	for i := range dataMat {
		if labelMat[i] > 0 {
			dataPlus = append(dataPlus, plotter.XY{X: dataMat[i][0], Y: dataMat[i][1]})
		} else {
			dataMinus = append(dataMinus, plotter.XY{X: dataMat[i][0], Y: dataMat[i][1]})
		}
	}

	// 绘制正样本
	sPlus, err := plotter.NewScatter(dataPlus)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	sPlus.GlyphStyle.Color = color.RGBA{R: 255, G: 215, B: 0, A: 255}
	// sPlus.GlyphStyle.Radius = vg.Points(3)
	// sPlus.GlyphStyle.Shape = plotutil.Shape(0)
	p.Add(sPlus)

	// 绘制负样本
	sMinus, err := plotter.NewScatter(dataMinus)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	sMinus.GlyphStyle.Color = color.RGBA{R: 0, G: 0, B: 205, A: 255}
	// sMinus.GlyphStyle.Radius = vg.Points(3)
	// sMinus.GlyphStyle.Shape = plotutil.Shape(0)
	p.Add(sMinus)

	// 绘制支持向量
	sv := make(plotter.XYs, 0)
	for i, alpha := range alphas {
		if alpha > 0 {
			sv = append(sv, plotter.XY{X: dataMat[i][0], Y: dataMat[i][1]})
		}
	}
	svScatter, err := plotter.NewScatter(sv)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	svScatter.GlyphStyle.Color = color.RGBA{R: 255, A: 255}
	svScatter.GlyphStyle.Radius = vg.Points(5)
	svScatter.GlyphStyle.Shape = plotutil.Shape(1)
	p.Add(svScatter)

	// plotutil.AddScatters(p, "+", dataPlus, "-", dataMinus, sv)

	// 绘制分离超平面
	x1 := Min(dataMat)
	x2 := Max(dataMat)

	y1 := (-b - w[0]*x1) / w[1]
	y2 := (-b - w[0]*x2) / w[1]
	line, err := plotter.NewLine(plotter.XYs{{X: x1, Y: y1}, {X: x2, Y: y2}})
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	line.Color = color.RGBA{B: 255, A: 255}
	line.Width = vg.Points(1.5)
	p.Add(line)

	// 绘制分离超平面
	// p.X.Max = Max(dataMat)
	// p.X.Min = Min(dataMat)
	// lineFunc := func(x float64) float64 {
	// 	return (-b - w[0]*x) / w[1]
	// }
	// line := plotter.NewFunction(lineFunc)
	// line.Color = color.RGBA{G: 255, A: 255}
	// p.Add(line)

	// 保存图表
	if err := p.Save(8*vg.Inch, 8*vg.Inch, "classifier.png"); err != nil {
		log.Fatalf("error: %v\n", err)
	}
}

// 辅助函数，用于计算两个向量的点积
func dot(X, A []float64) float64 {
	sum := 0.0
	for i := range X {
		sum += X[i] * A[i]
	}
	return sum
}

// 线性核函数
func kernelTrans(X, A []float64) float64 {
	return dot(X, A)
}

// getW 计算权重向量w
func getW(dataMat [][]float64, labelMat []float64, alphas []float64) []float64 {
	m, n := len(dataMat), len(dataMat[0])
	w := make([]float64, n)

	// 计算权重向量w
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			w[j] += alphas[i] * labelMat[i] * dataMat[i][j]
		}
	}

	return w
}

// 辅助函数，用于向量元素相乘
func multiply(vec1, vec2 []float64) []float64 {
	result := make([]float64, len(vec1))
	for i := range vec1 {
		result[i] = vec1[i] * vec2[i]
	}
	return result
}

func showVec(vec []float64, min float64) {
	for i := range vec {
		if vec[i] > min {
			fmt.Printf("%f ", vec[i])
		}
	}
	fmt.Println()
}

// 辅助函数，用于计算向量元素之和
func sum(vec []float64) float64 {
	result := 0.0
	for _, value := range vec {
		result += value
	}
	return result
}

func smoSimple(dataMatIn [][]float64, classLabels []float64, C, toler float64, maxIter int) (float64, []float64) {
	dataMatrix := dataMatIn

	labelMat := make([]float64, len(classLabels))
	copy(labelMat, classLabels)

	b := 0.0
	m, _ := len(dataMatrix), len(dataMatrix[0])

	alphas := make([]float64, m)

	rand.Seed(time.Now().UnixNano())
	iter_num := 0
	for iter_num < maxIter {
		alphaPairsChanged := 0
		for i := 0; i < m; i++ {
			fXi := float64(0.0)
			for j := 0; j < m; j++ {
				fXi += alphas[j] * labelMat[j] * kernelTrans(dataMatrix[j], dataMatrix[i])
			}
			fXi += b
			Ei := fXi - float64(labelMat[i])

			if (classLabels[i]*Ei < -toler && alphas[i] < C) || (classLabels[i]*Ei > toler && alphas[i] > 0) {
				j := selectJrand(i, m)
				fXj := float64(0)
				for k := 0; k < m; k++ {
					fXj += alphas[k] * labelMat[k] * kernelTrans(dataMatrix[k], dataMatrix[j])
				}
				fXj += b
				Ej := fXj - float64(labelMat[j])

				alphaIold := alphas[i]
				alphaJold := alphas[j]
				var L, H float64
				if labelMat[i] != labelMat[j] {
					L = math.Max(0, alphas[j]-alphas[i])
					H = math.Min(C, C+alphas[j]-alphas[i])
				} else {
					L = math.Max(0, alphas[j]+alphas[i]-C)
					H = math.Min(C, alphas[j]+alphas[i])
				}

				if L == H {
					fmt.Println("L==H")
					continue
				}

				eta := 2.0*kernelTrans(dataMatrix[i], dataMatrix[j]) - kernelTrans(dataMatrix[i], dataMatrix[i]) - kernelTrans(dataMatrix[j], dataMatrix[j])
				if eta >= 0 {
					fmt.Println("eta>=0")
					continue
				}

				alphas[j] -= labelMat[j] * (Ei - Ej) / eta
				alphas[j] = clipAlpha(alphas[j], H, L)

				if math.Abs(alphas[j]-alphaJold) < 0.00001 {
					fmt.Println("j not moving enough")
					continue
				}

				alphas[i] += labelMat[j] * labelMat[i] * (alphaJold - alphas[j])

				b1 := b - Ei - labelMat[i]*(alphas[i]-alphaIold)*kernelTrans(dataMatrix[i], dataMatrix[i]) - labelMat[j]*(alphas[j]-alphaJold)*kernelTrans(dataMatrix[i], dataMatrix[j])
				b2 := b - Ej - labelMat[i]*(alphas[i]-alphaIold)*kernelTrans(dataMatrix[i], dataMatrix[j]) - labelMat[j]*(alphas[j]-alphaJold)*kernelTrans(dataMatrix[j], dataMatrix[j])

				if 0 < alphas[i] && C > alphas[i] {
					b = b1
				} else if 0 < alphas[j] && C > alphas[j] {
					b = b2
				} else {
					b = (b1 + b2) / 2
				}

				alphaPairsChanged += 1
				fmt.Printf("iter: %d i: %d, pairs changed %d\n", iter_num, i, alphaPairsChanged)
			}
		}
		if alphaPairsChanged == 0 {
			iter_num += 1
		} else {
			iter_num = 0
		}
		fmt.Printf("iteration number: %d\n", iter_num)
	}
	return b, alphas
}
