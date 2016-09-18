package actions

import (
	"encoding/json"
	"testing"

	"github.com/corybuecker/steam-stats-fetcher/database"
	"github.com/corybuecker/steam-stats-fetcher/steam"
	"github.com/stretchr/testify/assert"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var steamFetcher steam.Fetcher

var sampleResponse = "{\"response\": {\"games\": [{\"appid\": 15, \"name\": \"game\", \"playtime_forever\": 32}]}}"

type fakejsonfetcher struct {
	response string
}

func (jsonfetcher *fakejsonfetcher) Fetch(url string, destination interface{}) error {
	if err := json.Unmarshal([]byte(jsonfetcher.response), destination); err != nil {
		return err
	}
	return nil
}

var session *mgo.Session
var err error
var mongoDB *database.MongoDB
var result []bson.M

func init() {
	session, _ = mgo.Dial("localhost:27017")
	session.SetMode(mgo.Monotonic, true)

	mongoDB = &database.MongoDB{Collection: session.DB("steam_stats_fetcher_test").C("update_steam_test")}
	mongoDB.Collection.DropCollection()

	steamFetcher = steam.Fetcher{
		Configuration: struct {
			SteamAPIKey string `bson:"steam_api_key"`
			SteamID     string `bson:"steam_id"`
		}{
			SteamAPIKey: "API KEY",
			SteamID:     "ID",
		},
	}
	steamFetcher.Jsonfetcher = &fakejsonfetcher{
		response: sampleResponse,
	}
}

func TestDataUpdating(t *testing.T) {
	UpdateSteam(&steamFetcher, mongoDB)
	result, _ := mongoDB.GetInt("steam_id", 15)
	assert.Equal(t, "game", result["name"], "should have been equal")
}
