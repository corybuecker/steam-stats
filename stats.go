package main

import (
	"log"
	"os"

	"github.com/corybuecker/steam-stats/configuration"
	"github.com/corybuecker/steam-stats/database"
	"github.com/corybuecker/steam-stats/fetcher"
	"github.com/corybuecker/steam-stats/giantbomb"
	"github.com/corybuecker/steam-stats/jobs"
	"github.com/corybuecker/steam-stats/steam"
	"github.com/corybuecker/steam-stats/storage"
	"github.com/dancannon/gorethink"
)

func main() {
	commandLineArgs := os.Args[1:]

	if len(commandLineArgs) != 1 {
		log.Fatalf("you must provide one argument, the connection string for RethinkDB")
	}

	session, err := gorethink.Connect(gorethink.ConnectOpts{Address: commandLineArgs[0]})
	if err != nil {
		log.Fatalln(err.Error())
	}
	var rethinkDB database.RethinkDB
	rethinkDB = database.RethinkDB{Session: session}

	var config configuration.Configuration
	config = configuration.Configuration{}

	if err := config.Load(&rethinkDB); err != nil {
		log.Fatal(err)
	}

	var steamFetcher = &steam.Fetcher{SteamAPIKey: config.SteamAPIKey, SteamID: config.SteamID}
	var giantBombFetcher = &giantbomb.Fetcher{GiantBombAPIKey: config.GiantBombAPIKey}
	var job = &jobs.Job{Fetcher: &fetcher.JSONFetcher{}, Database: &rethinkDB}

	storage.Setup(&rethinkDB, "videogames", []string{"ownedgames", "giantbomb"})

	job.OwnedGamesFetch(steamFetcher)
	job.OwnedGamesSearch(steamFetcher, giantBombFetcher)
	job.OwnedGamesFetchByID(steamFetcher, giantBombFetcher)

}
