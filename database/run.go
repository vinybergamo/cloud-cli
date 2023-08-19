package database

import (
	"fmt"
	"github.com/vinybergamo/cloud-cli/cli"
	"github.com/vinybergamo/cloud-cli/utils"
	"time"
)

func Run() {
	options := []cli.Option{
		{Label: "Create", Action: Create},
		{Label: "Link", Action: Link},
	}

	cli.PromptSelect(options)
}

func Create() {
	app := Database{}
	app.SetName()

	fmt.Println("Application created successfully: ", app.Name)
}

func Link() {
	database := Database{}

	err := utils.ValidateEmpty(database.Name, "")
	if err != nil {
		database.SelectDatabase()
	}

	err = utils.ValidateEmpty(database.App.Name, "")
	if err != nil {
		database.App.SelectApp()
	}

	fmt.Printf("Linking database %v with application %v.\n", database.Name, database.App.Name)
	time.Sleep(time.Second * 3)

	fmt.Printf("Database succesfully linked")
}
