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

var fakeDatabase fakeRethinkDB

type fakeRethinkDB struct {
	Entry map[string]interface{}
}

func (rethinkDB *fakeRethinkDB) UpdateEntry(databaseName string, tableName string, record map[string]interface{}) error {
	rethinkDB.Entry = record
	return nil
}
func (rethinkDB *fakeRethinkDB) CreateTable(databaseName string, tableName string) error {
	return nil
}
func (rethinkDB *fakeRethinkDB) CreateDatabase(databaseName string) error {
	return nil
}
func (rethinkDB *fakeRethinkDB) ListDatabases() ([]string, error) {
	return nil, nil
}
func (rethinkDB *fakeRethinkDB) ListTables(databaseName string) ([]string, error) {
	return nil, nil
}

func init() {
	fakeDatabase = fakeRethinkDB{}
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
	if err := steamFetcher.GetOwnedGames(&FakeFetcher{}); err != nil {
		t.Error(err)
	}
	if steamFetcher.OwnedGames.Response.Games[0].ID != 10 {
		t.Error("expected ID of 10")
	}
}

func TestDataUpdating(t *testing.T) {
	if err := steamFetcher.GetOwnedGames(&FakeFetcher{}); err != nil {
		t.Error(err)
	}
	if err := steamFetcher.UpdateOwnedGames(&fakeDatabase); err != nil {
		t.Error(err)
	}
	if fakeDatabase.Entry["id"] != 10 {
		t.Error("expected the entry to have an ID of 10")
	}
}
