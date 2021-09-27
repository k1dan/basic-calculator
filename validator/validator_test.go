package validator

import (
	customerrors "basic-calculator/custom-errors"
	"testing"
)

type testPair struct {
	input 	string
	result	string
	error	error
}

var testCases = []testPair{
	{"   14   +   5  -  4", "14+5-4", nil},
	{"14a -  2c  )", "", customerrors.InvalidCharsError},
	{"14+--24", "", customerrors.InvalidSequenceOfCharsError},
}

func TestValidateInput(t *testing.T) {
	for _, pair := range testCases {
		result, err := ValidateInput(pair.input)
		if result != pair.result {
			t.Error(
				"For", pair.input,
				"expected", pair.result,
				"got", result,
			)
		}
		if err != pair.error {
			t.Error(
				"For", pair.input,
				"expected error", pair.error,
				"got", err,
			)
		}
	}
}
