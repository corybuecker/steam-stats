package storage

import "github.com/corybuecker/steam-stats-fetcher/database"

// Setup will create any needed non-existant databases and tables.
func Setup(databaseConnection database.Interface, database string, tables []string) error {
	var existingDatabases []string
	var existingTables []string
	var err error

	if existingDatabases, err = databaseConnection.ListDatabases(); err != nil {
		return err
	}

	if !contains(existingDatabases, database) {
		if err = databaseConnection.CreateDatabase(database); err != nil {
			return err
		}
	}

	if existingTables, err = databaseConnection.ListTables(database); err != nil {
		return err
	}

	for _, table := range tables {
		if !contains(existingTables, table) {
			if err = databaseConnection.CreateTable(database, table); err != nil {
				return err
			}
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
