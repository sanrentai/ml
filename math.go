package ml

func Sign(v float64) float64 {
	if v < 0 {
		return -1
	} else if v > 0 {
		return 1
	}
	return 0
}

func SignVec(v []float64) []float64 {
	r := make([]float64, len(v))
	for i := range r {
		r[i] = Sign(v[i])
	}
	return r
}
