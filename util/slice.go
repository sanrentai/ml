package util

func Count[T comparable](slice []T, v T) int {
	count := 0
	for _, x := range slice {
		if x == v {
			count++
		}
	}
	return count
}

// contains 函数用于检查切片中是否包含指定元素
func Contains[T comparable](slice []T, item T) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// Remove 函数用于从切片中移除指定元素
func Remove[T comparable](slice []T, item T) []T {
	index := -1
	for i, s := range slice {
		if s == item {
			index = i
			break
		}
	}
	if index == -1 {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
