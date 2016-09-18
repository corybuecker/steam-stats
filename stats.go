package main

import (
	"os"

	"github.com/corybuecker/steam-stats-fetcher/actions"
	"github.com/urfave/cli"
)

func main() {
	var databaseHost string

	app := cli.NewApp()

	app.Name = "steam-stats-fetcher"

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
	}

	app.Run(os.Args)
}
