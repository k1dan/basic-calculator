package main

import (
	"basic-calculator/calculator"
	"basic-calculator/parser"
	"basic-calculator/validator"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}

	text, err = validator.ValidateInput(text)
	if err != nil {
		log.Fatalln(err)
	}

	data := parser.ParseChars(text)
	result, err := calculator.Calculate(data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(result)
}

