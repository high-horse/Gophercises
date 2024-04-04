package parse

import (
	"encoding/json"
	"log"
	"os"
)

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Story map[string]Chapter

func ParseFiles(filename string) Story {

	file, err := os.Open(filename)
	check(err, "error opening file")
	defer file.Close()

	d := json.NewDecoder(file)
	var story Story
	err = d.Decode(&story)
	check(err, "parsing the story")

	// fmt.Println("%+v\n", story)

	return story
}

func check(err error, msg string) {
	if err != nil {
		log.Fatalf("ERROR: %s \n%v", msg, err)
		os.Exit(1)
	}
}
