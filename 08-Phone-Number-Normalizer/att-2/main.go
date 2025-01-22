package main

import (
	"bytes"
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
	println("starting main func")
	storage.PrepareDB()
	
	DB, err := storage.InitDB()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
}