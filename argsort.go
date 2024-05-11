package ml

import "sort"

// argsort 返回数组排序后的索引数组
func Argsort(arr []float64) []int {
	// 创建索引数组
	indexes := make([]int, len(arr))
	for i := range indexes {
		indexes[i] = i
	}

	// 排序索引数组，使用 arr[indexes[i]] 作为比较函数
	sort.Slice(indexes, func(i, j int) bool {
		return arr[indexes[i]] < arr[indexes[j]]
	})

	return indexes
}
