package main

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"

	"github.com/codegangsta/cli"
	"github.com/corybuecker/goconfig"
	"github.com/corybuecker/steam-stats-fetcher/database"
	"github.com/corybuecker/steam-stats-fetcher/fetcher"
	"github.com/corybuecker/steam-stats-fetcher/giantbomb"
	"github.com/corybuecker/steam-stats-fetcher/jobs"
	"github.com/corybuecker/steam-stats-fetcher/steam"
)

func getDatabase(databaseHost string) (database.Interface, error) {
	session, err := mgo.Dial(databaseHost)
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	return &database.MongoDB{Collection: session.DB("test").C("steam_stats_fetcher")}, nil
}

func main() {
	databaseHost := "localhost"

	app := cli.NewApp()
	app.Name = "steam-stats-fetcher"

	var rethinkDB database.Interface
	var err error

	if rethinkDB, err = getDatabase(databaseHost); err != nil {
		log.Fatal(err)
	}

	var steamFetcher steam.Fetcher
	var giantBombFetcher giantbomb.Fetcher
	session, _ := mgo.Dial(databaseHost)

	goconfig.Get(session, "steam", &steamFetcher)
	goconfig.Get(session, "giantbomb", &giantBombFetcher)

	var job = &jobs.Job{Fetcher: &fetcher.JSONFetcher{}, Database: rethinkDB}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "host",
			Value:       "localhost",
			Usage:       "connection host for RethinkDB",
			Destination: &databaseHost,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "steam",
			Usage: "update all owned games from steam",
			Action: func(c *cli.Context) {
				job.OwnedGamesFetch(&steamFetcher)
			},
		},

		{
			Name:  "search",
			Usage: "search for the name of all owned games in GiantBomb",
			Action: func(c *cli.Context) {
				job.OwnedGamesSearch(&steamFetcher, &giantBombFetcher)
			},
		},

		{
			Name:  "fetch",
			Usage: "fetch all known games from GiantBomb",
			Action: func(c *cli.Context) {
				job.OwnedGamesFetchByID(&steamFetcher, &giantBombFetcher)
			},
		},
	}
	app.Run(os.Args)
}
