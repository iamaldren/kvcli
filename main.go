package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func init() {
	cli.VersionFlag = cli.BoolFlag{
		Name: "version, v",
		Usage: "print only the version",
	}
}

func cliInfo(app *cli.App) {
	app.Name = "Simple Key-Value with Command Line Interface"
	app.Usage = "An in-memory storage app with CLI"
	app.Author = "@iamaldren"
	app.Version = "1.0.0"
}

func commands(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name: "test",
			Aliases: []string{"t"},
			Usage: "threading the waters",
			Action: func(c *cli.Context) {
				fmt.Println("Woah! You created your first cli commands!")
			},
		},
	}
}

func main() {
	app := cli.NewApp()

	cliInfo(app)
	commands(app)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}