package customerrors

import "errors"

var (
	InvalidSequenceOfCharsError = errors.New("Input contains unacceptable sequence of characters")
	InvalidCharsError = errors.New("Input contains unacceptable characters")
	CalculationError = errors.New("Unable to make calculations")
)
