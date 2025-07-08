package mathz

import (
	"math"
	"testing"
)

func TestNormalize(t *testing.T) {
	tests := []struct {
		name string
		num  float64
		min  float64
		max  float64
		want float64
	}{{
		name: "Normalize(0, -1, 1) should return 0.5",
		num:  0,
		min:  -1,
		max:  1,
		want: 0.5,
	}, {
		name: "Normalize(0, 0, 0) should return math.NaN()",
		num:  0,
		min:  0,
		max:  0,
		want: math.NaN(),
	}, {
		name: "Normalize(0, 1, 2) should return -1",
		num:  0,
		min:  1,
		max:  2,
		want: -1,
	}, {
		name: "Normalize(math.MaxFloat64, -1, 1) should return (math.MaxFloat64 + 1) / 2",
		num:  math.MaxFloat64,
		min:  -1,
		max:  1,
		want: (math.MaxFloat64 + 1) / 2,
	}, {
		name: "Normalize(math.MaxFloat64, -math.MaxFloat64, math.MaxFloat64) should return math.NaN()",
		num:  math.MaxFloat64,
		min:  -math.MaxFloat64,
		max:  math.MaxFloat64,
		want: math.NaN(),
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Normalize(tt.num, tt.min, tt.max)
			if math.IsNaN(got) && math.IsNaN(tt.want) {
				return
			}
			if got != tt.want {
				t.Errorf("Normalize(%f, %f, %f) = %f; want %f", tt.num, tt.min, tt.max, got, tt.want)
			}
		})
	}
}
