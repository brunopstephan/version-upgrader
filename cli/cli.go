package cli

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func GetOption() int {
	var opt string
	fmt.Println("Select the type of change you want to make:")
	prompt := &survey.Select{
		Message: "Choose an option:",
		Options: []string{
			"Breaking Change (X.0.0)",
			"Release/Feature (0.X.0)",
			"Fix/Other (0.0.X)",
			"Don't update",
		},
	}

	err := survey.AskOne(prompt, &opt)
	if err != nil {
		fmt.Println("Prompt failed:", err)
		return -1
	}

	indexMap := map[string]int{
		"Breaking Change (X.0.0)":  0,
		"Release/Feature (0.X.0)": 1,
		"Fix/Other (0.0.X)":        2,
		"Don't update":                            -1,
	}

	return indexMap[opt]
}
