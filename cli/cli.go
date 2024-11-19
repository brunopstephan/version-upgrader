package cli

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func GetOption() int {
	var opt string
	fmt.Println("Select the type of change you want to make:")
	prompt := &survey.Select{
		Message: "Chose a option:",
		Options: []string{
			"Breaking Change (updates first number)",
			"Release/Feature (updates second number)",
			"Fix/Other (updates third number)",
			"Don't update",
		},
	}

	err := survey.AskOne(prompt, &opt)
	if err != nil {
		fmt.Println("Prompt failed:", err)
		return -1
	}

	indexMap := map[string]int{
		"Breaking Change (updates first number)":  0,
		"Release/Feature (updates second number)": 1,
		"Fix/Other (updates third number)":        2,
		"Don't update":                            -1,
	}

	return indexMap[opt]
}
