package main

import (
	"database/sql"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	_ "modernc.org/sqlite" // SQLite driver

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite" // SQLite driver for migration
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Store struct {
	DB *sql.DB
}

func NewStore() *Store {
	// ConnectDB returns a database connection depending on the environment.

	driver := "sqlite"
	connexion := "/tmp/stoik.db"
	migrationString := "sqlite://" + connexion
	migrationFiles := "file://migrations"
	// Connecting to DB
	log.Println("connecting to the database... ")

	db, err := sql.Open(driver, connexion)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Println("could not connect to the database")
		panic(err)
	}
	log.Println("OK")

	// Migrating
	log.Println("INFO: migrating...")
	m, err := migrate.New(migrationFiles, migrationString)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if !errors.Is(err, migrate.ErrNoChange) {
		if err != nil {
			panic(err)
		}
		log.Println("success!")
	} else {
		log.Println("no change required")
	}

	return &Store{DB: db}
}

type Resources struct {
	Store    Store
	Validate *validator.Validate
}
