package main

import "fmt"

func main1() {
	listOPosts, listClasses := loadDataSet()
	fmt.Println(listClasses)
	myVocabList := createVocabList(listOPosts)
	fmt.Println(myVocabList)
	fmt.Println(setOfWords2Vec(myVocabList, listOPosts[0]))
	fmt.Println(setOfWords2Vec(myVocabList, listOPosts[3]))
}

func loadDataSet() ([][]string, []int) {
	postingList := [][]string{
		{"my", "dog", "has", "flea", "problems", "help", "please"},
		{"maybe", "not", "take", "him", "to", "dog", "park", "stupid"},
		{"my", "dalmation", "is", "so", "cute", "I", "love", "him"},
		{"stop", "posting", "stupid", "worthless", "garbage"},
		{"mr", "licks", "ate", "my", "steak", "how", "to", "stop", "him"},
		{"quit", "buying", "worthless", "dog", "food", "stupid"},
	}

	classVec := []int{0, 1, 0, 1, 0, 1}
	return postingList, classVec
}

func createVocabList(dataSet [][]string) []string {
	vocabSet := make(map[string]struct{})
	for _, document := range dataSet {
		for _, word := range document {
			vocabSet[word] = struct{}{}
		}
	}

	vocabList := make([]string, 0, len(vocabSet))
	for word := range vocabSet {
		vocabList = append(vocabList, word)
	}
	return vocabList
}

func setOfWords2Vec(vocabList []string, inputSet []string) []int {
	vocabMap := make(map[string]int)
	for i, word := range vocabList {
		vocabMap[word] = i
	}

	returnVec := make([]int, len(vocabList))
	for _, word := range inputSet {
		index, found := vocabMap[word]
		if found {
			returnVec[index] = 1
		} else {
			fmt.Printf("the word: %s is not in my Vocabulary!\n", word)
		}
	}
	return returnVec
}
