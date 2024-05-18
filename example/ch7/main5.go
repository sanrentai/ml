package main

import (
	"fmt"
	"image/color"

	"github.com/sanrentai/ml"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main5() {
	datArr, labelArr := loadDataSet("horseColicTraining2.txt")
	_, aggClassEst := adaBosstTrainDS(datArr, labelArr, 10000)
	plotROC(aggClassEst, labelArr)

}

func plotROC(predStrengths []float64, labels []float64) {
	ySum := 0.0
	plt := plot.New()
	numPosClas := 0.0
	cur := []float64{1.0, 1.0}
	for i := 0; i < len(labels); i++ {
		if labels[i] == 1.0 {
			numPosClas += 1
		}
	}
	yStep := 1.0 / float64(numPosClas)
	xStep := 1.0 / (float64(len(labels)) - numPosClas)
	sortedIndicies := ml.Argsort(predStrengths)

	for _, index := range sortedIndicies {
		var (
			delX, delY float64
		)
		if labels[index] == 1.0 {
			delX = 0
			delY = yStep
		} else {
			delX = xStep
			delY = 0
			ySum += cur[1]
		}
		line, _ := plotter.NewLine(plotter.XYs{
			{X: cur[0], Y: cur[1]},
			{X: cur[0] - delX, Y: cur[1] - delY},
		})
		line.Color = color.RGBA{B: 255, A: 255}
		plt.Add(line)
		cur[0] -= delX
		cur[1] -= delY
	}
	line, _ := plotter.NewLine(plotter.XYs{
		{X: 0, Y: 0},
		{X: 1, Y: 1},
	})
	line.Dashes = []vg.Length{vg.Points(10), vg.Points(10)} //虚线
	line.Color = color.RGBA{B: 255, A: 255}
	plt.Add(line)
	plt.X.Label.Text = "False Positive Rate"
	plt.Y.Label.Text = "True Positive Rate"
	plt.Title.Text = "ROC Curve"
	plt.X.Min = 0
	plt.X.Max = 1
	plt.Y.Min = 0
	plt.Y.Max = 1
	plt.Save(8*vg.Inch, 8*vg.Inch, "roc.png")
	fmt.Printf("the Area Under the Curve is: %v", ySum*xStep)
}
