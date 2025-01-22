package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
    *sql.DB
}

const (
    drivername = "sqlite3"
    datasourcename = "../phone.db"
)

func InitDB() (DB, error) {
    fmt.Println("Opening database...")
    database, err := sql.Open(drivername, datasourcename)
    if err != nil {
        fmt.Println("Error opening database:", err)
        return DB{}, err
    }

    fmt.Println("Database opened successfully")
    return DB{database}, nil
}

func PrepareDB() {
	database, err := sql.Open(drivername, datasourcename)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	createTables(database)
	seedTables(database)
}

func createTables(db *sql.DB) error {
	stmt := `
		CREATE TABLE IF NOT EXISTS phone_numbers(
			id INTEGER NOT NULL PRIMARY KEY,
			phone VARCHAR(255)
		) `
		
	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	return nil
}

func seedTables(db *sql.DB) error {
	phonebook := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}
	
	for _, num := range phonebook {
		stmt := `
			INSERT INTO phone_numbers(phone) VALUES($1) 
		`
		_, err := db.Exec(stmt, num)
		if err != nil {
			log.Printf("%+v", err)
		}
	}
	return nil
}