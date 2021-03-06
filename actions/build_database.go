package actions

import (
	"time"

	"github.com/corybuecker/steamfetcher/database"

	mgo "gopkg.in/mgo.v2"
)

func getSession(databaseHost string) (*mgo.Session, error) {
	var session *mgo.Session
	var err error

	if session, err = mgo.DialWithTimeout(databaseHost, time.Second); err != nil {
		return nil, err
	}
	session.SetSocketTimeout(time.Second)
	session.SetSyncTimeout(time.Second)
	session.SetMode(mgo.Monotonic, true)
	return session, nil
}

func GetDatabase(databaseHost string) (database.Interface, error) {
	var session *mgo.Session
	var err error

	if session, err = getSession(databaseHost); err != nil {
		return nil, err
	}

	database := &database.MongoDB{Collection: session.DB("steamfetcher").C("games")}
	database.SetSession(session)

	return database, nil
}
