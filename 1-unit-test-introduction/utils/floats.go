package utils

import "math"

func IsFloatsSoftEqualEps(epsilon float64, nums ...float64) bool {
	for i := 0; i < len(nums)-1; i++ {
		if math.Abs(nums[i]-nums[i+1]) >= epsilon {
			return false
		}
	}
	return true
}

func IsFloatsSoftEqual(nums ...float64) bool {
	return IsFloatsSoftEqualEps(1e-9, nums...)
}
