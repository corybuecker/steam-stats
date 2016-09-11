package steam

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/corybuecker/steam-stats-fetcher/test"
)

var steamFetcher Fetcher
var fakeDatabase test.FakeDatabase

var sampleResponse string = "{\"response\": {\"games\": [{\"appid\": 10, \"playtime_forever\": 32}]}}"

type fakejsonfetcher struct {
	response string
}

func (jsonfetcher *fakejsonfetcher) Fetch(url string, destination interface{}) error {
	if err := json.Unmarshal([]byte(jsonfetcher.response), destination); err != nil {
		return err
	}
	return nil
}

func init() {
	fakeDatabase = test.FakeDatabase{}
	steamFetcher = Fetcher{SteamAPIKey: "API KEY", SteamID: "ID"}
	steamFetcher.Jsonfetcher = &fakejsonfetcher{
		response: sampleResponse,
	}
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
	if err := steamFetcher.GetOwnedGames(); err != nil {
		t.Error(err)
	}
	if steamFetcher.OwnedGames.Response.Games[0].ID != 10 {
		t.Error("expected ID of 10")
	}
}

func TestDataUpdating(t *testing.T) {
	if err := steamFetcher.GetOwnedGames(); err != nil {
		t.Error(err)
	}
	if err := steamFetcher.UpdateOwnedGames(&fakeDatabase); err != nil {
		t.Error(err)
	}
	if fakeDatabase.Entry["id"] != 10 {
		t.Error("expected the entry to have an ID of 10")
	}
}
