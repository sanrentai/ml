package main

import (
	"fmt"

	"github.com/sanrentai/ml/matrix"
)

func main() {
	randMat := matrix.Rand(4, 4)
	fmt.Println(randMat)
	invRandMat := randMat.I()
	fmt.Println(invRandMat)
	myEye := matrix.Multiplication(randMat, invRandMat, 1.0)
	fmt.Println(myEye)
	fmt.Println(matrix.Sub(myEye, matrix.Eye(4)))
}
