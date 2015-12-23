package storage

import "testing"

var fakeDatabase fakeRethinkDB

type fakeRethinkDB struct {
	databaseCreated   string
	existingDatabases []string
	tableCreated      string
	existingTables    []string
}

func (rethinkDB *fakeRethinkDB) Upsert(databaseName string, tableName string, record map[string]interface{}) error {
	return nil
}
func (rethinkDB *fakeRethinkDB) CreateTable(databaseName string, tableName string) error {
	rethinkDB.tableCreated = tableName
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
	return rethinkDB.existingTables, nil
}

func (rethinkDB *fakeRethinkDB) Filter(function func()) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}

func (rethinkDB *fakeRethinkDB) RowsWithoutField(databaseName string, tableName string, fieldToExclude string) ([]map[string]interface{}, error) {
	return []map[string]interface{}{}, nil
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
		t.Errorf("expected %s, got %s", "videogames", fakeDatabase.databaseCreated)
	}
}

func TestDoesNotCreateExistingDatabase(t *testing.T) {
	fakeDatabase.databaseCreated = ""
	fakeDatabase.existingDatabases = []string{"videogames"}

	if err := Setup(&fakeDatabase, "videogames", []string{}); err != nil {
		t.Error(err)
	}
	if fakeDatabase.databaseCreated != "" {
		t.Errorf("expected \"\", got %s", fakeDatabase.databaseCreated)
	}
}

func TestCreatesTables(t *testing.T) {
	fakeDatabase.existingDatabases = []string{"videogames"}

	if err := Setup(&fakeDatabase, "videogames", []string{"test1"}); err != nil {
		t.Error(err)
	}
	if fakeDatabase.tableCreated != "test1" {
		t.Error("expected that it created the table")
	}
}

func TestDoesNotCreateExistingTables(t *testing.T) {
	fakeDatabase.existingDatabases = []string{"videogames"}
	fakeDatabase.existingTables = []string{"test1"}

	if err := Setup(&fakeDatabase, "videogames", []string{"test1", "test2"}); err != nil {
		t.Error(err)
	}
	if fakeDatabase.tableCreated != "test2" {
		t.Error("expected that it created the missing table")
	}
}
