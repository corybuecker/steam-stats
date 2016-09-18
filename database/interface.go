package database

import mgo "gopkg.in/mgo.v2"

type Interface interface {
	SetSession(*mgo.Session)
	GetSession() *mgo.Session
	GetInt(string, int) (map[string]interface{}, error)
	UpsertIntField(string, int, interface{}) error
}
