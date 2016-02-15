package database

import "github.com/dancannon/gorethink"

type Interface interface {
	Upsert(string, string, map[string]interface{}) error
	CreateTable(string, string) error
	CreateDatabase(string) error
	ListDatabases() ([]string, error)
	ListTables(string) ([]string, error)
	RowsWithoutFields(string, string, []string) ([]map[string]interface{}, error)
	RowsWithField(string, string, string) ([]map[string]interface{}, error)
	GetRow(string, string, string) (map[string]interface{}, error)
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
func (rethinkDB *RethinkDB) RowsWithoutFields(databaseName string, tableName string, fieldsToExclude []string) ([]map[string]interface{}, error) {
	var rows []map[string]interface{}
	var err error
	var cursor *gorethink.Cursor

	filterFunction := func(game gorethink.Term) gorethink.Term {
		return game.HasFields(fieldsToExclude).Eq(false)
	}

	if cursor, err = gorethink.DB(databaseName).Table(tableName).Filter(filterFunction).Run(rethinkDB.Session); err != nil {
		return nil, err
	}

	if err = cursor.All(&rows); err != nil {
		return nil, err
	}
	return rows, nil
}
func (rethinkDB *RethinkDB) RowsWithField(databaseName string, tableName string, fieldToInclude string) ([]map[string]interface{}, error) {
	var rows []map[string]interface{}
	var err error
	var cursor *gorethink.Cursor

	filterFunction := func(game gorethink.Term) gorethink.Term {
		return game.HasFields(fieldToInclude).Eq(true)
	}

	if cursor, err = gorethink.DB(databaseName).Table(tableName).Filter(filterFunction).Run(rethinkDB.Session); err != nil {
		return nil, err
	}

	if err = cursor.All(&rows); err != nil {
		return nil, err
	}
	return rows, nil
}
func (rethinkDB *RethinkDB) GetRow(databaseName string, tableName string, field string) (map[string]interface{}, error) {
	var row map[string]interface{}
	var err error
	var cursor *gorethink.Cursor

	if cursor, err = gorethink.DB(databaseName).Table(tableName).Get(field).Run(rethinkDB.Session); err != nil {
		return nil, err
	}

	if err = cursor.One(&row); err != nil {
		return nil, err
	}
	return row, nil
}
