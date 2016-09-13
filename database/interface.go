package database

type Interface interface {
	Upsert(string, map[string]interface{}) error
}
