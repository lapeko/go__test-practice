package mathz

import "math"

func Normalize(x, min, max float64) float64 {
	if max == min {
		return math.NaN()
	}
	return (x - min) / (max - min)
}
