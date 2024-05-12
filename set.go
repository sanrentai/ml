package ml

// 用于数组去重
func set(arr []any) []any {
	uniqueMap := make(map[any]bool)
	uniqueSlice := []any{}

	for _, v := range arr {
		if !uniqueMap[v] {
			uniqueMap[v] = true
			uniqueSlice = append(uniqueSlice, v)
		}
	}

	return uniqueSlice
}
