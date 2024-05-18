package ml

func Sign(v float64) float64 {
	if v < 0 {
		return -1
	} else if v > 0 {
		return 1
	}
	return 0
}
