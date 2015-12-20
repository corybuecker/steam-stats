package storage

import "testing"

var fakeDatabase fakeRethinkDB

type fakeRethinkDB struct {
	databaseCreated   string
	existingDatabases []string
}

func (rethinkDB *fakeRethinkDB) UpdateEntry(tableName string, record map[string]interface{}) error {
	return nil
}
func (rethinkDB *fakeRethinkDB) CreateTable(tableName string) error {
	return nil
}
func (rethinkDB *fakeRethinkDB) CreateDatabase(databaseName string) error {
	rethinkDB.databaseCreated = databaseName
	return nil
}
func (rethinkDB *fakeRethinkDB) ListDatabases() ([]string, error) {
	return rethinkDB.existingDatabases, nil
}
func (rethinkDB *fakeRethinkDB) ListTables(databaseName string) ([]string, error) {
	return make([]string, 5, 5), nil
}

func init() {
	fakeDatabase = fakeRethinkDB{}
}

func TestCreatesDatabase(t *testing.T) {
	fakeDatabase.databaseCreated = ""
	fakeDatabase.existingDatabases = []string{}

	if err := Setup(&fakeDatabase, "videogames", []string{}); err != nil {
		t.Error(err)
	}
	if fakeDatabase.databaseCreated != "videogames" {
		t.Error("expected that it created the database")
	}
}

func TestDoesNotCreateExistingDatabase(t *testing.T) {
	fakeDatabase.databaseCreated = ""
	fakeDatabase.existingDatabases = []string{"videogames"}

	if err := Setup(&fakeDatabase, "videogames", []string{}); err != nil {
		t.Error(err)
	}
	if fakeDatabase.databaseCreated != "" {
		t.Error("expected that it did not create the database")
	}
}
