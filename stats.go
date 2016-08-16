package main

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"

	"github.com/corybuecker/jsonfetcher"
	"github.com/corybuecker/mgoconfig"
	"github.com/corybuecker/steam-stats-fetcher/database"
	"github.com/corybuecker/steam-stats-fetcher/giantbomb"
	"github.com/corybuecker/steam-stats-fetcher/jobs"
	"github.com/corybuecker/steam-stats-fetcher/steam"
	"github.com/urfave/cli"
)

func getMongoSession(databaseHost string) (*mgo.Session, error) {
	session, err := mgo.Dial(databaseHost)
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	return session, nil
}

func main() {
	databaseHost := "localhost"

	app := cli.NewApp()
	app.Name = "steam-stats-fetcher"

	var mongoSession *mgo.Session
	var mongoDatabase database.Interface
	var err error

	if mongoSession, err = getMongoSession(databaseHost); err != nil {
		log.Fatal(err)
	}

	mongoDatabase = &database.MongoDB{Collection: mongoSession.DB("steam_stats_fetcher").C("games")}

	var steamFetcher steam.Fetcher
	var giantBombFetcher giantbomb.Fetcher

	mgoconfig.Get(mongoSession, "steam_stats_fetcher", "steam", &steamFetcher)
	mgoconfig.Get(mongoSession, "steam_stats_fetcher", "giantbomb", &giantBombFetcher)

	steamFetcher.Jsonfetcher = &jsonfetcher.Jsonfetcher{}
	giantBombFetcher.Jsonfetcher = &jsonfetcher.Jsonfetcher{}

	var job = &jobs.Job{Database: mongoDatabase}

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
			Action: func(c *cli.Context) error {
				job.OwnedGamesFetch(&steamFetcher)
				return nil
			},
		},

		{
			Name:  "search",
			Usage: "search for the name of all owned games in GiantBomb",
			Action: func(c *cli.Context) error {
				job.OwnedGamesSearch(&steamFetcher, &giantBombFetcher)
				return nil
			},
		},

		{
			Name:  "fetch",
			Usage: "fetch all known games from GiantBomb",
			Action: func(c *cli.Context) error {
				job.OwnedGamesFetchByID(&steamFetcher, &giantBombFetcher)
				return nil
			},
		},
	}
	app.Run(os.Args)
}
