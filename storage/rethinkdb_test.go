package storage

import (
	"log"
	"testing"

	"github.com/corybuecker/steam-stats/fetcher"
	r "github.com/dancannon/gorethink"
)

var rethinkdb *RethinkDB
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
	rethinkdb = &RethinkDB{Name: "test_create", Tables: []string{"ownedgames"}, Session: session}
}

func TestUpdateOwnedGames(t *testing.T) {
	ownedGames := &fetcher.OwnedGames{
		Response: fetcher.Response{
			Games: []fetcher.OwnedGame{
				{
					ID:              1,
					PlaytimeForever: 1,
					Name:            "testgame",
				},
			},
		},
	}
	rethinkdb.EnsureExists()
	err := rethinkdb.UpdateOwnedGames(ownedGames)

	if err != nil {
		t.Error(err)
	}
	var row interface{}
	res, err := r.DB("test_create").Table("ownedgames").Filter(map[string]interface{}{"id": 1}).Run(session)
	err = res.One(&row)
	if err != nil {
		t.Error(err)
	}
}

func TestDatabaseCreation(t *testing.T) {
	if err := rethinkdb.EnsureExists(); err != nil {
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

	if err := rethinkdb.EnsureExists(); err != nil {
		t.Error(err.Error())
	}

	if cursor, err = r.DB("test_create").TableList().Run(session); err != nil {
		t.Error(err.Error())
	}

	if err = cursor.All(&knownTables); err != nil {
		t.Error(err.Error())
	}

	if contains(knownTables, "ownedgames") == false {
		t.Error("expected table to have been created")
	}
}
