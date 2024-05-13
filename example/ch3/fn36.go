package main

import (
	"fmt"

	"github.com/sanrentai/ml/tree"
)

func fn36() {
	fmt.Println(retrieveTree(1))
	myTree := retrieveTree(0)
	fmt.Println(myTree)                //3
	fmt.Println(myTree.GetNumLeafs())  //3
	fmt.Println(myTree.GetTreeDepth()) //2
}

func retrieveTree(i int) *tree.Tree {
	listOfTrees := []*tree.Tree{
		tree.NewTree("no surfacing",
			tree.NewTree("no"),
			tree.NewTree("flippers",
				tree.NewTree("no"),
				tree.NewTree("yes"),
			),
		),
		tree.NewTree("no surfacing",
			tree.NewTree("no"),
			tree.NewTree("flippers",
				tree.NewTree("head",
					tree.NewTree("no"),
					tree.NewTree("yes"),
				),
				tree.NewTree("yes"),
			),
		),
	}
	return listOfTrees[i]
}
