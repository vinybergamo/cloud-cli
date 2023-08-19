package git

import (
	"fmt"
	"github.com/vinybergamo/cloud-cli/cli"
	"github.com/vinybergamo/cloud-cli/utils"
	"time"
)

func Run() {
	options := []cli.Option{
		{Label: "Sync", Action: Sync},
	}

	cli.PromptSelect(options)
}

func Sync() {
	git := Git{}

	git.SetRepository()

	err := utils.ValidateEmpty(git.App.Name, "")
	if err != nil {
		git.App.SelectApp()
	}

	fmt.Printf("Syncing %v with %v\n", git.App.Name, git.Repository)
	time.Sleep(time.Second * 4)
	fmt.Println("Git synced successfully.")
}
