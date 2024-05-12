package main

import (
	"fmt"

	"github.com/sanrentai/ml"
)

func fn36() {
	fmt.Println(retrieveTree(1))
	myTree := retrieveTree(0)
	fmt.Println(myTree)                  //3
	fmt.Println(ml.GetNumLeafs(myTree))  //3
	fmt.Println(ml.GetTreeDepth(myTree)) //2
}

func retrieveTree(i int) *ml.Tree {
	listOfTrees := []*ml.Tree{
		ml.NewTree("no surfacing",
			ml.NewTree("no"),
			ml.NewTree("flippers",
				ml.NewTree("no"),
				ml.NewTree("yes"),
			),
		),
		ml.NewTree("no surfacing",
			ml.NewTree("no"),
			ml.NewTree("flippers",
				ml.NewTree("head",
					ml.NewTree("no"),
					ml.NewTree("yes"),
				),
				ml.NewTree("yes"),
			),
		),
	}
	return listOfTrees[i]
}
