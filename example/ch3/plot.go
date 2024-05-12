package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var (
	leafNode     = color.RGBA{R: 0, G: 0, B: 0, A: 255}
	decisionNode = color.RGBA{R: 255, G: 0, B: 0, A: 255}
)

// PlotNode 绘制决策树节点
func PlotNode(p *plot.Plot, nodeTxt string, centerPt, parentPt plotter.XY, nodeType color.Color) {
	// p.X.Label.Text = nodeTxt
	// text := plotter.Annotation(plotter.XY{X: parentPt[0].X, Y: parentPt[0].Y}, nodeTxt)
	// text.TextStyle.Font.Size = vg.Points(12)
	// text.TextStyle.Color = color.Black
	// p.Add(text)

	// plotutil.AddLines(p, plotter.XYs{centerPt, parentPt})

	line, _ := plotter.NewLine(plotter.XYs{centerPt, parentPt})
	// line.LineStyle = nodeType
	labels, _ := plotter.NewLabels(plotter.XYLabels{
		XYs: []plotter.XY{
			centerPt,
		},
		Labels: []string{nodeTxt},
	})
	// var rt vg.Rectangle
	for i, t := range labels.TextStyle {
		labels.Offset.X = 0 - t.Width(nodeTxt)/2
		labels.Offset.Y = 0 - t.Height(nodeTxt)/2
		labels.TextStyle[i].Color = nodeType

	}

	// p.Add(plotter.NewGlyphBoxes())

	// fmt.Println(labels)
	p.Add(labels)
	p.Add(line)

	// plotutil.AddBoxPlots(p, 0.1, nodeTxt, plotter.Values{0.1, 0.1})
	// l, _ := plotter.NewLabels(plotter.XYLabels{
	// 	Labels: []string{nodeTxt},
	// })

	// line, s, _ := plotter.NewLinePoints(plotter.Labels{
	// 	XYs:    plotter.XYs{parentPt, centerPt},
	// 	Labels: []string{nodeTxt},
	// })
	// line.LineStyle = nodeType

	// p.Add(line)
	// p.Add(s)
	// p.Add(l)
}

// CreatePlot 创建绘图
func CreatePlot() {
	p := plot.New()

	p.Title.Text = "Decision Tree"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.X.Min = 0
	p.X.Max = 1
	p.Y.Min = 0
	p.Y.Max = 1

	// 创建决策节点
	PlotNode(p, "a decision node", plotter.XY{X: 0.5, Y: 0.1}, plotter.XY{X: 0.1, Y: 0.5}, decisionNode)

	// 创建叶子节点
	PlotNode(p, "a leaf node", plotter.XY{X: 0.8, Y: 0.1}, plotter.XY{X: 0.3, Y: 0.8}, leafNode)

	if err := p.Save(8*vg.Inch, 8*vg.Inch, "decision_tree.png"); err != nil {
		panic(err)
	}
}
