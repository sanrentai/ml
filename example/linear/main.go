package main

import (
	"fmt"
	"math/rand"
	"time"

	// "gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成示例数据
	X := make([]float64, 100)
	y := make([]float64, 100)
	for i := range X {
		X[i] = 2 * rand.Float64()              // 生成0到2之间的随机数
		y[i] = 4 + 3*X[i] + rand.NormFloat64() // 线性关系模型 y = 4 + 3*x + 噪声
	}

	// 将数据转换为矩阵形式
	XMat := mat.NewDense(len(X), 1, X)
	yVec := mat.NewVecDense(len(y), y)

	// 构建设计矩阵（添加截距项）
	var ones []float64
	for range X {
		ones = append(ones, 1)
	}
	onesMat := mat.NewDense(len(ones), 1, ones)
	designMat := mat.NewDense(len(X), 2, nil)
	designMat.Augment(onesMat, XMat)

	// 拟合线性回归模型
	var reg mat.Dense
	reg.Solve(designMat, yVec)

	// 提取模型参数
	coef := make([]float64, 2)
	coef[0] = reg.At(0, 0) // 截距
	coef[1] = reg.At(1, 0) // 系数

	fmt.Println("Intercept:", coef[0])
	fmt.Println("Coefficient:", coef[1])

	// 在测试集上进行预测
	XTest := make([]float64, 20)
	for i := range XTest {
		XTest[i] = 2 * rand.Float64() // 生成0到2之间的随机数
	}
	XTestMat := mat.NewDense(len(XTest), 1, XTest)
	var yPred mat.Dense
	yPred.Mul(XTestMat, mat.NewDense(1, 2, coef))

	// 可视化结果
	XTestSlice := XTestMat.ColView(0)
	yPredSlice := yPred.ColView(0)
	plotResults(XTestSlice, yPredSlice, X, y)
}

func plotResults(xPred, yPred mat.Vector, X, y []float64) {
	p := plot.New()

	points := make(plotter.XYs, len(X))
	for i := range points {
		points[i].X = X[i]
		points[i].Y = y[i]
	}
	scatter, err := plotter.NewScatter(points)
	if err != nil {
		panic(err)
	}
	p.Add(scatter)

	line, _ := plotter.NewLine(plotter.XYs{
		{xPred.AtVec(0), yPred.AtVec(0)},
		{xPred.AtVec(xPred.Len() - 1), yPred.AtVec(yPred.Len() - 1)},
	})
	p.Add(line)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, "linear_regression.png"); err != nil {
		panic(err)
	}
}
