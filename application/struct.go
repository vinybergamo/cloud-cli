package application

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/vinybergamo/cloud-cli/cli"
	"github.com/vinybergamo/cloud-cli/utils"
)

type App struct {
	Name string
}

func (a *App) SetName() {
	prompt := promptui.Prompt{
		Label:   "Application Name",
		Default: utils.FormatString(utils.GenerateRandomName()),
		Validate: func(s string) error {
			message := "Invalid application name: Min length 3"
			if len(s) < 3 {
				return errors.New(message)
			}
			return nil
		},
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	name := utils.FormatString(result)
	a.Name = name
}

func (a *App) SelectApp() {
	var options []cli.Option

	names := []string{"App 1", "App 2"}
	for _, name := range names {
		options = append(options, cli.Option{
			Label: name,
			Action: func() {
				a.Name = name
			},
		})
	}

	cli.PromptSelect(options, "Select an application")
}
