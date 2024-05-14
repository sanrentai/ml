package main

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/mmcdole/gofeed"
	"github.com/sanrentai/ml/util"
)

func main6() {
	fp := gofeed.NewParser()
	ny, err := fp.ParseURL("xxx")
	if err != nil {
		panic(err)
	}
	sf, err := fp.ParseURL("xxx")
	if err != nil {
		panic(err)
	}
	// fmt.Println(ny)
	localWords(ny, sf)
}

func calcMostFreq(vocabList []string, fullText []string) []string {
	freqDict := make(map[string]int)
	for _, token := range vocabList {
		freqDict[token] = util.Count(fullText, token)
	}

	var sortedFreq []string
	for token, count := range freqDict {
		sortedFreq = append(sortedFreq, fmt.Sprintf("%s:%d", token, count))
	}

	sort.Slice(sortedFreq, func(i, j int) bool {
		return freqDict[sortedFreq[i]] > freqDict[sortedFreq[j]]
	})

	if len(sortedFreq) > 30 {
		return sortedFreq[:30]
	}
	return sortedFreq
}

// localWords 函数用于从 RSS 源获取数据并进行分类
func localWords(feed1, feed0 *gofeed.Feed) ([]string, []float64, []float64) {
	docList := [][]string{}
	classList := []int{}
	fullText := []string{}
	minLen := util.Min(len(feed1.Items), len(feed0.Items))
	for i := 0; i < minLen; i++ {
		wordList := textParse(feed1.Items[i].Description)
		docList = append(docList, wordList)
		fullText = append(fullText, wordList...)
		classList = append(classList, 1) // NY is class 1
		wordList = textParse(feed0.Items[i].Description)
		docList = append(docList, wordList)
		fullText = append(fullText, wordList...)
		classList = append(classList, 0)
	}
	vocabList := createVocabList(docList)
	top30Words := calcMostFreq(vocabList, fullText)
	for _, pairW := range top30Words {
		if util.Contains(vocabList, pairW) {
			vocabList = util.Remove(vocabList, pairW)
		}
	}
	trainingSet := make([]int, 2*minLen)
	for i := 0; i < len(trainingSet); i++ {
		trainingSet[i] = i
	}
	testSet := []int{}
	for i := 0; i < 20; i++ {
		randIndex := rand.Intn(len(trainingSet))
		testSet = append(testSet, trainingSet[randIndex])
		trainingSet = append(trainingSet[:randIndex], trainingSet[randIndex+1:]...)
	}
	trainMat := [][]int{}
	trainClasses := []int{}
	for _, docIndex := range trainingSet {
		trainMat = append(trainMat, bagOfWords2VecMN(vocabList, docList[docIndex]))
		trainClasses = append(trainClasses, classList[docIndex])
	}
	p0V, p1V, pSpam := trainNB0(trainMat, trainClasses)
	errorCount := 0
	for _, docIndex := range testSet {
		wordVector := bagOfWords2VecMN(vocabList, docList[docIndex])
		if classifyNB(wordVector, p0V, p1V, pSpam) != classList[docIndex] {
			errorCount++
		}
	}
	fmt.Printf("the error rate is: %f\n", float64(errorCount)/float64(len(testSet)))
	return vocabList, p0V, p1V
}

// getTopWords 函数用于输出最频繁出现的词汇
func getTopWords(ny, sf *gofeed.Feed) {
	vocabList, p0V, p1V := localWords(ny, sf)
	topNY := []string{}
	topSF := []string{}
	for i := 0; i < len(p0V); i++ {
		if p0V[i] > -6.0 {
			topSF = append(topSF, vocabList[i])
		}
		if p1V[i] > -6.0 {
			topNY = append(topNY, vocabList[i])
		}
	}
	fmt.Println("SF**SF**SF**SF**SF**SF**SF**SF**SF**SF**SF**SF**SF**SF**SF**SF**")
	for _, word := range topSF {
		fmt.Println(word)
	}
	fmt.Println("NY**NY**NY**NY**NY**NY**NY**NY**NY**NY**NY**NY**NY**NY**NY**NY**")
	for _, word := range topNY {
		fmt.Println(word)
	}
}
