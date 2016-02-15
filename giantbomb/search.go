package giantbomb

import (
	"fmt"
	"log"
	"net/url"

	"github.com/corybuecker/steam-stats/database"
	"github.com/corybuecker/steam-stats/fetcher"
	"github.com/corybuecker/steam-stats/ratelimiters"
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
	RateLimiter     *ratelimiters.GiantBombRateLimiter
}

func (fetcher *Fetcher) generateFetchURL(id int) string {
	return fmt.Sprintf("http://www.giantbomb.com/api/games/?filter=id:%d&api_key=%s&format=json&field_list=id,name,site_detail_url", id, fetcher.GiantBombAPIKey)
}

func (fetcher *Fetcher) generateSearchURL(name string) string {
	return fmt.Sprintf("http://www.giantbomb.com/api/games/?api_key=%s&format=json&filter=name:%s&field_list=id,name,site_detail_url", fetcher.GiantBombAPIKey, url.QueryEscape(name))
}

func (fetcher *Fetcher) FindGameByID(jsonfetcher fetcher.Interface, id int) error {
	if fetcher.RateLimiter == nil {
		fetcher.RateLimiter = &ratelimiters.GiantBombRateLimiter{}
	}

	log.Printf("fetching %d in the GiantBomb API", id)

	err := jsonfetcher.Fetch(fetcher.generateFetchURL(id), &fetcher.SearchResults)

	if err := fetcher.RateLimiter.ObeyRateLimit(); err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func (fetcher *Fetcher) FindOwnedGame(jsonfetcher fetcher.Interface, gameName string) error {
	if fetcher.RateLimiter == nil {
		fetcher.RateLimiter = &ratelimiters.GiantBombRateLimiter{}
	}

	log.Printf("searching for %s in the GiantBomb API", gameName)

	err := jsonfetcher.Fetch(fetcher.generateSearchURL(gameName), &fetcher.SearchResults)

	if err := fetcher.RateLimiter.ObeyRateLimit(); err != nil {
		return err
	}

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
