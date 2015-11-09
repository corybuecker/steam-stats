package main

import (
	"log"

	"github.com/corybuecker/steam-stats/configuration"
	"github.com/corybuecker/steam-stats/database"
	"github.com/corybuecker/steam-stats/fetcher"
	"github.com/dancannon/gorethink"
)

func main() {
	var db *database.DB
	var config = new(configuration.Configuration)
	if err := config.Load("./config.json"); err != nil {
		log.Fatal(err)
	}
	session, err := gorethink.Connect(gorethink.ConnectOpts{Address: "localhost:28015"})
	if err != nil {
		log.Fatalln(err.Error())
	}

	db = &database.DB{Name: "videogames", Tables: []string{"steam", "mygames", "giantbomb"}, Session: session}
	var steamFetcher = &fetcher.SteamFetcher{Storage: db, SteamAPIKey: config.SteamAPIKey, SteamID: config.SteamID, GiantBombAPIKey: config.GiantBombAPIKey}
	if err := db.EnsureExists(); err != nil {
		log.Fatalln(err.Error())
	}

	ownedGames, _ := steamFetcher.GetOwnedGames(fetcher.JSONFetcher{})

	log.Println(ownedGames.Response)

}
