package cli

import (
	"fmt"
	"os"
)

func SelectOption(option string, options map[string]func(), callback func()) {
	if function, ok := options[option]; ok {
		function()
	} else {
		fmt.Println("Invalid option")

		if callback != nil {
			callback()
		} else {
			os.Exit(1)
		}
	}
}
