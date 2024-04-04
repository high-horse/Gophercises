package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Message struct {
	Text string
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		msg := Message{Text: "Welcome to our server!"}
		tmpl.Execute(w, msg)
	} else if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		var p Person
		err := decoder.Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Printf("Received: %+v\n", p)
	}
}
