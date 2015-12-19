package giantbomb

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/corybuecker/steam-stats/steam"
)

type FakeFetcher struct{}

func (fetcher *FakeFetcher) Fetch(url string, data interface{}) error {
	var sampleResponse string = "{\"results\": [{\"id\": 1, \"name\": \"foundgame\"}]}"
	if err := json.Unmarshal([]byte(sampleResponse), data); err != nil {
		return err
	}
	return nil
}

var gbFetcher Fetcher

func init() {
	gbFetcher = Fetcher{GiantBombAPIKey: "API KEY"}
}

func TestURLIncludesAPIKey(t *testing.T) {
	if strings.Contains(gbFetcher.generateSearchURL("test"), "API KEY") != true {
		t.Error("expected URL to contain API KEY")
	}
}

func TestURLIncludesSearchName(t *testing.T) {
	if strings.Contains(gbFetcher.generateSearchURL("test"), "test") != true {
		t.Error("expected URL to contain search term")
	}
}

func TestDataMarshalling(t *testing.T) {
	foundGames, _ := gbFetcher.FindOwnedGame(&FakeFetcher{}, &steam.OwnedGame{Name: "gamename"})
	if foundGames.Results[0].ID != 1 {
		t.Error("expected ID of 1")
	}
}
