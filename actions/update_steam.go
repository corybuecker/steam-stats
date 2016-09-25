package actions

import (
	"github.com/corybuecker/steamfetcher/database"
	"github.com/corybuecker/steamfetcher/steam"
)

func UpdateSteam(databaseHost string) error {
	var database database.Interface
	var err error
	var steamFetcher = &steam.Fetcher{}

	if database, err = GetDatabase(databaseHost); err != nil {
		return err
	}

	if err = steamFetcher.UpdateOwnedGames(database); err != nil {
		return err
	}

	return nil
}
