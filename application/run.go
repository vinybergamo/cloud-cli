package application

import (
	"fmt"
	"github.com/vinybergamo/cloud-cli/cli"
)

func Run() {
	options := []cli.Option{
		{Label: "Create", Action: Create},
	}

	cli.PromptSelect(options)
}

func Create() {
	app := App{}
	app.SetName()

	fmt.Println("Application created successfully: ", app.Name)
}
