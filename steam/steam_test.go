package steam

import (
	"testing"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/corybuecker/steam-stats-fetcher/database"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var fetcher = Fetcher{}
var mongoDB database.Interface
var err error

func setupDatabase() {
	session, _ := mgo.DialWithTimeout("localhost", time.Millisecond*500)
	session.DB("steam_stats_fetcher").C("configuration").Insert(bson.M{"id": "steam", "steam_api_key": "API KEY", "steam_id": "ID"})
	mongoDB = &database.MongoDB{Collection: session.DB("steam_stats_fetcher").C("steam_test")}
	mongoDB.SetSession(session)
}

func TestRunner(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	setupDatabase()

	fetcher.configure(mongoDB)

	t.Run("url includes api key", testURLIncludesAPIKey)
	t.Run("url includes steam id", urlIncludesSteamID)
	t.Run("data marshalling", dataMarshalling)
	t.Run("data updating", dataUpdating)
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
	assert.Equal(t, 10, fetcher.OwnedGames.Response.Games[0].ID, "should be equal")
}

func dataUpdating(t *testing.T) {
	fetcher.UpdateOwnedGames(mongoDB)
	result, _ := mongoDB.GetInt("steam_id", 10)
	assert.Equal(t, "game", result["name"], "should have been equal")
}
