package actions

import (
	"log"

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

	if err = search(database); err != nil {
		return err
	}

	return nil
}

func search(db database.Interface) (err error) {
	var results []database.Game

	if results, err = db.GetAllGamesWithoutURL(); err != nil {
		return err
	}

	for _, game := range results {
		log.Printf("searching for %s", game.Name)
		searchResults, _ := wikipediasearch.Search(game.Name, true)

		if len(searchResults) > 0 {
			spew.Printf("found %v\n", searchResults)

			if err = db.UpsertIntField("steam_id", game.ID, bson.M{"wikipediaURL": searchResults[0].URL}); err != nil {
				return err
			}
		}
	}

	return nil
}
