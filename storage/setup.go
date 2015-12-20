package storage

func Setup(databaseConnection DatabaseInterface, database string, tables []string) error {
	var existingDatabases []string
	var err error

	if existingDatabases, err = databaseConnection.ListDatabases(); err != nil {
		return err
	}

	if !contains(existingDatabases, database) {
		if err = databaseConnection.CreateDatabase(database); err != nil {
			return err
		}
	}

	return nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
