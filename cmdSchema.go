package main

import "strconv"

type Validator func(string) (bool, string) // success, msg
type TextInputSchema struct {
	header string
	placeholder string
	footer string
	validators []Validator
}

func initialize() *TextInputSchema {
	validators := make([]Validator, 1)
	validators[0] = func (input string) (bool, string) {
		_, err := strconv.Atoi(input)
		if err != nil {
			return false, "Value must be an integer."
		}
		return true, "Slay"
	}
	test := &TextInputSchema{
		header: "Enter an input file path:", 
		placeholder: "./video.mp4",
		footer: "(press esc to quit)",
		validators: validators,
	}
	return test
}

