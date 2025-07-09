package utils

import (
	"errors"
	"fmt"
	"testing"
)

func TestStringifyError(t *testing.T) {
	fName := "StringifyError"
	err := errors.New("test error")
	tests := []struct {
		name string
		err  error
		want string
	}{{
		name: fmt.Sprintf("%s with provided error should return stringified error", fName),
		err:  err,
		want: fmt.Sprintf("error(%q)", err.Error()),
	}, {
		name: fmt.Sprintf("%s with nil error should return \"nil\"", fName),
		err:  nil,
		want: "nil",
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringifyError(tt.err)
			if got != tt.want {
				t.Errorf("%s(%v) = %s; want: %s", fName, tt.err, got, tt.want)
			}
		})
	}
}
