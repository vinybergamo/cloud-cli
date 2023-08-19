package ports

import (
	"fmt"
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

	err := utils.ValidateEmpty(port.App.Name, "")
	if err != nil {
		port.App.SelectApp()
	}

	port.SetPort()
	fmt.Printf("Setting Port %v for application %v.", port.Port, port.App.Name)
}
