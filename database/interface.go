package database

import mgo "gopkg.in/mgo.v2"

type Interface interface {
	SetSession(*mgo.Session)
	GetSession() *mgo.Session
	GetInt(string, int) (map[string]interface{}, error)
	UpsertIntField(string, int, interface{}) error
	GetAllGamesWithoutURL() ([]Game, error)
}

type Game struct {
	ID              int    `json:"appid" bson:"steam_id"`
	Name            string `json:"name" bson:"name"`
	PlaytimeForever int    `json:"playtime_forever" bson:"playtimeForever"`
	PlaytimeRecent  int    `json:"playtime_2weeks" bson:"playtimeRecent"`
	WikipediaURL    string `bson:"wikipediaURL"`
}
