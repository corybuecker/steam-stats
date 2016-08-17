package giantbomb

import (
	"encoding/json"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/corybuecker/steam-stats-fetcher/ratelimiters"
	"github.com/corybuecker/steam-stats-fetcher/test"
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

type fakejsonfetcher struct {
	response string
}

func (jsonfetcher *fakejsonfetcher) Fetch(url string, destination interface{}) error {
	if err := json.Unmarshal([]byte(jsonfetcher.response), destination); err != nil {
		return err
	}
	return nil
}

var sampleResponse = "{\"results\": [{\"id\": 1, \"name\": \"foundgame\", \"site_detail_url\": \"foundgame.com\"}]}"

var gbFetcher Fetcher

func init() {
	gbFetcher = Fetcher{GiantBombAPIKey: "API KEY", RateLimiter: &ratelimiters.GiantBombRateLimiter{Clock: &TestClock{}}}
	gbFetcher.Jsonfetcher = &fakejsonfetcher{
		response: sampleResponse,
	}
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

func TestDataUpdating(t *testing.T) {
	if err := gbFetcher.FindOwnedGame("gamename"); err != nil {
		t.Error(err)
	}
	if err := gbFetcher.UpdateFoundGames(1, &fakeDatabase); err != nil {
		t.Error(err)
	}
	log.Printf("%v", fakeDatabase.Entry)
	if fakeDatabase.Entry["url"] != "foundgame.com" {
		t.Error("expected the entry to have an ID of 10")
	}
}

func TestDataUpdatingWithMoreThanOneResponse(t *testing.T) {
	fakeDatabase.Entry = nil
	sampleResponse = "{\"results\": [{\"id\": 10, \"name\": \"newgame\", \"site_detail_url\": \"newgame.com\"}, {\"id\": 11, \"name\": \"newgame 2\", \"site_detail_url\": \"newgame.com\"}]}"

	gbFetcher.Jsonfetcher = &fakejsonfetcher{
		response: sampleResponse,
	}
	if err := gbFetcher.FindOwnedGame("newgame"); err != nil {
		t.Error(err)
	}
	if err := gbFetcher.UpdateFoundGames(10, &fakeDatabase); err != nil {
		t.Error(err)
	}
	assert.Empty(t, fakeDatabase.Entry, "should be empty")
}

func TestFetchGameById(t *testing.T) {
	sampleResponse = "{\"results\": [{\"id\": 10, \"name\": \"newgame\", \"site_detail_url\": \"newgame.com\"}]}"

	gbFetcher.Jsonfetcher = &fakejsonfetcher{
		response: sampleResponse,
	}

	if err := gbFetcher.FindGameByID(10); err != nil {
		t.Error(err)
	}

	if err := gbFetcher.UpdateFoundGames(10, &fakeDatabase); err != nil {
		t.Error(err)
	}
	if fakeDatabase.Entry["url"] != "newgame.com" {
		t.Error("expected the entry to have an ID of 10")
	}
}
