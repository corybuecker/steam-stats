package giantbomb

import (
	"encoding/json"
	"strings"
	"testing"
)

var fakeDatabase fakeRethinkDB

type fakeRethinkDB struct {
	Entry map[string]interface{}
}

func (rethinkDB *fakeRethinkDB) Upsert(databaseName string, tableName string, record map[string]interface{}) error {
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

func (rethinkDB *fakeRethinkDB) RowsWithoutField(databaseName string, tableName string, fieldToExclude string) ([]map[string]interface{}, error) {
	return []map[string]interface{}{
		{
			"name": "mario",
		},
	}, nil
}

type FakeFetcher struct{}

func (fetcher *FakeFetcher) Fetch(url string, data interface{}) error {
	var sampleResponse string = "{\"results\": [{\"id\": 1, \"name\": \"foundgame\"}]}"
	if err := json.Unmarshal([]byte(sampleResponse), data); err != nil {
		return err
	}
	return nil
}

var gbFetcher Fetcher

func init() {
	gbFetcher = Fetcher{GiantBombAPIKey: "API KEY"}
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
