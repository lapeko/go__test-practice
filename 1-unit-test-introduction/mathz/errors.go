package mathz

import "errors"

const (
	SqrOfNegative = "cannot take sqrt of negative number"
)

var (
	ErrDivideByZero = errors.New("division by zero")
)
