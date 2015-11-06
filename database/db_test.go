package database

import (
	"log"
	"testing"

	r "github.com/dancannon/gorethink"
)

var database *DB
var session *r.Session

func init() {
	var err error

	session, err = r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	r.DBDrop("test_create").Run(session)
	database = &DB{Name: "test_create", Tables: []string{"steam"}, Session: session}
}

func TestDatabaseCreation(t *testing.T) {
	if err := database.EnsureExists(); err != nil {
		t.Error(err.Error())
	}

	if check, _ := exists("test_create", session); check == false {
		t.Error("expected database to have been created")
	}
}

func TestTableCreation(t *testing.T) {
	var cursor *r.Cursor
	var err error
	var knownTables []string

	if err := database.EnsureExists(); err != nil {
		t.Error(err.Error())
	}

	if cursor, err = r.DB("test_create").TableList().Run(session); err != nil {
		t.Error(err.Error())
	}

	if err = cursor.All(&knownTables); err != nil {
		t.Error(err.Error())
	}

	if contains(knownTables, "steam") == false {
		t.Error("expected table to have been created")
	}
}
