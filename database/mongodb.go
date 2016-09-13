package database

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoDB struct {
	Collection *mgo.Collection
}

func (mongoDB *MongoDB) Upsert(id string, record map[string]interface{}) error {
	if _, err := mongoDB.Collection.Upsert(map[string]string{"id": id}, bson.M{"$set": record}); err != nil {
		return err
	}
	return nil
}
