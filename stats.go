package main

import (
	"log"

	"github.com/corybuecker/steam-stats/database"
	"github.com/dancannon/gorethink"
)

func main() {
	var db *database.DB
	session, err := gorethink.Connect(gorethink.ConnectOpts{Address: "localhost:28015"})
	if err != nil {
		log.Fatalln(err.Error())
	}

	db = &database.DB{Name: "videogames", Tables: []string{"steam", "mygames", "giantbomb"}, Session: session}

	if err := db.EnsureExists(); err != nil {
		log.Fatalln(err.Error())
	}

}
