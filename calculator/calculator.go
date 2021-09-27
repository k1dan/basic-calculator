package calculator

import (
	customerrors "basic-calculator/custom-errors"
	"strconv"
)

func Calculate(input []string) (int, error) {
	result, err := strconv.Atoi(input[0])
	if err != nil {
		return 0, customerrors.CalculationError
	}
	for i := 1; i < (len(input) - 1); i++ {
		if input[i] == "+" {
			value, err := strconv.Atoi(input[i+1])
			if err != nil {
				return 0, customerrors.CalculationError
			}
			result = result + value
		} else if input[i] == "-" {
			value, err := strconv.Atoi(input[i+1])
			if err != nil {
				return 0, customerrors.CalculationError
			}
			result = result - value
		}
	}
	return result, nil
}
