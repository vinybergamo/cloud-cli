package ports

import (
	"fmt"
	"github.com/vinybergamo/cloud-cli/application"
	"github.com/vinybergamo/cloud-cli/cli"
	"github.com/vinybergamo/cloud-cli/utils"
)

func Run() {
	options := []cli.Option{
		{Label: "Set", Action: Set},
	}

	cli.PromptSelect(options)
}

func Set() {
	port := Port{}
	app := application.App{}

	err := utils.ValidateEmpty(app.Name, "")
	if err != nil {
		app.SelectApp()
	}

	port.SetPort()
	fmt.Printf("Setting Port %v for application %v.", port.Port, app.Name)
}
