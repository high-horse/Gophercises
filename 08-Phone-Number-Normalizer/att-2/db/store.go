package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
    *sql.DB
}

type Phone struct{
	Id int
	PhNumber string
}

const (
    drivername = "sqlite3"
    datasourcename = "./phone_number.db" 
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
	migrateTables(database)
	seedData(database)
	database.Close()
}

func migrateTables(db *sql.DB) error {
	log.Println("creating tables...")
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

func seedData(db *sql.DB) error {
	log.Println("seedng data ...")
	trucate := `
		DELETE FROM phone_numbers;
	`
	_, err := db.Exec(trucate)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	
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


func (db DB) GetPhoneWithId( id int) (*Phone, error) {
    stmt := `
        SELECT id, phone FROM phone_numbers WHERE id = $1
    `

    phone := Phone{}

    err := db.QueryRow(stmt, id).Scan(&phone.Id, &phone.PhNumber)
    if err != nil {
        return nil, fmt.Errorf("failed to get phone with id %d: %w", id, err)
    }

    return &phone, nil
}

func (db DB) GetPhoneWithPh(ph string) (*Phone, error) {
	stmt := `
		SELECT id, phone from phone_numbers WHERE phone = $1
	`
	phone := Phone{}
	err := db.QueryRow(stmt, ph).Scan(&phone.Id, &phone.PhNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get phone with number %+v", ph)
	}
	
	return &phone, nil
}

func (db DB) GetAllPhones() ([]Phone, error) {
	stmt := `
		SELECT id, phone FROM phone_numbers
	`

	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var phones []Phone
	for  rows.Next() {
		var ph Phone
		if err := rows.Scan(&ph.Id, &ph.PhNumber); err!= nil {
			return  phones, err
		}
		phones = append(phones, ph)
	}
	if err := rows.Err() ; err != nil {
		return phones, err
	}
	
	return phones, nil
}

func (db *DB) DeletePhRecord(id int) error {
	stmt := `
		DELETE FROM phone_numbers where id = $1
	`
	result, err := db.Exec(stmt, id)
	if err != nil {
		return err 
	}
	
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if affectedRows <=0 {
	    return errors.New(fmt.Sprintf("no rows affected, affected rows: %v", affectedRows))
	}
	fmt.Printf("deleted id %v \n", id)
	
	return nil	
}

func (db *DB) UpdatePhRecord(id int, ph string) error {
    stmt := `
        UPDATE phone_numbers SET phone = $2 WHERE id = $1
    `
    result, err := db.Exec(stmt, id, ph)
    if err != nil {
        return err
    }

    affectedRows, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if affectedRows > 0 {
        fmt.Printf("Updated id %v with phone %v\n", id, ph)
    } else {
        fmt.Printf("No update necessary for id %v\n", id)
    }
    
    return nil
}


func (db *DB) UpdatePhRecord1(id int, ph string) error {
	stmt := `
		UPDATE phone_numbers SET phone = $2 WHERE id = $1
	`
	_, err := db.Exec(stmt, id, ph)
	fmt.Println("updated ", id, " ", ph)
	return err
}