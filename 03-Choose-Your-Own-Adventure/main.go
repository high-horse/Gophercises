package main

import (
	"cyoa/parse"
	"flag"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

var S parse.Story

func main() {
	port := flag.Int("port", 8000, "port to run the app on")
	filename := flag.String("file", "files/adventure.json", "file containing adventure json")
	flag.Parse()
	println("Generating Story...")

	// parse file and read from file
	S = parse.ParseFiles(*filename)
	println("server running on port ", *port)
	http.HandleFunc("/", handler)
	http.HandleFunc("/{arc}", handleArc)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
	

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
