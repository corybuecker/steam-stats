package storage

import (
	"log"
	"testing"

	"github.com/corybuecker/steam-stats/giantbomb"
	"github.com/corybuecker/steam-stats/steam"
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
}

func TestUpdateOwnedGames(t *testing.T) {
	beforeEach()
	ownedGames := &steam.OwnedGames{
		Response: steam.Response{
			Games: []steam.OwnedGame{
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

func TestUpdateGiantBomb(t *testing.T) {
	beforeEach()
	results := &giantbomb.Search{
		Results: []giantbomb.SearchResult{
			{ID: 1, Name: "test"},
		},
	}
	rethinkdb.EnsureExists()
	err := rethinkdb.UpdateGiantBomb(results.Results)

	if err != nil {
		t.Error(err)
	}
	var row interface{}
	res, err := r.DB("test_create").Table("giantbomb").Filter(map[string]interface{}{"id": 1}).Run(session)
	err = res.One(&row)
	if err != nil {
		t.Error(err)
	}
}

func beforeEach() {
	r.DBDrop("test_create").Run(session)
	rethinkdb = &RethinkDB{Name: "test_create", Tables: []string{"ownedgames", "giantbomb"}, Session: session}
}

func TestDatabaseCreation(t *testing.T) {
	beforeEach()
	if err := rethinkdb.EnsureExists(); err != nil {
		t.Error(err.Error())
	}

	if check, _ := exists("test_create", session); check == false {
		t.Error("expected database to have been created")
	}
}

func TestTableCreation(t *testing.T) {
	beforeEach()
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
