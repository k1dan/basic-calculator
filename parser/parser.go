package parser

func ParseChars(input string) []string {
	var result []string
	tempChar := ""
	end := false
	index := 0
	for _, char := range input {
		if char == 43 || char == 45 || char == 10 {
			end = true
		} else {
			tempChar = tempChar + string(char)
		}
		if end == true {
			result = append(result, tempChar)
			if char != 10 {
				result = append(result, string(char))
			}
			tempChar = ""
			end = false
			index++
		}
	}
	result = append(result, string(input[len(input)-1]))
	return result
}
