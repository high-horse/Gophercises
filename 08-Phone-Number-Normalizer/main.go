package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "root"
	dbname = "gophercise_8"
)

func main(){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	// db, err := sql.Open("postgres", psqlInfo)
	// must(err)
	// must(resetDB(db, dbname))
	// defer db.Close()

	psqlInfo  = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	must(db.Ping())

	_, err = insertPhone(db, "1234567890")
	must(err)
	_, err = insertPhone(db, "1234567890")
	must(err)
	_, err = insertPhone(db, "123 456 7891")
	must(err)
	_, err = insertPhone(db, "(123) 456 7892")
	must(err)
	_, err = insertPhone(db, "(123) 456-7893")
	must(err)
	_, err = insertPhone(db, "123-456-7894")
	must(err)
	_, err = insertPhone(db, "123-456-7890")
	must(err)
	_, err = insertPhone(db, "1234567892")
	must(err)
	id, err := insertPhone(db, "(123)456-7892")
	must(err)
	ph, err := getPhone(db, id)
	must(err)
	fmt.Println(ph)

	phones, err := allPhone(db)
	must(err)

	for _, phone := range phones {
		fmt.Printf("working on... %+v\n", phone)
		normalized := normalize(phone.Value)
		existingPhone, err := findPhone(db, normalized)
		must(err)

		if existingPhone != nil && existingPhone.ID != phone.ID{
			must(deletePhone(db, phone.ID))
			fmt.Printf("Deleted duplicate: %+v\n", phone)
		} else {
			must(updatePhone(db, phone.ID, normalized))
			fmt.Printf("Updated: %+v\n", phone)
		}
	}
}


func createTable(db *sql.DB, name string) error {
	statement := `
	CREATE TABLE IF NOT EXISTS phone_numbers
	(
		id SERIAL PRIMARY KEY,
		value VARCHAR(255)
	)`
	_, err := db.Exec(statement)
	return err
}

func getPhone(db *sql.DB, id int) (string, error) {
	var phone string
	statement := `SELECT value FROM phone_numbers WHERE id=$1`
	err := db.QueryRow(statement, id).Scan(&phone)
	if err != nil {
		return "", err
	}
	return phone, nil
}

func findPhone(db *sql.DB, ph string) (*Phone, error) {
	var phone Phone
	statement := `SELECT * FROM phone_numbers WHERE value=$1`
	err := db.QueryRow(statement, ph).Scan(&phone.ID, &phone.Value)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &phone, nil
}

func updatePhone(db *sql.DB, id int, ph string) error {
	statement := `UPDATE phone_numbers SET value=$1 WHERE id=$2`
	_, err := db.Exec(statement, ph, id)
	return err
}

func deletePhone(db *sql.DB, id int) error {
	statement := `DELETE FROM phone_numbers WHERE id=$1`
	_, err := db.Exec(statement, id)
	return err
}

type Phone struct {
	ID int
	Value string
}

func allPhone(db *sql.DB) ([]Phone, error) {
	rows, err := db.Query("SELECT id, value FROM phone_numbers")
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


func resetDB(db *sql.DB, name string) error {

	_, err := db.Exec("DROP DATABASE IF EXISTS " + dbname)
	if err != nil {
		panic(err)
	}
	return createDB(db, dbname)
}

func must(err error) {
	if err != nil {	
		panic(err)
	}
}

func createDB(db *sql.DB, name string) error  {
	_, err := db.Exec("CREATE DATABASE "+name)
	if err != nil {
		panic(err)
	}
	return nil
}

func normalize_reglar(phone string) string {
	var buff bytes.Buffer
	for _, ch := range phone {
		if ch >= '0' && ch <= '9' {
			buff.WriteRune(ch)
		}
	}

	return buff.String()
}

func normalize(phone string) string {
	// re := regexp.MustCompile("[^0-9]")

	re := regexp.MustCompile("\\D")

	return re.ReplaceAllLiteralString(phone, "")
}

