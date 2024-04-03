package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	urlshort "url-short"

	"gopkg.in/yaml.v2"
)

func main() {
	ymlFile := flag.String("yml" , "", "pass yaml file name")
	flag.Parse()

	parsed_yaml, err := ParseYaml(*ymlFile)
	if err != nil {
		log.Fatal("Error parsing yml file: ", err)
		return
	}

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	var yaml string
	if(len(parsed_yaml) < 1) {
		yaml = `
			- path: /urlshort
			url: https://github.com/gophercises/urlshort
			- path: /urlshort-final
			url: https://github.com/gophercises/urlshort/tree/solution
		`
	} else {
		yaml, err = GenerateYMLString(parsed_yaml)
		if err != nil {
			log.Fatal("Error parsing yml file: ", err)
			return
		}
	}
	
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}


type URLmap struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func ParseYaml(filename string) ([]URLmap, error) {
	var mapped_data []URLmap
	// var mapped_data map[string]string


	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err  = yaml.Unmarshal(data, &mapped_data)
	if err != nil {
		return nil, err
	}
	return mapped_data, nil
}

func GenerateYMLString(data []URLmap) (string, error) {
	yamlBytes, err := yaml.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(yamlBytes), nil

}