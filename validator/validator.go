package validator

import (
	customerrors "basic-calculator/custom-errors"
	"strings"
)

func ValidateInput(input string) (string, error) {
	input = strings.ReplaceAll(input, " ", "")
	if !isCharsValid(input) {
		return "", customerrors.InvalidCharsError
	}
	if !isSequenceValid(input) {
		return "", customerrors.InvalidSequenceOfCharsError
	}
	input = strings.TrimSuffix(input, "\n")
	return input, nil
}

func isCharsValid(input string) bool {
	for _, char := range input {
		if (char < 48 || char > 57) && char != 43 && char != 45 && char != 10 {
			return false
		}
	}
	return true
}

func isSequenceValid(input string) bool {
	if input[0] == 43 || input[len(input)-1] == 43 || input[0] == 45 || input[len(input)-1] == 45 {
		return false
	}
	for i := 1; i < (len(input) - 2); i++ {
		if (input[i] == 43 || input[i] == 45) && (input[i+1] == 43 || input[i+1] == 45) {
			return false
		}
	}
	return true
}
