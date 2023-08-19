package cli

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
)

func PromptInput(prompt promptui.Prompt) string {

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("[ CLI INPUT ERROR ] %v", err)
		os.Exit(1)
	}

	return result
}
