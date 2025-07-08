package mathz

import (
	"errors"
	"fmt"
	"github.com/lapeko/go__test-practice/utils"
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
			if !utils.IsFloatsSoftEqual(got, tt.want) {
				t.Errorf("Normalize(%f, %f, %f) = %f; want %f", tt.num, tt.min, tt.max, got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	fName := "TestDivide"
	tests := []struct {
		name      string
		a         float64
		b         float64
		want      float64
		wantError error
	}{{
		name:      fmt.Sprintf("%s(%g, %g) should return %g, error.New(%q)", fName, 0.0, 0.0, 0.0, ErrDivideByZero.Error()),
		a:         0,
		b:         0,
		want:      0,
		wantError: ErrDivideByZero,
	}, {
		name:      fmt.Sprintf("%s(%g, %g) should return %g, %v", fName, 1.0, 2.0, 0.5, nil),
		a:         1,
		b:         2,
		want:      0.5,
		wantError: nil,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotError := Divide(tt.a, tt.b)
			if got != tt.want || !errors.Is(gotError, tt.wantError) {
				t.Errorf(
					"%s(%g, %g)\n\tgot: %g, %v;\n\twant: %g, %s",
					fName,
					tt.a,
					tt.b,
					got,
					utils.StringifyError(gotError),
					tt.want,
					utils.StringifyError(tt.wantError),
				)
			}
		})
	}
}

func TestMustSqrtFail(t *testing.T) {
	fName := "MustSqrt"
	payload := -1.0
	testName := fmt.Sprintf("%s(%g)", fName, payload)

	t.Run(testName, func(t *testing.T) {
		defer func() {
			got := recover()
			if got == nil {
				t.Fatalf("%s\n\tgot panic:\tnil\n\twant panic:\t%q", testName, SqrOfNegative)
			} else if got != SqrOfNegative {
				t.Fatalf("%s\n\tgot panic:\t%q\n\twant panic:\t%q", testName, got, SqrOfNegative)
			}
		}()
		MustSqrt(payload)
	})
}

func TestMustSqrt(t *testing.T) {
	tests := []struct {
		name string
		num  float64
		want float64
	}{{
		name: "MustSqrt(0)",
		num:  0,
		want: 0,
	}, {
		name: "MustSqrt(4)",
		num:  4,
		want: 2,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if got := recover(); got != nil {
					t.Fatalf("%s\n\tgot: unexpected panic(%q)\n\twant: %g", tt.name, got, tt.want)
				}
			}()
			got := MustSqrt(tt.num)
			if got != tt.want {
				t.Fatalf("%s\n\tgot: %g\n\twant: %g", tt.name, got, tt.want)
			}
		})
	}
}
