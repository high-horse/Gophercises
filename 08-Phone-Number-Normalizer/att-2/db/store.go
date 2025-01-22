package db

import (
    "database/sql"
    "fmt"
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