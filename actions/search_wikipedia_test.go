package actions

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"testing"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/corybuecker/steamfetcher/database"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var session *mgo.Session
var testPayload []byte
var db *database.MongoDB

func generateURL(search string) string {
	return fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=query&format=json&titles=%s&prop=info|redirects&inprop=url&redirects", url.QueryEscape(search))
}

func TestSearchWikipedia(t *testing.T) {
	testPayload, _ = ioutil.ReadFile("../test_json/mount_blade.json")

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", generateURL("Mount & Blade"), httpmock.NewBytesResponder(200, testPayload))

	session, _ = getSession("localhost")
	session.DB("steamfetcher_test").DropDatabase()
	db = &database.MongoDB{Collection: session.DB("steamfetcher_test").C("games")}
	db.SetSession(session)

	db.Collection.Insert(bson.M{"steam_id": 1, "name": "Mount & Blade"})

	t.Run("updating entries that do not have the wikipedia URL", testSearch)
}

func testSearch(t *testing.T) {
	search(db)
	game, _ := db.GetInt("steam_id", 1)
	assert.Equal(t, "https://en.wikipedia.org/wiki/Mount_%26_Blade", game["wikipediaURL"])
}
