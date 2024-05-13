package main

import (
	"fmt"

	"github.com/sanrentai/ml/tree"
)

func fn39() {
	myTree := retrieveTree(0)
	fmt.Println(myTree)
	filename := "classifierStorage.txt"
	err := myTree.Store(filename)
	if err != nil {
		panic(err)
	}
	myTree2, err := tree.GrabTree(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(myTree2)
}
