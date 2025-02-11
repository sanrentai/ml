package main

import (
	"fmt"

	"github.com/sanrentai/ml/tree"
	"gonum.org/v1/plot"

	// "gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var (
	totalW, totalD, xOff, yOff float64
)

// CreatePlot 绘制树形结构
func CreatePlotTree(inTree *tree.Tree) {
	p := plot.New()
	p.HideAxes()
	totalW = float64(inTree.GetNumLeafs())
	totalD = float64(inTree.GetTreeDepth())
	p.X.Min = 0
	p.X.Max = 1
	p.Y.Min = 0
	p.Y.Max = 1

	xOff = -0.5 / totalW
	yOff = 1.0
	PlotTree(p, inTree, plotter.XY{X: 0.5, Y: 1.0}, "")

	if err := p.Save(5*vg.Inch, 5*vg.Inch, "tree_plot.png"); err != nil {
		panic(err)
	}
}

func PlotMidText(p *plot.Plot, cntrPt, parentPt plotter.XY, txtString string) {
	xMid := (parentPt.X-cntrPt.X)/2 + cntrPt.X
	yMid := (parentPt.Y-cntrPt.Y)/2 + cntrPt.Y

	labels, _ := plotter.NewLabels(plotter.XYLabels{
		XYs: []plotter.XY{
			{
				X: xMid,
				Y: yMid,
			},
		},
		Labels: []string{txtString},
	})
	for _, t := range labels.TextStyle {
		labels.Offset.X = 0 - t.Width(txtString)/2
		labels.Offset.Y = 0 - t.Height(txtString)
	}
	p.Add(labels)
}

func PlotTree(p *plot.Plot, myTree *tree.Tree, parentPt plotter.XY, nodeTxt string) {
	numLeafs := myTree.GetNumLeafs()
	// depth := ml.GetTreeDepth(myTree)

	cntrPt := plotter.XY{
		X: xOff + (1+float64(numLeafs))/2.0/totalW,
		Y: yOff,
	}
	PlotMidText(p, cntrPt, parentPt, nodeTxt)
	PlotNode(p, myTree.Value.(string), cntrPt, parentPt, decisionNode)
	yOff -= 1.0 / totalD
	for _, key := range myTree.Keys {
		child := myTree.Children[key]
		if len(child.Children) > 0 {
			PlotTree(p, child, cntrPt, fmt.Sprintf("%v", key))
		} else {
			xOff += 1.0 / totalW
			PlotNode(p, child.Value.(string), plotter.XY{
				X: xOff,
				Y: yOff,
			}, cntrPt, leafNode)
			PlotMidText(p, plotter.XY{
				X: xOff,
				Y: yOff,
			}, cntrPt, fmt.Sprintf("%v", key))
		}
	}
	yOff += 1.0 / totalD
}

func plot_tree() {
	// 构建示例树
	// 绘制树形结构
	myTree := retrieveTree(0)
	myTree.AddChild(3, tree.NewTree("maybe"))
	// myTree.Children = append(myTree.Children, tree.NewTree("maybe"))
	CreatePlotTree(myTree)
}
