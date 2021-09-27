package parser

import (
	"testing"
)

type testPair struct {
	input 	string
	result	[]string
}

var testCases = []testPair{
	{"14+5-4",[]string{"14", "+", "5", "-", "4"}},
	{"20-4+1",[]string{"20", "-", "4", "+", "1"}},
}

func TestParseChars(t *testing.T) {
	for _, pair := range testCases {
		result := ParseChars(pair.input)
		for i:= 0; i < len(result)-1; i++ {
			if result[i] != pair.result[i] {
				t.Error(
					"For", pair.input,
					"expected", pair.result,
					"got", result,
				)
			}
		}
	}
}
