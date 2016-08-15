package test

type FakeDatabase struct {
	Entry             map[string]interface{}
	DatabaseCreated   string
	ExistingDatabases []string
	TableCreated      string
	ExistingTables    []string
	Rows              map[string]interface{}
}

func (fakeDatabase *FakeDatabase) Upsert(id map[string]interface{}, record map[string]interface{}) error {
	fakeDatabase.Entry = record
	return nil
}

func (fakeDatabase *FakeDatabase) RowsWithoutFields(fieldsToExclude []string) ([]map[string]interface{}, error) {
	return []map[string]interface{}{
		{
			"id":   10,
			"name": "mario",
		},
	}, nil
}

func (fakeDatabase *FakeDatabase) RowsWithField(fieldsToInclude string) ([]map[string]interface{}, error) {
	return []map[string]interface{}{
		{
			"giantbombId": 10,
			"id":          34,
		},
	}, nil
}

func (fakeDatabase *FakeDatabase) GetRow(field string) (map[string]interface{}, error) {
	return fakeDatabase.Rows, nil
}
