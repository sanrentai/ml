package ml

import "math"

func VecMin(vec []float64) float64 {
	min := vec[0]
	for _, v := range vec {
		if v < min {
			min = v
		}
	}
	return min
}

func VecMax(vec []float64) float64 {
	max := vec[0]
	for _, v := range vec {
		if v > max {
			max = v
		}
	}
	return max
}

func VecDot(a, b []float64) float64 {
	return Dot(a, b)
}

func VecProd(a float64, b []float64) []float64 {
	r := make([]float64, len(b))
	for i := 0; i < len(b); i++ {
		r[i] = a * b[i]
	}
	return r
}

func VecSum(vec []float64) float64 {
	sum := 0.0
	for _, v := range vec {
		sum += v
	}
	return sum
}

func VecMean(vec []float64) float64 {
	return VecSum(vec) / float64(len(vec))
}

func VecSub(a, b []float64) []float64 {
	r := make([]float64, len(a))
	for i := 0; i < len(a); i++ {
		r[i] = a[i] - b[i]
	}
	return r
}

func VecAdd(a, b []float64) []float64 {
	r := make([]float64, len(a))
	for i := 0; i < len(a); i++ {
		r[i] = a[i] + b[i]
	}
	return r
}

func VecMul(vec []float64, s []float64) []float64 {
	r := make([]float64, len(vec))
	for i := 0; i < len(vec); i++ {
		r[i] = vec[i] * s[i]
	}
	return r
}

func VecDiv(vec []float64, s float64) []float64 {
	r := make([]float64, len(vec))
	for i := 0; i < len(vec); i++ {
		r[i] = vec[i] / s
	}
	return r
}

func VecExp(vec []float64) []float64 {
	r := make([]float64, len(vec))
	for i := 0; i < len(vec); i++ {
		r[i] = math.Exp(vec[i])
	}
	return r
}

func VecLog(vec []float64) []float64 {
	r := make([]float64, len(vec))
	for i := 0; i < len(vec); i++ {
		r[i] = math.Log(vec[i])
	}
	return r
}

func VecPow(vec []float64, p float64) []float64 {
	r := make([]float64, len(vec))
	for i := 0; i < len(vec); i++ {
		r[i] = math.Pow(vec[i], p)
	}
	return r
}

func VecSqrt(vec []float64) []float64 {
	r := make([]float64, len(vec))
	for i := 0; i < len(vec); i++ {
		r[i] = math.Sqrt(vec[i])
	}
	return r
}

func VecAbs(vec []float64) []float64 {
	r := make([]float64, len(vec))
	for i := 0; i < len(vec); i++ {
		r[i] = math.Abs(vec[i])
	}
	return r
}

func VecSigmoid(vec []float64) []float64 {
	r := make([]float64, len(vec))
	for i := 0; i < len(vec); i++ {
		r[i] = 1.0 / (1.0 + math.Exp(-vec[i]))
	}
	return r
}

func VecNormalize(vec []float64) []float64 {
	sum := VecSum(vec)
	if sum == 0 {
		return vec
	}
	return VecDiv(vec, sum)
}
