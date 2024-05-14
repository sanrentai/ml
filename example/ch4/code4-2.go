package main

import (
	"fmt"
	"math"
)

func main2() {
	listOPosts, listClasses := loadDataSet()
	myVocabList := createVocabList(listOPosts)
	trainMat := [][]int{}
	for _, postinDoc := range listOPosts {
		trainMat = append(trainMat, setOfWords2Vec(myVocabList, postinDoc))
	}
	p0V, p1V, pAb := trainNB0(trainMat, listClasses)
	fmt.Println(pAb)
	fmt.Println(p0V)
	fmt.Println(p1V)
}

func trainNB0(trainMatrix [][]int, trainCategory []int) ([]float64, []float64, float64) {
	numTrainDocs := len(trainMatrix)
	numWords := len(trainMatrix[0])
	pAbusive := float64(sum(trainCategory)) / float64(numTrainDocs)
	p0Num := make([]int, numWords)
	p1Num := make([]int, numWords)
	for i := 0; i < numWords; i++ {
		p0Num[i] = 1
		p1Num[i] = 1
	}
	p0Denom := 2.0
	p1Denom := 2.0
	for i := 0; i < numTrainDocs; i++ {
		if trainCategory[i] == 1 {
			p1Num = addVectors(p1Num, trainMatrix[i])
			p1Denom += float64(sum(trainMatrix[i]))
		} else {
			p0Num = addVectors(p0Num, trainMatrix[i])
			p0Denom += float64(sum(trainMatrix[i]))
		}
	}
	p1Vect := make([]float64, numWords)
	p0Vect := make([]float64, numWords)
	for i := 0; i < numWords; i++ {
		p1Vect[i] = math.Log(float64(p1Num[i]) / p1Denom) // float64(p1Num[i]) / p1Denom //
		p0Vect[i] = math.Log(float64(p0Num[i]) / p0Denom) // float64(p0Num[i]) / p0Denom //
	}
	return p0Vect, p1Vect, pAbusive
}

func testingNB() {
	// 加载数据集
	listOPosts, listClasses := loadDataSet()
	// 创建词汇表
	myVocabList := createVocabList(listOPosts)
	trainMat := [][]int{}
	// 将文档转换为词袋向量
	for _, postinDoc := range listOPosts {
		trainMat = append(trainMat, setOfWords2Vec(myVocabList, postinDoc))
	}
	// 训练朴素贝叶斯模型
	p0V, p1V, pAb := trainNB0(trainMat, listClasses)
	// 测试样本
	testEntry1 := []string{"love", "my", "dalmation"}
	thisDoc1 := setOfWords2Vec(myVocabList, testEntry1)
	fmt.Print(testEntry1, " classified as: ", classifyNB(thisDoc1, p0V, p1V, pAb), "\n")

	testEntry2 := []string{"stupid", "garbage"}
	thisDoc2 := setOfWords2Vec(myVocabList, testEntry2)
	fmt.Print(testEntry2, " classified as: ", classifyNB(thisDoc2, p0V, p1V, pAb), "\n")
}

func sum[T int | float64](arr []T) T {
	var sum T
	for _, val := range arr {
		sum += val
	}
	return sum
}

func addVectors(vec1, vec2 []int) []int {
	result := make([]int, len(vec1))
	for i := range vec1 {
		result[i] = vec1[i] + vec2[i]
	}
	return result
}

func classifyNB(vec2Classify []int, p0Vec, p1Vec []float64, pClass1 float64) int {
	p1 := sum(multiplyVectors(vec2Classify, p1Vec)) + math.Log(pClass1)
	p0 := sum(multiplyVectors(vec2Classify, p0Vec)) + math.Log(1.0-pClass1)
	if p1 > p0 {
		return 1
	} else {
		return 0
	}
}

func multiplyVectors[T, K int | float64](vec1 []K, vec2 []T) []float64 {
	result := make([]float64, len(vec1))
	for i := range vec1 {
		result[i] = float64(vec1[i]) * float64(vec2[i])
	}
	return result
}
