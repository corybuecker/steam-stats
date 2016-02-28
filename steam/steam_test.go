package steam

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/corybuecker/steam-stats-fetcher/test"
)

type FakeFetcher struct{}

func (fetcher *FakeFetcher) Fetch(url string, data interface{}) error {
	var sampleResponse string = "{\"response\": {\"games\": [{\"appid\": 10, \"playtime_forever\": 32}]}}"
	if err := json.Unmarshal([]byte(sampleResponse), data); err != nil {
		return err
	}
	return nil
}

var steamFetcher Fetcher
var fakeDatabase test.FakeDatabase

func init() {
	fakeDatabase = test.FakeDatabase{}
	steamFetcher = Fetcher{SteamAPIKey: "API KEY", SteamID: "ID"}
}

func TestURLIncludesAPIKey(t *testing.T) {
	if strings.Contains(steamFetcher.generateURL(), "API KEY") != true {
		t.Error("expected URL to contain API KEY")
	}
}
func TestURLIncludesSteamID(t *testing.T) {
	if strings.Contains(steamFetcher.generateURL(), "ID") != true {
		t.Error("expected URL to contain Steam ID")
	}
}

func TestDataMarshalling(t *testing.T) {
	if err := steamFetcher.GetOwnedGames(&FakeFetcher{}); err != nil {
		t.Error(err)
	}
	if steamFetcher.OwnedGames.Response.Games[0].ID != 10 {
		t.Error("expected ID of 10")
	}
}

func TestDataUpdating(t *testing.T) {
	if err := steamFetcher.GetOwnedGames(&FakeFetcher{}); err != nil {
		t.Error(err)
	}
	if err := steamFetcher.UpdateOwnedGames(&fakeDatabase); err != nil {
		t.Error(err)
	}
	if fakeDatabase.Entry["id"] != 10 {
		t.Error("expected the entry to have an ID of 10")
	}
}

func TestFetching(t *testing.T) {
	var games []string
	var err error
	if games, err = steamFetcher.FetchOwnedGamesWithoutGiantBomb(&fakeDatabase); err != nil {
		t.Error(err)
	}

	if games[0] != "mario" {
		t.Error("expected to have fetched the games")
	}
}

func TestFetchingWithGB(t *testing.T) {
	var games []int
	var err error
	if games, err = steamFetcher.FetchOwnedGamesGiantBombID(&fakeDatabase); err != nil {
		t.Error(err)
	}

	if games[0] != 10 {
		t.Error("expected to have fetched the games")
	}
}
