package database

import (
	"errors"
	"github.com/manifoldco/promptui"
	"github.com/vinybergamo/cloud-cli/application"
	"github.com/vinybergamo/cloud-cli/cli"
	"github.com/vinybergamo/cloud-cli/utils"
)

type Database struct {
	Name string
	App  application.App
}

func (d *Database) SetName() {
	prompt := promptui.Prompt{
		Label:   "Database Name",
		Default: utils.FormatString(utils.GenerateRandomName()),
		Validate: func(s string) error {
			message := "Invalid database name: Min length 3"
			if len(s) < 3 {
				return errors.New(message)
			}
			return nil
		},
	}

	result := cli.PromptInput(prompt)

	name := utils.FormatString(result)
	d.Name = name
}

func (d *Database) SelectDatabase() {
	var options []cli.Option

	names := []string{"Database 1", "Database 2"}
	for _, name := range names {
		options = append(options, cli.Option{
			Label: name,
			Action: func() {
				d.Name = name
			},
		})
	}

	cli.PromptSelect(options, "Select database")
}
