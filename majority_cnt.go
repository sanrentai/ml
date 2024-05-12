package ml

// MajorityCnt 返回列表中出现次数最多的元素
func MajorityCnt[T comparable](classList []T) T {
	classCount := make(map[T]int)
	for _, vote := range classList {
		classCount[vote]++
	}
	var major T
	maxCount := 0
	for vote, count := range classCount {
		if count > maxCount {
			maxCount = count
			major = vote
		}
	}
	return major
}
