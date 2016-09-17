package database

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoDB struct {
	Collection *mgo.Collection
}

func (mongoDB *MongoDB) UpsertIntField(field string, searchValue int, record interface{}) error {
	if _, err := mongoDB.Collection.Upsert(map[string]int{field: searchValue}, bson.M{"$set": record}); err != nil {
		return err
	}
	return nil
}

func (mongoDB *MongoDB) GetInt(field string, searchValue int) (result map[string]interface{}, err error) {
	if err = mongoDB.Collection.Find(map[string]int{field: searchValue}).One(&result); err != nil {
		return nil, err
	}

	return result, nil
}
