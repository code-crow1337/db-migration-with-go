package connection

import "database/sql"

import (
	_ "github.com/lib/pq"
)

// db Globally holds the database connection after it's been initialized.
var db *sql.DB

// Init sets up setting the connection pool global variable.
func Init(dataSourceName string) error {
	var err error
	// Init the database connection
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}

	return db.Ping()
}

// GetDB returns the database instance for direct use
func GetDB() *sql.DB {
	return db
}
