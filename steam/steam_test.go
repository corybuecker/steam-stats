package steam

import (
	"testing"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/corybuecker/mgoconfig"
	"github.com/corybuecker/steamfetcher/database"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var fetcher = &Fetcher{}
var mongoDB database.Interface
var err error

func setupDatabase() {
	session, _ := mgo.DialWithTimeout("localhost", time.Millisecond*500)
	fetcher.ConfigurationSettings = &mgoconfig.Configuration{Database: "steam_test", Key: "steam", Session: session}
	session.DB("steam_test").DropDatabase()

	session.DB("steam_test").C("configuration").Insert(bson.M{"id": "steam", "steam_api_key": "API KEY", "steam_id": "ID"})
	mongoDB = &database.MongoDB{Collection: session.DB("steam_test").C("games")}
	mongoDB.SetSession(session)
}

func TestSteam(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	setupDatabase()

	fetcher.configure(mongoDB)

	t.Run("url includes api key", testURLIncludesAPIKey)
	t.Run("url includes steam id", urlIncludesSteamID)
	t.Run("data marshalling", dataMarshalling)
	t.Run("data updating", dataUpdating)
	t.Run("configuration not found", configurationNotFound)
}

func testURLIncludesAPIKey(t *testing.T) {
	assert.Contains(t, fetcher.generateURL(), "API KEY")
}
func urlIncludesSteamID(t *testing.T) {
	assert.Contains(t, fetcher.generateURL(), "ID")
}

func dataMarshalling(t *testing.T) {
	httpmock.RegisterResponder("GET", fetcher.generateURL(),
		httpmock.NewStringResponder(200, "{\"response\": {\"games\": [{\"appid\": 10, \"name\": \"game\", \"playtime_forever\": 32}]}}"))

	fetcher.getOwnedGames()
	assert.Equal(t, 10, fetcher.OwnedGames.Response.Games[0].ID)
}

func dataUpdating(t *testing.T) {
	fetcher.UpdateOwnedGames(mongoDB)
	result, _ := mongoDB.GetInt("steam_id", 10)
	assert.Equal(t, "game", result["name"])
}

func configurationNotFound(t *testing.T) {
	session := fetcher.ConfigurationSettings.Session
	fetcher.ConfigurationSettings = &mgoconfig.Configuration{Database: "steam_test", Key: "missing", Session: session}
	err := fetcher.UpdateOwnedGames(mongoDB)
	assert.EqualError(t, err, "the steam configuration could not be fetched")
}
