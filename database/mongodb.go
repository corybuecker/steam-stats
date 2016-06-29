package database

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoDB struct {
	Collection *mgo.Collection
}

func (mongoDB *MongoDB) Upsert(id map[string]interface{}, record map[string]interface{}) error {
	if _, err := mongoDB.Collection.Upsert(id, bson.M{"$set": record}); err != nil {
		return err
	}
	return nil
}
func (mongoDB *MongoDB) RowsWithoutFields(fieldsToExclude []string) ([]map[string]interface{}, error) {
	var rows []map[string]interface{}
	var err error
	var outputRaw = make(map[string]interface{})

	for _, field := range fieldsToExclude {
		outputRaw[field] = map[string]bool{"$exists": false}
	}

	if err = mongoDB.Collection.Find(outputRaw).All(&rows); err != nil {
		return nil, err
	}
	return rows, nil
}
func (mongoDB *MongoDB) RowsWithField(fieldToInclude string) ([]map[string]interface{}, error) {
	var rows []map[string]interface{}
	var err error
	if err = mongoDB.Collection.Find(bson.M{fieldToInclude: bson.M{"$exists": true}}).All(&rows); err != nil {
		return nil, err
	}
	return rows, nil
}

func (mongoDB *MongoDB) GetRow(field string) (map[string]interface{}, error) {
	var row map[string]interface{}
	var err error
	if err = mongoDB.Collection.Find(bson.M{"id": field}).One(row); err != nil {
		return nil, err
	}
	return row, nil
}
