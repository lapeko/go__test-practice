package utils

import (
	"fmt"
	"testing"
)

func TestIsFloatsSoftEqual(t *testing.T) {
	fName := "IsFloatsSoftEqual"
	tests := []struct {
		name string
		num1 float64
		num2 float64
		want bool
	}{{
		name: fmt.Sprintf("%s(1, 1 + 1e-9) should return false", fName),
		num1: 1,
		num2: 1 + 1e-9,
		want: false,
	}, {
		name: fmt.Sprintf("%s(1, 1 + 1e-10) should return true", fName),
		num1: 1,
		num2: 1 + 1e-10,
		want: true,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsFloatsSoftEqual(tt.num1, tt.num2)
			if got != tt.want {
				t.Fatalf("%s(%g, %g) = %t; want %t", fName, tt.num1, tt.num2, got, tt.want)
			}
		})
	}
}

func TestIsFloatsSoftEqualEps(t *testing.T) {
	fName := "IsFloatsSoftEqualEps"
	tests := []struct {
		name    string
		epsilon float64
		num1    float64
		num2    float64
		want    bool
	}{{
		name:    fmt.Sprintf("%s(1e-2, 1, 1 + 1e-2) should return false", fName),
		epsilon: 1e-2,
		num1:    1,
		num2:    1 + 1e-2,
		want:    false,
	}, {
		name:    fmt.Sprintf("%s(1e-1, 1, 1 + 1e-2) should return true", fName),
		epsilon: 1e-1,
		num1:    1,
		num2:    1 + 1e-2,
		want:    true,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsFloatsSoftEqualEps(tt.epsilon, tt.num1, tt.num2)
			if got != tt.want {
				t.Fatalf("%s(%g, %g, %g) = %t; want %t", fName, tt.epsilon, tt.num1, tt.num2, got, tt.want)
			}
		})
	}
}
