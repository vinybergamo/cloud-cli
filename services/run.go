package services

import (
	"github.com/vinybergamo/cloud-cli/cli"
	"github.com/vinybergamo/cloud-cli/services/ports"
)

func Run() {
	options := []cli.Option{
		{Label: "Ports", Action: ports.Run},
	}

	cli.PromptSelect(options)
}
