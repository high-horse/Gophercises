package main

import (
	"bytes"
	"fmt"
	"regexp"

	phdb"phone-number-normalizer/db"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "root"
	dbname = "gophercise_8"
	driver = "postgres"
)

func main(){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	
	must(phdb.ResetDB(driver, psqlInfo, dbname))
	must(phdb.Migrate(driver,psqlInfo, dbname))
	must(phdb.Seed(driver, psqlInfo, dbname))

	db, err := phdb.OpenDB(driver, psqlInfo, dbname)
	must(err)

	defer db.Close()

	phones, err := db.AllPhone()
	must(err)

	for _, phone := range phones {
		fmt.Printf("working on... %+v\n", phone)
		normalized := normalize(phone.Value)
		existingPhone, err := db.FindPhone(normalized)
		must(err)

		if existingPhone != nil && existingPhone.ID != phone.ID{
			must(db.DeletePhone( phone.ID))
			fmt.Printf("Deleted duplicate: %+v\n", phone)
		} else {
			must(db.UpdatePhone(phone.ID, normalized))
			fmt.Printf("Updated: %+v\n", phone)
		}
	}
}


func must(err error) {
	if err != nil {	
		panic(err)
	}
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

