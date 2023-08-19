package cli

type Option struct {
	Label    string
	Action   func()
	Selected string
}

func (o *Option) setSelected(value string) {
	o.Selected = value
}
