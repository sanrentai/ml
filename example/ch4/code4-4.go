package main

func bagOfWords2VecMN(vocabList []string, inputSet []string) []int {
	returnVec := make([]int, len(vocabList))
	for _, word := range inputSet {
		for i, v := range vocabList {
			if v == word {
				returnVec[i]++
				break
			}
		}
	}
	return returnVec
}
