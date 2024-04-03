package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	urlshort "url-short"
)

func main() {
	ymlFile := flag.String("yml", "", "pass yaml file name")
	flag.Parse()

	var parsedYAML []urlshort.PathURL
	var err error

	if *ymlFile != "" {
		data, err := os.ReadFile(*ymlFile)
		if err != nil {
			log.Fatalf("Error reading yml file: %v", err)
		}
		parsedYAML, err = urlshort.ParseYAML(data)
		if err != nil {
			log.Fatalf("Error parsing yml file: %v", err)
		}
	} else {
		// Use the default YAML string
		defaultYAML := []byte(`
            - path: /urlshort
              url: https://github.com/gophercises/urlshort
            - path: /urlshort-final
              url: https://github.com/gophercises/urlshort/tree/solution
        `)
		parsedYAML, err = urlshort.ParseYAML(defaultYAML)
		if err != nil {
			log.Fatalf("Error parsing default YAML: %v", err)
		}
	}

	mux := defaultMux()
	pathsToUrls := urlshort.BuildMap(parsedYAML)
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
