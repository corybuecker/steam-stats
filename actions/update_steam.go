package actions

import (
	"log"

	"github.com/corybuecker/steam-stats-fetcher/database"
	"github.com/corybuecker/steam-stats-fetcher/steam"
)

func UpdateSteam(steamFetcher *steam.Fetcher, database database.Interface) {
	if err := steamFetcher.GetOwnedGames(); err != nil {
		log.Fatalln(err.Error())
	}

	if err := steamFetcher.UpdateOwnedGames(database); err != nil {
		log.Fatalln(err.Error())
	}
}
