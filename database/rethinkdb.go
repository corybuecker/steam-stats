package database

import "github.com/dancannon/gorethink"

type Interface interface {
	Upsert(string, string, map[string]interface{}) error
	CreateTable(string, string) error
	CreateDatabase(string) error
	ListDatabases() ([]string, error)
	ListTables(string) ([]string, error)
	RowsWithoutField(string, string, string) ([]map[string]interface{}, error)
}

type RethinkDB struct {
	Session *gorethink.Session
}

func (rethinkDB *RethinkDB) Upsert(databaseName string, tableName string, record map[string]interface{}) error {
	if _, err := gorethink.DB(databaseName).Table(tableName).Insert(record, gorethink.InsertOpts{Conflict: "update"}).RunWrite(rethinkDB.Session); err != nil {
		return err
	}
	return nil
}
func (rethinkDB *RethinkDB) CreateTable(databaseName string, tableName string) error {
	if _, err := gorethink.DB(databaseName).TableCreate(tableName).Run(rethinkDB.Session); err != nil {
		return err
	}
	return nil
}
func (rethinkDB *RethinkDB) CreateDatabase(databaseName string) error {
	if _, err := gorethink.DBCreate(databaseName).Run(rethinkDB.Session); err != nil {
		return err
	}
	return nil
}
func (rethinkDB *RethinkDB) ListDatabases() ([]string, error) {
	var databases []string
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
	var tables []string
	var err error
	var cursor *gorethink.Cursor

	if cursor, err = gorethink.DB(databaseName).TableList().Run(rethinkDB.Session); err != nil {
		return nil, err
	}

	if err = cursor.All(&tables); err != nil {
		return nil, err
	}
	return tables, nil
}
func (rethinkDB *RethinkDB) RowsWithoutField(databaseName string, tableName string, fieldToExclude string) ([]map[string]interface{}, error) {
	var rows []map[string]interface{}
	var err error
	var cursor *gorethink.Cursor

	filterFunction := func(game gorethink.Term) gorethink.Term {
		return game.HasFields(fieldToExclude).Eq(false)
	}

	if cursor, err = gorethink.DB(databaseName).Table(tableName).Filter(filterFunction).Run(rethinkDB.Session); err != nil {
		return nil, err
	}

	if err = cursor.All(&rows); err != nil {
		return nil, err
	}
	return rows, nil
}
