package database

type Interface interface {
	GetInt(string, int) (map[string]interface{}, error)
	UpsertIntField(string, int, interface{}) error
}
