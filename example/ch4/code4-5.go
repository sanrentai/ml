package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

func textParse(bigString string) []string {
	regex := regexp.MustCompile(`\s+`)
	listOfTokens := regex.Split(bigString, -1)
	var result []string
	for _, tok := range listOfTokens {
		if len(tok) > 2 {
			result = append(result, strings.ToLower(tok))
		}
	}
	return result
}

func spamTest() {
	var docList [][]string
	var classList []int
	var fullText []string

	for i := 1; i <= 25; i++ {
		spamBytes, _ := os.ReadFile(fmt.Sprintf("email/spam/%d.txt", i))
		hamBytes, _ := os.ReadFile(fmt.Sprintf("email/ham/%d.txt", i))
		spamWordList := textParse(string(spamBytes))
		hamWordList := textParse(string(hamBytes))

		docList = append(docList, spamWordList)
		docList = append(docList, hamWordList)

		fullText = append(fullText, spamWordList...)
		fullText = append(fullText, hamWordList...)

		classList = append(classList, 1)
		classList = append(classList, 0)
	}
	// fmt.Println(docList)
	vocabList := createVocabList(docList)
	var trainingSet []int
	testSet := make([]int, 0)
	// 创建存储训练集的索引值的列表
	for i := 0; i < 50; i++ {
		trainingSet = append(trainingSet, i)
	}

	// 从50个邮件中，随机挑选出40个作为训练集，10个作为测试集
	for i := 0; i < 10; i++ {
		// 随机选取索引值
		randIndex := rand.Intn(len(trainingSet))
		// 添加测试集的索引值
		testSet = append(testSet, trainingSet[randIndex])
		// 在训练集的列表中删除添加到测试集的索引值
		trainingSet = append(trainingSet[:randIndex], trainingSet[randIndex+1:]...)
	}

	var trainMat [][]int
	var trainClasses []int

	for _, docIndex := range trainingSet {
		trainMat = append(trainMat, setOfWords2Vec(vocabList, docList[docIndex]))
		trainClasses = append(trainClasses, classList[docIndex])
	}

	p0V, p1V, pSpam := trainNB0(trainMat, trainClasses)
	errorCount := 0

	for _, docIndex := range testSet {
		wordVector := setOfWords2Vec(vocabList, docList[docIndex])
		if classifyNB(wordVector, p0V, p1V, pSpam) != classList[docIndex] {
			errorCount++
			fmt.Println("classification error", docList[docIndex])
		}
	}
	fmt.Println("the error rate is:", float64(errorCount)/float64(len(testSet)))
}
