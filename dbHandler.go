package main

import (
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// connect handles connecting to the specified database (dbType) with
// the specifed connection string (dsn).
// It can currently handle PostGreSQL and SQLite databases.
// For SQLite databases, the dsn is the path to the database file.
func connect(dbType string, dsn string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	if dbType == "PGSQL" {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else if dbType == "SQLite" {
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	} else {
		return nil, errors.New("Unsupported database specified in connect function")
	}

	if err != nil {
		return nil, err
	}

	return db, nil

}
