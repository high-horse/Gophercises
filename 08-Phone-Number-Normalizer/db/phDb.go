package db

import (
	"database/sql"
	"fmt"
)


func ResetDB(driver, psqlInfo, dbname string) error {

	db, err := sql.Open(driver, psqlInfo)
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP DATABASE IF EXISTS " + dbname)
	if err != nil {
		panic(err)
	}

	db.Close()
	return CreateDB(driver,psqlInfo, dbname)
}

func CreateDB(driver, psqlInfo, dbname string) error  {

	db, err := sql.Open(driver, psqlInfo)
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE DATABASE "+dbname)
	if err != nil {
		panic(err)
	}

	db.Close()
	return nil
}

func Migrate(driver, psqlInfo, dbname string) error {

	psqlInfo  = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)

	db, err := sql.Open(driver, psqlInfo)
	if err != nil {
		return err
	}
	
	statement := `
	CREATE TABLE IF NOT EXISTS phone_numbers
	(
		id SERIAL PRIMARY KEY,
		value VARCHAR(255)
	)`
	_, err = db.Exec(statement)
	if err != nil {
		return err
	}

	return db.Close()
}




func Seed(driver, psqlInfo, dbname string) error {
	psqlInfo  = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)

	db, err := sql.Open(driver, psqlInfo)
	if err != nil {
		return err
	}

	ph_nos:= []string{	
	"1234567890",
	"123 456 7891",
	"1234567890",
	"123 456 7891",
	"123 456 7891",
	"123 456 7891",
	"(123) 456 7892",
	"123 456 7891",
	"(123) 456-7893",
	"123 456 7891",
	"123-456-7894",
	"123 456 7891",
	"123-456-7890",
	"123 456 7891",
	"1234567892",
	"123 456 7891",
	"(123)456-7892",
	"123 456 7891",
	}

	for _, ph := range ph_nos {
		_, err := insertPhone(db, ph)
		if err != nil {
			return err
		}
	}
	return db.Close()
}


func insertPhone(db *sql.DB, phone string) (int, error) {
	statement := `
		INSERT INTO phone_numbers (value) VALUES($1) RETURNING id
	`
	var id int
	err := db.QueryRow(statement, phone).Scan(&id)
	if err != nil{
		return -1, err
	}
	return id, nil
}


type DB struct {
	db *sql.DB
}

func OpenDB(driver string, psqlInfo string, dbname string) (*DB, error) {
	psqlInfo  = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) GetPhone( id int) (string, error) {
	var phone string
	statement := `SELECT value FROM phone_numbers WHERE id=$1`
	err := db.db.QueryRow(statement, id).Scan(&phone)
	if err != nil {
		return "", err
	}
	return phone, nil
}

func (db *DB)FindPhone( ph string) (*Phone, error) {
	var phone Phone
	statement := `SELECT * FROM phone_numbers WHERE value=$1`
	err := db.db.QueryRow(statement, ph).Scan(&phone.ID, &phone.Value)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &phone, nil
}

func (db *DB)UpdatePhone( id int, ph string) error {
	statement := `UPDATE phone_numbers SET value=$1 WHERE id=$2`
	_, err := db.db.Exec(statement, ph, id)
	return err
}

func (db *DB)DeletePhone(id int) error {
	statement := `DELETE FROM phone_numbers WHERE id=$1`
	_, err := db.db.Exec(statement, id)
	return err
}

func (db *DB) AllPhone() ([]Phone, error) {
	rows, err := db.db.Query("SELECT id, value FROM phone_numbers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var phones []Phone
	for rows.Next() {
		var phone Phone
		err := rows.Scan(&phone.ID, &phone.Value)
		if err != nil {
			return nil, err
		}
		phones = append(phones, phone)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return phones, nil
}

type Phone struct {
	ID int
	Value string
}