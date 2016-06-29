package database

type Interface interface {
	Upsert(map[string]interface{}, map[string]interface{}) error
	RowsWithoutFields([]string) ([]map[string]interface{}, error)
	RowsWithField(string) ([]map[string]interface{}, error)
	GetRow(string) (map[string]interface{}, error)
}
