package giantbomb

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/corybuecker/steam-stats/ratelimiters"
	"github.com/corybuecker/steam-stats/test"
)

type TestClock struct {
	sleepDuration time.Duration
}

func (clock *TestClock) Sleep(d time.Duration) {
	clock.sleepDuration = d
}
func (clock *TestClock) Now() time.Time {
	return time.Date(2009, time.January, 1, 0, 0, 0, 0, time.UTC)
}

var fakeDatabase test.FakeDatabase

type FakeFetcher struct{}

var sampleResponse = "{\"results\": [{\"id\": 1, \"name\": \"foundgame\"}]}"

func (fetcher *FakeFetcher) Fetch(url string, data interface{}) error {
	if err := json.Unmarshal([]byte(sampleResponse), data); err != nil {
		return err
	}
	return nil
}

var gbFetcher Fetcher

func init() {
	gbFetcher = Fetcher{GiantBombAPIKey: "API KEY", RateLimiter: &ratelimiters.GiantBombRateLimiter{Clock: &TestClock{}}}
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

func TestFetchGameById(t *testing.T) {
	sampleResponse = "{\"results\": [{\"id\": 10, \"name\": \"newgame\"}]}"

	if err := gbFetcher.FindGameByID(&FakeFetcher{}, 10); err != nil {
		t.Error(err)
	}

	if err := gbFetcher.UpdateFoundGames(&fakeDatabase); err != nil {
		t.Error(err)
	}
	if fakeDatabase.Entry["name"] != "newgame" {
		t.Error("expected the entry to have an ID of 10")
	}
}
