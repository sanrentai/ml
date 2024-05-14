package util

func Shape(arr [][]float64) (int, int) {
	m := len(arr)
	if m == 0 {
		panic("empty array")
	}
	n := len(arr[0])
	return m, n
}
