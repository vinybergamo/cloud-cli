package main

import (
	"github.com/vinybergamo/cloud-cli/application"
	"github.com/vinybergamo/cloud-cli/cli"
	"github.com/vinybergamo/cloud-cli/database"
	"github.com/vinybergamo/cloud-cli/git"
	"github.com/vinybergamo/cloud-cli/services"
)

func main() {
	options := []cli.Option{
		{Label: "Application", Action: application.Run, HideSelected: true},
		{Label: "Database", Action: database.Run, HideSelected: true},
		{Label: "Git", Action: git.Run},
		{Label: "Services", Action: services.Run, HideSelected: true},
	}

	cli.PromptSelect(options)
}
