package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(databaseURL string) (*sql.DB, error) {
	var err error
	db, err = sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping the database: %v", err)
	}

	log.Println("Successfully connected to the database.")
	return db, nil
}

func RunMigrations(databaseUrl string) {
	db, err := InitDB(databaseUrl)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	m, err := migrate.NewWithDatabaseInstance(
		"file:///app/migrations",
		"postgres", driver,
	)
	if err != nil {
		log.Fatalf("Error creating migration instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error running migrations: %v", err)
	}

	log.Println("Migrations applied successfully!")
}
