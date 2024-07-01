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

	db, err := sql.Open("postgres", psqlInfo)
	must(err)
	must(resetDB(db, dbname))
	defer db.Close()

	psqlInfo  = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	must(db.Ping())
	must(createTable(db, "phonebook"))
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

// func insertPhone(db *sql.DB, phone string) (int, error) {
// 	statement := `
// 		INSERT INTO phone_numbers (value) VALUES(`+ phone +`)
// 	`
// 	some, err := db.Exec(statement)
// 	if err != nil{
// 		return -1, err
// 	}
// }


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

