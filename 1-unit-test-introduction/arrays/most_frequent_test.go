package arrays

import (
	"fmt"
	"testing"
)

type test struct {
	name  string
	array []int
	want  int
}

func (t *test) addPrefix2Name(name string) {
	t.name = fmt.Sprintf("%s%s", name, t.name)
}

var tests = []test{{
	name:  "([1]) should return 1",
	array: []int{1},
	want:  1,
}, {
	name:  "([-1, 1, -1, 1]) should return -1",
	array: []int{-1, 1, -1, 1},
	want:  -1,
}}

var funcs = map[string]func([]int) int{
	"MostFrequent":      MostFrequent,
	"MostFrequentNaive": MostFrequentNaive,
}

func TestAllMostFrequentImplementations(t *testing.T) {
	for fName, fn := range funcs {
		for _, tt := range tests {
			tt.addPrefix2Name(fName)
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := fn(tt.array)
				if got != tt.want {
					t.Errorf("%s(%v) = %d; want %d", fName, tt.array, got, tt.want)
				}
			})
		}
	}
}
