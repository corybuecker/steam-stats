package database

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session
var err error
var mongoDB *MongoDB

func init() {
	session, err = mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	session.SetMode(mgo.Monotonic, true)

	mongoDB = &MongoDB{Collection: session.DB("test").C("games")}
}

func TestRunner(t *testing.T) {
	session.DB("test").C("games").DropCollection()
	t.Run("upsert int field with new data", testUpsertIntFieldWithNewData)
	t.Run("upsert int field with existing data", testUpsertIntFieldWithExistingData)
	t.Run("upsert int field with error", testUpsertIntFieldWithError)
	t.Run("get int with error", testGetIntWithError)
	session.DB("test").C("games").DropCollection()
}

func testUpsertIntFieldWithNewData(t *testing.T) {
	mongoDB.UpsertIntField("steam_id", 1, map[string]interface{}{"test": true})
	result, _ := mongoDB.GetInt("steam_id", 1)
	assert.Equal(t, true, result["test"], "should have been equal")
}

func testUpsertIntFieldWithExistingData(t *testing.T) {
	mongoDB.UpsertIntField("steam_id", 1, map[string]interface{}{"test": false})
	result, _ := mongoDB.GetInt("steam_id", 1)
	assert.Equal(t, false, result["test"], "should have been equal")
}

func testUpsertIntFieldWithError(t *testing.T) {
	err := mongoDB.UpsertIntField("", 1, map[string]interface{}{"test": false})
	assert.EqualError(t, err, "Cannot call setValue on the root object")
}

func testGetIntWithError(t *testing.T) {
	_, err := mongoDB.GetInt("", 1)
	assert.EqualError(t, err, "not found")
}
