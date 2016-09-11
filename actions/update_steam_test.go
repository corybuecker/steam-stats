package actions

import (
	"encoding/json"
	"testing"

	"github.com/corybuecker/steam-stats-fetcher/steam"
	"github.com/corybuecker/steam-stats-fetcher/test"
)

var steamFetcher steam.Fetcher
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
	steamFetcher = steam.Fetcher{SteamAPIKey: "API KEY", SteamID: "ID"}
	steamFetcher.Jsonfetcher = &fakejsonfetcher{
		response: sampleResponse,
	}
}

func TestUpdateSteam(t *testing.T) {
	UpdateSteam(&steamFetcher, &fakeDatabase)

	if fakeDatabase.Entry["id"] != 10 {
		t.Error("expected the entry to have an ID of 10")
	}
}
