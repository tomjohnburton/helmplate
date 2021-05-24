package main

import (
	"github.com/urfave/cli"
	"log"
	"os"

	"github.com/tomjohnburton/helmplate/create"
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
			},
			Action: func(c *cli.Context) error {
				return nil
			},



		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
