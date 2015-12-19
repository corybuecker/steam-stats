package giantbomb

import (
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

func (fetcher *Fetcher) FindOwnedGame(jsonfetcher fetcher.JSONFetcherInterface, ownedGame *steam.OwnedGame) (*Search, error) {
	var data Search = Search{}

	err := jsonfetcher.Fetch(fetcher.generateSearchURL(ownedGame.Name), &data)

	time.Sleep(time.Second)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
