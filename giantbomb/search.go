package giantbomb

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/corybuecker/steam-stats/database"
	"github.com/corybuecker/steam-stats/fetcher"
)

type Search struct {
	Results []SearchResult `json:"results"`
}
type SearchResult struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"site_detail_url"`
}

type Fetcher struct {
	GiantBombAPIKey string
	SearchResults   Search
}

func (fetcher *Fetcher) generateSearchURL(name string) string {
	return fmt.Sprintf("http://www.giantbomb.com/api/games/?api_key=%s&format=json&filter=name:%s&field_list=id,name,site_detail_url", fetcher.GiantBombAPIKey, url.QueryEscape(name))
}

func (fetcher *Fetcher) FindOwnedGame(jsonfetcher fetcher.Interface, gameName string) error {
	log.Printf("searching for %s in the GiantBomb API", gameName)

	err := jsonfetcher.Fetch(fetcher.generateSearchURL(gameName), &fetcher.SearchResults)

	time.Sleep(time.Second)

	if err != nil {
		return err
	}

	return nil
}

func (fetcher *Fetcher) UpdateFoundGames(database database.Interface) error {
	for _, foundGame := range fetcher.SearchResults.Results {
		foundGameMap := map[string]interface{}{
			"id":   foundGame.ID,
			"name": foundGame.Name,
			"url":  foundGame.URL,
		}

		if err := database.Upsert("videogames", "giantbomb", foundGameMap); err != nil {
			return err
		}
	}
	return nil
}
