package main

import (
	"fmt"

	"github.com/sanrentai/ml/tree"
	"github.com/sanrentai/ml/util"
)

func main() {
	// fn31()
	// fn32()
	// fn33()
	// fn34()
	// fn35()
	// fn36()
	plot_tree()

	// plottest()

	// fn38()
	// fn39()
	dataset, err := util.File2Dataset("lenses.txt")
	if err != nil {
		panic(err)
	}
	// fmt.Println(dataset)
	lensesLabels := []string{"age", "prescript", "astigmatic", "tearRate"}
	lensesTree := tree.CreateTree(dataset, lensesLabels)
	fmt.Println(lensesTree)
	CreatePlotTree(lensesTree)
}

// 算法称为ID3,不完美，无法之间处理数值型数据，第九章学习另一个决策树构造算法 CART
// 还有其他的算法 C4.5
