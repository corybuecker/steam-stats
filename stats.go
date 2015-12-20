package main

import (
	"log"

	"github.com/corybuecker/steam-stats/configuration"
	"github.com/corybuecker/steam-stats/database"
	"github.com/corybuecker/steam-stats/fetcher"
	"github.com/corybuecker/steam-stats/giantbomb"
	"github.com/corybuecker/steam-stats/steam"
	"github.com/corybuecker/steam-stats/storage"
	"github.com/dancannon/gorethink"
)

func main() {

	var config = new(configuration.Configuration)

	if err := config.Load("./config.json"); err != nil {
		log.Fatal(err)
	}

	var steamFetcher = &steam.Fetcher{SteamAPIKey: config.SteamAPIKey, SteamID: config.SteamID}
	var giantBombFetcher = &giantbomb.Fetcher{GiantBombAPIKey: config.GiantBombAPIKey}

	session, err := gorethink.Connect(gorethink.ConnectOpts{Address: "localhost:28015"})

	if err != nil {
		log.Fatalln(err.Error())
	}

	rethinkDB := database.RethinkDB{Session: session}

	storage.Setup(&rethinkDB, "videogames", []string{"ownedgames", "giantbomb"})

	if err := steamFetcher.GetOwnedGames(&fetcher.JSONFetcher{}); err != nil {
		log.Fatalln(err.Error())
	}

	if err := steamFetcher.UpdateOwnedGames(&rethinkDB); err != nil {
		log.Fatalln(err.Error())
	}

	var ownedGamesWithoutGiantBombID []string

	if ownedGamesWithoutGiantBombID, err = steamFetcher.FetchOwnedGames(&rethinkDB); err != nil {
		log.Fatalln(err.Error())
	}

	for _, ownedGame := range ownedGamesWithoutGiantBombID {
		if err := giantBombFetcher.FindOwnedGame(&fetcher.JSONFetcher{}, ownedGame); err != nil {
			log.Println("errr")
		}
		if err := giantBombFetcher.UpdateFoundGames(&rethinkDB); err != nil {
			log.Println("errr")
		}
	}
}
