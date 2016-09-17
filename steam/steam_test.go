package steam

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/corybuecker/steam-stats-fetcher/database"
	"github.com/stretchr/testify/assert"

	mgo "gopkg.in/mgo.v2"
)

var steamFetcher Fetcher

var sampleResponse = "{\"response\": {\"games\": [{\"appid\": 10, \"name\": \"game\", \"playtime_forever\": 32}]}}"

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

func init() {
	session, err = mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	session.SetMode(mgo.Monotonic, true)

	mongoDB = &database.MongoDB{Collection: session.DB("test").C("games")}

	steamFetcher = Fetcher{
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

func TestURLIncludesAPIKey(t *testing.T) {
	assert.Contains(t, steamFetcher.generateURL(), "API KEY", "should include the steam api key")

}
func TestURLIncludesSteamID(t *testing.T) {
	assert.Contains(t, steamFetcher.generateURL(), "ID", "should include the steam ID")
}

func TestDataMarshalling(t *testing.T) {
	steamFetcher.GetOwnedGames()
	assert.Equal(t, 10, steamFetcher.OwnedGames.Response.Games[0].ID, "should be equal")
}

func TestDataUpdating(t *testing.T) {
	if err := steamFetcher.GetOwnedGames(); err != nil {
		t.Error(err)
	}
	if err := steamFetcher.UpdateOwnedGames(mongoDB); err != nil {
		t.Error(err)
	}

	result, _ := mongoDB.GetInt("steam_id", 10)
	assert.Equal(t, "game", result["name"], "should have been equal")
}
