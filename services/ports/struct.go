package ports

import (
	"errors"
	"github.com/manifoldco/promptui"
	"github.com/vinybergamo/cloud-cli/cli"
	"strconv"
)

type Port struct {
	Port string
}

func (p *Port) SetPort() {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid port")
		}
		if len(input) < 4 {
			return errors.New("Invalid port: Min length 4")
		} else if len(input) > 5 {
			return errors.New("Invalid port: Max length 5")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Application port",
		Default:  "3000",
		Validate: validate,
	}

	result := cli.PromptInput(prompt)

	p.Port = result
}
