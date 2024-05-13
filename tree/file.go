package tree

import (
	"encoding/gob"
	"os"
)

// 将分类器存储在硬盘上，不用每次对数据分类时重新学习一遍
// 存储决策树到文件
func (inputTree *Tree) Store(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(inputTree)
	if err != nil {
		return err
	}

	return nil
}

// 从文件加载决策树
func GrabTree(filename string) (*Tree, error) {
	var inputTree Tree
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&inputTree)
	if err != nil {
		return nil, err
	}
	return &inputTree, nil
}
