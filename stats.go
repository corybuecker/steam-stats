package main

import (
	"log"

	"github.com/corybuecker/steam-stats/configuration"
	"github.com/corybuecker/steam-stats/fetcher"
	"github.com/corybuecker/steam-stats/storage"
	"github.com/dancannon/gorethink"
)

func main() {
	var rethinkdb *storage.RethinkDB
	var config = new(configuration.Configuration)
	if err := config.Load("./config.json"); err != nil {
		log.Fatal(err)
	}
	session, err := gorethink.Connect(gorethink.ConnectOpts{Address: "localhost:28015"})
	if err != nil {
		log.Fatalln(err.Error())
	}

	rethinkdb = &storage.RethinkDB{Name: "videogames", Tables: []string{"ownedgames", "giantbomb"}, Session: session}
	var steamFetcher = &fetcher.SteamFetcher{SteamAPIKey: config.SteamAPIKey, SteamID: config.SteamID, GiantBombAPIKey: config.GiantBombAPIKey}
	if err := rethinkdb.EnsureExists(); err != nil {
		log.Fatalln(err)
	}

	ownedGames, _ := steamFetcher.GetOwnedGames(fetcher.JSONFetcher{})
	if err := rethinkdb.UpdateOwnedGames(ownedGames); err != nil {
		log.Fatalln(err.Error())
	}

	log.Println(ownedGames.Response)

}
