package main

import (
	"github.com/vinybergamo/cloud-cli/application"
	"github.com/vinybergamo/cloud-cli/cli"
	"github.com/vinybergamo/cloud-cli/services/ports"
)

func main() {
	options := []cli.Option{
		{Label: "Application", Action: application.Run},
		{Label: "Ports", Action: ports.Run},
	}

	cli.PromptSelect(options)
}
