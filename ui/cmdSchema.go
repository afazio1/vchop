package ui

import "strconv"

type Steps struct {
	Steps map[Step]TextInputSchema
}
type Step string
const (
	Input Step = "input"
	Duration Step = "duration"
	Noise Step = "noise"
	Output Step = "output"
)
type Validator func(string) (bool, string) // success, msg
type TextInputSchema struct {
	header string
	placeholder string
	footer string
	validators []Validator
}

func Initialize() *Steps {
	validators := make([]Validator, 1)
	validators[0] = func (input string) (bool, string) {
		_, err := strconv.Atoi(input)
		if err != nil {
			return false, "Value must be an integer."
		}
		return true, "Slay"
	}
	steps := &Steps{
		map[Step]TextInputSchema {
			Input: {
				header: "Enter an input file path:",
				placeholder: "./video.mp4",
				footer: "(press esc to quit)",
				validators: validators,
			},
			Output: {
				header: "Enter an output file path:",
				placeholder: "./output.mp4",
				footer: "(press esc to quit)",
				validators: validators,
			},
			Noise: {
				header: "Enter an noise level for silence:",
				placeholder: "-30",
				footer: "(press esc to quit)",
				validators: validators,
			},
		},
	}
	return steps
}

