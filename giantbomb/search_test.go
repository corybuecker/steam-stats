package giantbomb

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/corybuecker/steam-stats/test"
)

var fakeDatabase test.FakeDatabase

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
	if err := gbFetcher.FindOwnedGame(&FakeFetcher{}, "gamename"); err != nil {
		t.Error(err)
	}
	if gbFetcher.SearchResults.Results[0].ID != 1 {
		t.Error("expected ID of 1")
	}
}

func TestDataUpdating(t *testing.T) {
	if err := gbFetcher.FindOwnedGame(&FakeFetcher{}, "gamename"); err != nil {
		t.Error(err)
	}
	if err := gbFetcher.UpdateFoundGames(&fakeDatabase); err != nil {
		t.Error(err)
	}
	if fakeDatabase.Entry["name"] != "foundgame" {
		t.Error("expected the entry to have an ID of 10")
	}
}
