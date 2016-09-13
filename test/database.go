package test

type FakeDatabase struct {
	Entry map[string]interface{}
}

func (fakeDatabase *FakeDatabase) Upsert(id string, record map[string]interface{}) error {
	fakeDatabase.Entry = record
	return nil
}
