package giantbomb

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/corybuecker/steam-stats/fetcher"
	"github.com/corybuecker/steam-stats/steam"
)

type Search struct {
	Results []SearchResult `json:"results"`
}
type SearchResult struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Fetcher struct {
	GiantBombAPIKey string
}

func (fetcher *Fetcher) generateSearchURL(name string) string {
	return fmt.Sprintf("http://www.giantbomb.com/api/games/?api_key=%s&format=json&filter=name:%s&field_list=id,name", fetcher.GiantBombAPIKey, url.QueryEscape(name))
}

func (fetcher *Fetcher) FindOwnedGame(jsonfetcher fetcher.JSONFetcher, ownedGame *steam.OwnedGame) (*Search, error) {
	response, err := jsonfetcher.Fetch(fetcher.generateSearchURL(ownedGame.Name))
	time.Sleep(time.Second)

	if err != nil {
		return nil, err
	}

	var searchResults = new(Search)

	err = json.Unmarshal(response, searchResults)

	if err != nil {
		return nil, err
	}

	return searchResults, nil
}
