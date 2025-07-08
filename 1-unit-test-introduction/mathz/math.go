package mathz

import "math"

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func Normalize(x, min, max float64) float64 {
	if max == min {
		return math.NaN()
	}
	return (x - min) / (max - min)
}

func MustSqrt(num float64) float64 {
	if num < 0 {
		panic(SqrOfNegative)
	}
	return math.Sqrt(num)
}
