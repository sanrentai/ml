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
		tree.NewTree("no surfacing").
			AddChild(0, tree.NewTree("no")).
			AddChild(1, tree.NewTree("flippers").
				AddChild(0, tree.NewTree("no")).
				AddChild(1, tree.NewTree("yes")),
			),
		tree.NewTree("no surfacing").
			AddChild(0, tree.NewTree("no")).
			AddChild(1, tree.NewTree("flippers").
				AddChild(0, tree.NewTree("head").
					AddChild(0, tree.NewTree("no")).
					AddChild(1, tree.NewTree("yes"))).
				AddChild(1, tree.NewTree("no")),
			),
	}
	return listOfTrees[i]
}
