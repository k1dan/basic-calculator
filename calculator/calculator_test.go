package calculator

import (
	customerrors "basic-calculator/custom-errors"
	"testing"
)

type testPair struct {
	input 	[]string
	result	int
	error	error
}

var testCases = []testPair{
	{[]string{"14", "+", "4", "-", "5"}, 13, nil},
	{[]string{"14", "-", "10", "-", "5"}, -1, nil},
	{[]string{"20a", "-", "4", "+", "1"}, 0, customerrors.CalculationError},
}

func TestCalculate(t *testing.T) {
	for _, pair := range testCases {
		result, err := Calculate(pair.input)
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
