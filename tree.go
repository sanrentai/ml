package ml

import (
	"fmt"
	"strings"
)

// TreeNode 表示决策树节点
type TreeNode struct {
	SplitFeature string
	Children     map[interface{}]*TreeNode
	LeafValue    interface{}
}

func CreateTree(dataSet [][]interface{}, labels []string) *Tree {
	classList := make([]interface{}, len(dataSet))
	for i, example := range dataSet {
		classList[i] = example[len(example)-1]
	}
	// 如果类别完全相同，则停止划分，返回该类别
	if countOccurrences(classList, classList[0]) == len(classList) {
		return &Tree{Value: classList[0]}
	}
	if len(dataSet[0]) == 1 {
		return &Tree{Value: MajorityCnt(classList)}
	}

	bestFeat := ChooseBestFeatureToSplit(dataSet)
	bestFeatLabel := labels[bestFeat]
	myTree := &Tree{
		Value: bestFeatLabel,
	}

	// 清空labels[bestFeat]
	labels = append(labels[:bestFeat], labels[bestFeat+1:]...)

	featValues := make([]interface{}, len(dataSet))
	for i, example := range dataSet {
		featValues[i] = example[bestFeat]
	}

	uniqueVals := set(featValues)

	for value := range uniqueVals {
		subLabels := make([]string, len(labels))
		copy(subLabels, labels)
		subDataSet := SplitDataSet(dataSet, bestFeat, value)
		if myTree.Children == nil {
			myTree.Children = make([]*Tree, 0)
		}
		myTree.Children = append(myTree.Children, CreateTree(subDataSet, subLabels))
	}

	return myTree
}

// countOccurrences 计算列表中指定元素出现的次数
func countOccurrences(list []interface{}, elem interface{}) int {
	count := 0
	for _, e := range list {
		if e == elem {
			count++
		}
	}
	return count
}

type Tree struct {
	Value    interface{}
	Children []*Tree
}

func NewTree(val interface{}, tree ...*Tree) *Tree {
	return &Tree{
		Value:    val,
		Children: tree,
	}
}

func (t *Tree) String() string {
	var sb strings.Builder

	if len(t.Children) == 0 {
		sb.WriteString(fmt.Sprintf(`"%v"`, t.Value))
	} else {
		sb.WriteString("{")
		sb.WriteString(fmt.Sprintf(`"%v": `, t.Value))
		sb.WriteString("{")
		for i, child := range t.Children {
			sb.WriteString(fmt.Sprintf(`%v: %v`, i, child))
			if i < len(t.Children)-1 {
				sb.WriteString(", ")
			}
		}
		sb.WriteString("}")
		sb.WriteString("}")
	}

	return sb.String()
}

// 获取叶节点的数目
func GetNumLeafs(myTree *Tree) int {
	if myTree == nil {
		return 0
	}
	if len(myTree.Children) == 0 {
		return 1
	}
	numLeafs := 0
	for _, child := range myTree.Children {
		numLeafs += GetNumLeafs(child)
	}
	return numLeafs
}

// 获取树的深度
func GetTreeDepth(myTree *Tree) int {
	if myTree == nil {
		return 0
	}
	if len(myTree.Children) == 0 {
		return 0
	}
	maxDepth := 0
	for _, child := range myTree.Children {
		thisDepth := 1 + GetTreeDepth(child)
		if thisDepth > maxDepth {
			maxDepth = thisDepth
		}
	}
	return maxDepth
}
