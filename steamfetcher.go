package main

import (
	"os"

	"github.com/corybuecker/steamfetcher/actions"
	"github.com/urfave/cli"
)

func main() {
	var databaseHost string

	app := cli.NewApp()

	app.Name = "steamfetcher"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "host",
			Value:       "localhost",
			Usage:       "connection host for MongoDB",
			Destination: &databaseHost,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "steam",
			Usage: "update all owned games from steam",
			Action: func(c *cli.Context) error {
				if err := actions.UpdateSteam(databaseHost); err != nil {
					return err
				}
				return nil
			},
		},

		{
			Name:  "wikipedia",
			Usage: "fetch games from wikipedia",
			Action: func(c *cli.Context) error {
				if err := actions.SearchWikipedia(databaseHost); err != nil {
					return err
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
