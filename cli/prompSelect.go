package cli

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func PromptSelect(options []Option, label ...string) (int, string, error) {
	option := Option{}

	var items []string
	for _, opt := range options {
		items = append(items, opt.Label)
	}

	promptLabel := "Select an option"
	if len(label) > 0 {
		promptLabel = label[0]
	}

	prompt := promptui.Select{
		Label: promptLabel,
		Items: items,
	}

	index, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		return index, result, err
	}
	option.setSelected(result)

	selectedOption := findSelectedOption(result, options)
	if selectedOption != nil {
		selectedOption.Action()
	}

	return index, result, err
}

func findSelectedOption(label string, options []Option) *Option {
	for _, opt := range options {
		if opt.Label == label {
			return &opt
		}
	}
	return nil
}
