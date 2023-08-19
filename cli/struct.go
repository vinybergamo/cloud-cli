package cli

import "github.com/manifoldco/promptui"

type Option struct {
	Label        string
	Action       func()
	HideSelected bool
	Selected     string
}

type Prompt struct {
	Input promptui.Prompt
}

func (o *Option) setSelected(value string) {
	o.Selected = value
}
