package git

import (
	"github.com/manifoldco/promptui"
	"github.com/vinybergamo/cloud-cli/application"
	"github.com/vinybergamo/cloud-cli/cli"
)

type Git struct {
	Repository string
	App        application.App
}

func (g *Git) SetRepository() {
	prompt := promptui.Prompt{
		Label: "Repository",
	}

	result := cli.PromptInput(prompt)

	g.Repository = result
}
