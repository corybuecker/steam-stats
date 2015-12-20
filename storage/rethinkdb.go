package storage

import "github.com/dancannon/gorethink"

type DatabaseInterface interface {
	UpdateEntry(string, map[string]interface{}) error
	CreateTable(string) error
	CreateDatabase(string) error
	ListDatabases() ([]string, error)
	ListTables(string) ([]string, error)
}

type RethinkDB struct {
	Session *gorethink.Session
}

func (rethinkDB *RethinkDB) UpdateEntry(tableName string, record map[string]interface{}) error {
	return nil
}
func (rethinkDB *RethinkDB) CreateTable(tableName string) error {
	return nil
}
func (rethinkDB *RethinkDB) CreateDatabase(databaseName string) error {
	return nil
}
func (rethinkDB *RethinkDB) ListDatabases() ([]string, error) {
	return make([]string, 5, 5), nil
}
func (rethinkDB *RethinkDB) ListTables(databaseName string) ([]string, error) {
	return make([]string, 5, 5), nil
}
