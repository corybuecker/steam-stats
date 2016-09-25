package actions

import (
	"testing"

	"github.com/corybuecker/steamfetcher/database"
	"github.com/stretchr/testify/assert"
)

func TestUpdateSteam(t *testing.T) {
	t.Run("opening the session", testGetSession)
	t.Run("building the database", testBuildDatabase)
	t.Run("building the database with error", testBuildDatabaseError)
}

func testGetSession(t *testing.T) {
	session, _ := getSession("localhost")
	err := session.Ping()
	assert.Equal(t, err, nil)
}

func testBuildDatabase(t *testing.T) {
	db, _ := GetDatabase("localhost")
	assert.IsType(t, &database.MongoDB{}, db)
}

func testBuildDatabaseError(t *testing.T) {
	_, err := GetDatabase("s")
	assert.EqualError(t, err, "no reachable servers")
}
