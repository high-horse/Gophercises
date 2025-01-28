package main

import (
	"bytes"
	"fmt"
	"log"
	storage "normalizer/db"
)

func Normalize(phone string) string {
	var buf  bytes.Buffer
	for _, ch := range phone  {
		if ch >= '0' && ch <= '9' {
			buf.WriteRune(ch)
		} 
	}
	return buf.String()
}

func main() {
	storage.PrepareDB()
   initNormalize()
}
func initNormalize() {
    DB, err := storage.InitDB()
    if err != nil {
        panic(err)
    }
    defer DB.Close()
    
    phones, err := DB.GetAllPhones()
    if err != nil {
        panic(err)
    }
    
    fmt.Println("total records", len(phones))
    for _, data := range phones {
        normalized := Normalize(data.PhNumber)
        
        preExistingRecord, err := DB.GetPhoneWithPh(normalized)
        if err != nil {
        	continue
        }

        
        if preExistingRecord != nil && preExistingRecord.Id != data.Id {
            // Delete the current record as the normalized number already exists with a different ID
            err = DB.DeletePhRecord(data.Id)
            if err != nil {
                log.Println("inside pre-existing ", err)
            }
        } else {
            // Update the record with the normalized phone number
            err = DB.UpdatePhRecord(data.Id, normalized)
            if err != nil {
                log.Println("inside else block ", err)
            }
        }
        
        fmt.Printf("original %v, normalized %v \n", data.PhNumber, normalized)
    }
}


func initNormalize_old() {
	DB, err := storage.InitDB()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	
	ph, err :=  DB.GetPhoneWithId(1);
	fmt.Printf("%+v \n", *ph)
	
	pones, err := DB.GetAllPhones()
	if err!= nil {
		panic(err)
	}
	
	println("toal records", len(pones))
	// panic("panid")
	for _, data := range pones {
		normalized := Normalize(data.PhNumber)
		
		preExistingRecord, err := DB.GetPhoneWithPh(normalized)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("original %v, normalized %v \n", data.PhNumber, normalized)
		
		if preExistingRecord != nil && preExistingRecord.Id != data.Id {	
			err = DB.DeletePhRecord(data.Id)
			if err != nil {
				log.Println("inside pre-existing ", err)
			}
		} else {
			if normalized != data.PhNumber {
			    err = DB.UpdatePhRecord(data.Id, normalized)
			    if err != nil {
			        log.Println("inside else block ", err)
			    }
			} else {
			    fmt.Printf("Skipping update for id %v as normalized value is the same\n", data.Id)
			}
		}
		
	}
}