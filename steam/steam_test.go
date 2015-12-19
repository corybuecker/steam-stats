package steam

import (
	"encoding/json"
	"strings"
	"testing"
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

func init() {
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
	ownedGames, _ := steamFetcher.GetOwnedGames(&FakeFetcher{})
	if ownedGames.Response.Games[0].ID != 10 {
		t.Error("expected ID of 10")
	}
}
