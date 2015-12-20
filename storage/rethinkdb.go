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
	if _, err := gorethink.DBCreate(databaseName).Run(rethinkDB.Session); err != nil {
		return err
	}
	return nil
}
func (rethinkDB *RethinkDB) ListDatabases() ([]string, error) {
	databases := make([]string, 0)
	var err error
	var cursor *gorethink.Cursor

	if cursor, err = gorethink.DBList().Run(rethinkDB.Session); err != nil {
		return nil, err
	}

	if err = cursor.All(&databases); err != nil {
		return nil, err
	}
	return databases, nil
}
func (rethinkDB *RethinkDB) ListTables(databaseName string) ([]string, error) {
	return make([]string, 5, 5), nil
}
