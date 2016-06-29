package test

type FakeDatabase struct {
	Entry             map[string]interface{}
	DatabaseCreated   string
	ExistingDatabases []string
	TableCreated      string
	ExistingTables    []string
	Rows              map[string]interface{}
}

func (rethinkDB *FakeDatabase) Upsert(id map[string]interface{}, record map[string]interface{}) error {
	rethinkDB.Entry = record
	return nil
}

func (rethinkDB *FakeDatabase) RowsWithoutFields(fieldsToExclude []string) ([]map[string]interface{}, error) {
	return []map[string]interface{}{
		{
			"id":   10,
			"name": "mario",
		},
	}, nil
}

func (rethinkDB *FakeDatabase) RowsWithField(fieldsToInclude string) ([]map[string]interface{}, error) {
	return []map[string]interface{}{
		{
			"giantbombId": 10,
			"id":          34,
		},
	}, nil
}

func (rethinkDB *FakeDatabase) GetRow(field string) (map[string]interface{}, error) {
	return rethinkDB.Rows, nil
}
