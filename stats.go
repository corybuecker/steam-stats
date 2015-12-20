package main

import (
	"log"

	"github.com/corybuecker/steam-stats/configuration"
	"github.com/corybuecker/steam-stats/database"
	"github.com/corybuecker/steam-stats/fetcher"
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
	// var giantBombFetcher = &giantbomb.Fetcher{GiantBombAPIKey: config.GiantBombAPIKey}

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

	//
	// rethinkdb = &storage.RethinkDB{Name: "videogames", Tables: []string{"ownedgames", "giantbomb", "steam_giantbomb"}, Session: session}
	// var steamFetcher = &steam.Fetcher{SteamAPIKey: config.SteamAPIKey, SteamID: config.SteamID}
	// var giantBombFetcher = &giantbomb.Fetcher{GiantBombAPIKey: config.GiantBombAPIKey}
	// if err := rethinkdb.EnsureExists(); err != nil {
	// 	log.Fatalln(err)
	// }
	//
	// ownedGames, _ := steamFetcher.GetOwnedGames(&fetcher.JSONFetcher{})
	//
	// if err := rethinkdb.UpdateOwnedGames(ownedGames); err != nil {
	// 	log.Fatalln(err.Error())
	// }
	//
	// for _, ownedGame := range ownedGames.Response.Games {
	// 	log.Printf("searching Giantbomb for --- %s", ownedGame.Name)
	// 	var search *giantbomb.Search
	// 	search, err = giantBombFetcher.FindOwnedGame(&fetcher.JSONFetcher{}, &ownedGame)
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 	}
	// 	if len(search.Results) > 0 {
	// 		if err := rethinkdb.UpdateGiantBomb(search.Results); err != nil {
	// 			log.Println(err.Error())
	// 		}
	// 	}
	// }
	//
}
