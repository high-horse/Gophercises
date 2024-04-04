package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
)

var S Story

func main() {
	filename := flag.String("file", "files/adventure.json", "file containing adventure json")
	flag.Parse()
	println("Generating Story...")

	// parse file and read from file
	S = ParseFiles(*filename)

	http.HandleFunc("/", handler)
	http.HandleFunc("/{arc}", handleArc)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func check(err error, msg string) {
	if err != nil {
		log.Fatalf("ERROR: %s \n%v", msg, err)
		os.Exit(1)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	msg := S["intro"]
	tmpl.Execute(w, msg)
}

func handleArc(w http.ResponseWriter, r *http.Request) {
	// parse the {string}
	arc := path.Base(r.URL.Path)
	println("Arc => ", arc)
	msg, ok := S[arc]
	if !ok {
		tmpl := template.Must(template.ParseFiles("templates/404page.html"))
		tmpl.Execute(w, nil)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, msg)
}
