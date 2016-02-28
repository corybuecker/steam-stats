package storage

import (
	"testing"

	"github.com/corybuecker/steam-stats-fetcher/test"
)

var fakeDatabase test.FakeDatabase

func init() {
	fakeDatabase = test.FakeDatabase{}
}

func TestCreatesDatabase(t *testing.T) {
	fakeDatabase.DatabaseCreated = ""
	fakeDatabase.ExistingDatabases = []string{}

	if err := Setup(&fakeDatabase, "videogames", []string{}); err != nil {
		t.Error(err)
	}
	if fakeDatabase.DatabaseCreated != "videogames" {
		t.Errorf("expected %s, got %s", "videogames", fakeDatabase.DatabaseCreated)
	}
}

func TestDoesNotCreateExistingDatabase(t *testing.T) {
	fakeDatabase.DatabaseCreated = ""
	fakeDatabase.ExistingDatabases = []string{"videogames"}

	if err := Setup(&fakeDatabase, "videogames", []string{}); err != nil {
		t.Error(err)
	}
	if fakeDatabase.DatabaseCreated != "" {
		t.Errorf("expected \"\", got %s", fakeDatabase.DatabaseCreated)
	}
}

func TestCreatesTables(t *testing.T) {
	fakeDatabase.ExistingDatabases = []string{"videogames"}

	if err := Setup(&fakeDatabase, "videogames", []string{"test1"}); err != nil {
		t.Error(err)
	}
	if fakeDatabase.TableCreated != "test1" {
		t.Error("expected that it created the table")
	}
}

func TestDoesNotCreateExistingTables(t *testing.T) {
	fakeDatabase.ExistingDatabases = []string{"videogames"}
	fakeDatabase.ExistingTables = []string{"test1"}

	if err := Setup(&fakeDatabase, "videogames", []string{"test1", "test2"}); err != nil {
		t.Error(err)
	}
	if fakeDatabase.TableCreated != "test2" {
		t.Error("expected that it created the missing table")
	}
}
