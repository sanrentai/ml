package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main2() {
	dataArr, labelMat := loadDataSet()
	wei := gradAscent(dataArr, labelMat)
	plotBestFit(wei, "gradAscent.png")
}

func plotBestFit(weights []float64, filename string) error {
	dataMat, labelMat := loadDataSet()
	// weightMat := mat.NewDense(len(weights), 1, weights)
	// 提取数据集中的数据
	var xys1, xys2 plotter.XYs

	for i, data := range dataMat {
		if labelMat[i] == 1 {
			xys1 = append(xys1, plotter.XY{
				X: data[1],
				Y: data[2],
			})
		} else {
			xys2 = append(xys2, plotter.XY{
				X: data[1],
				Y: data[2],
			})
		}
	}
	// 创建绘图
	p := plot.New()
	s1, err := plotter.NewScatter(xys1)
	if err != nil {
		return err
	}
	s1.GlyphStyle.Color = plotutil.Color(0)

	s2, err := plotter.NewScatter(xys2)
	if err != nil {
		return err
	}
	s1.GlyphStyle.Color = plotutil.Color(1)
	p.Add(s1, s2)

	plotutil.AddScatters(p, s1, s2)

	// 绘制拟合直线
	lineFunc := func(x float64) float64 {
		return (-weights[0] - weights[1]*x) / weights[2]
	}
	ff := plotter.NewFunction(lineFunc)
	ff.XMin = -3
	ff.XMax = 3
	p.Add(ff)
	p.X.Label.Text = "X1"
	p.Y.Label.Text = "X2"
	p.X.Max = 4
	p.X.Min = -4

	// 保存图像
	if err := p.Save(4*vg.Inch, 4*vg.Inch, filename); err != nil {
		return err
	}

	return nil
}
