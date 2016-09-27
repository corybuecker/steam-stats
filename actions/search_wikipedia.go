package actions

import (
	"github.com/corybuecker/steamfetcher/database"
	"github.com/corybuecker/wikipediasearch"
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/mgo.v2/bson"
)

func SearchWikipedia(databaseHost string) error {
	var database database.Interface
	var err error

	if database, err = GetDatabase(databaseHost); err != nil {
		return err
	}

	results, _ := database.GetAllGames()

	for _, game := range results {
		searchResults, _ := wikipediasearch.Search(game.Name, true)
		if len(searchResults) > 0 {
			spew.Printf("found %v\n", searchResults)
			database.UpsertIntField("steam_id", game.ID, bson.M{"wikipediaURL": searchResults[0].URL})
		}
	}

	return nil
}
