package main

import (
	"github.com/urfave/cli"
	"log"
	"os"

	"tomjohnburton/helmplate/create"
)

func main() {
	app := cli.NewApp()
	app.Name = "helmplate"
	app.Usage = "Generate helm formatted resources ready for templating"

	app.Commands = []cli.Command{
		{
			Name: "create",
			Usage: "Create a specified resource formatted to your chart",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "chart",
					Usage:       "Loads metadata from `Chart.yaml`",
					FilePath:    "/path/to/Chart.yaml",
					Required:    true,
					TakesFile:   true,
				},
				cli.StringFlag{
					Name: "name",
					Usage: "Customize name of the resource",
					Required: false,
				},
			},
			Action: func(c *cli.Context) error {
				create.Create(c.Args().Get(0), c.String("chart"), c.String("name"))
				return nil
			},

		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
