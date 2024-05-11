package ml

import "math"

// 计算两个向量的欧式距离
func EuclideanDistance(vec1, vec2 []float64) float64 {
	sum := 0.0
	for i := range vec1 {
		sum += math.Pow(vec1[i]-vec2[i], 2)
	}
	return math.Sqrt(sum)
}
