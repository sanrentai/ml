package tree

import (
	"fmt"
	"strings"

	"github.com/sanrentai/ml"
)

// 决策树
type Tree struct {
	Value    interface{}
	Keys     []interface{}
	Children map[interface{}]*Tree
}

func NewTree(val interface{}) *Tree {
	return &Tree{
		Value: val,
	}
}

func (myTree *Tree) AddChild(val interface{}, child *Tree) *Tree {
	if myTree.Children == nil {
		myTree.Children = make(map[interface{}]*Tree)
		myTree.Keys = make([]interface{}, 0)
	}
	myTree.Children[val] = child
	myTree.Keys = append(myTree.Keys, val)
	return myTree
}

// 分类
func (myTree *Tree) Classify(featureLabels []string, testVec []interface{}) interface{} {
	return classify(myTree, featureLabels, testVec)
}

func (t *Tree) String() string {
	var sb strings.Builder

	if len(t.Children) == 0 {
		sb.WriteString(fmt.Sprintf(`"%v"`, t.Value))
	} else {
		sb.WriteString("{")
		sb.WriteString(fmt.Sprintf(`"%v": `, t.Value))
		sb.WriteString("{")
		for i, key := range t.Keys {
			if k, ok := key.(string); ok {
				sb.WriteString(fmt.Sprintf(`"%v": %v`, k, t.Children[key]))
			} else {
				sb.WriteString(fmt.Sprintf(`%v: %v`, key, t.Children[key]))
			}

			if i < len(t.Keys)-1 {
				sb.WriteString(", ")
			}
		}
		sb.WriteString("}")
		sb.WriteString("}")
	}

	return sb.String()
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
		return &Tree{Value: ml.MajorityCnt(classList)}
	}

	bestFeat := ml.ChooseBestFeatureToSplit(dataSet)
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

	uniqueVals := ml.Set(featValues)

	for _, value := range uniqueVals {
		subLabels := make([]string, len(labels))
		copy(subLabels, labels)
		subDataSet := ml.SplitDataSet(dataSet, bestFeat, value)
		myTree.AddChild(value, CreateTree(subDataSet, subLabels))
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

// 获取叶节点的数目
func (myTree *Tree) GetNumLeafs() int {
	if myTree == nil {
		return 0
	}
	if len(myTree.Children) == 0 {
		return 1
	}
	numLeafs := 0
	for _, child := range myTree.Children {
		numLeafs += child.GetNumLeafs()
	}
	return numLeafs
}

// 获取树的深度
func (myTree *Tree) GetTreeDepth() int {
	if myTree == nil {
		return 0
	}
	if len(myTree.Children) == 0 {
		return 0
	}
	maxDepth := 0
	for _, child := range myTree.Children {
		thisDepth := 1 + child.GetTreeDepth()
		if thisDepth > maxDepth {
			maxDepth = thisDepth
		}
	}
	return maxDepth
}
