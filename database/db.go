package database

import "github.com/dancannon/gorethink"

type DB struct {
	Name    string
	Tables  []string
	Session *gorethink.Session
}

func (db *DB) EnsureExists() error {
	if err := ensureDBExists(db.Name, db.Session); err != nil {
		return err
	}
	if err := ensureTablesExist(db.Name, db.Tables, db.Session); err != nil {
		return err
	}
	return nil
}

func ensureDBExists(name string, session *gorethink.Session) error {
	var check bool
	var err error

	if check, err = exists(name, session); err != nil {
		return err
	}
	if check == false {
		if err := create(name, session); err != nil {
			return err
		}
	}
	return nil
}

func ensureTablesExist(name string, tables []string, session *gorethink.Session) error {
	var cursor *gorethink.Cursor
	var err error
	var knownTables []string
	if cursor, err = gorethink.DB(name).TableList().Run(session); err != nil {
		return err
	}
	if err = cursor.All(&knownTables); err != nil {
		return err
	}

	for _, table := range tables {
		if contains(knownTables, table) == false {
			if _, err := gorethink.DB(name).TableCreate(table).Run(session); err != nil {
				return err
			}
		}
	}
	return nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func exists(name string, session *gorethink.Session) (bool, error) {
	var databases []string
	var err error
	var cursor *gorethink.Cursor
	if cursor, err = gorethink.DBList().Run(session); err != nil {
		return false, err
	}
	if err = cursor.All(&databases); err != nil {
		return false, err
	}
	return contains(databases, name), nil
}

func create(name string, session *gorethink.Session) error {
	if _, err := gorethink.DBCreate(name).Run(session); err != nil {
		return err
	}
	return nil
}
