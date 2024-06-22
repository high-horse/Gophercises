package parse

import (
	"encoding/json"
	"cyoa/common"
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
	common.Check(err, "error opening file")
	defer file.Close()

	d := json.NewDecoder(file)
	var story Story
	err = d.Decode(&story)
	common.Check(err, "parsing the story")

	// fmt.Println("%+v\n", story)

	return story
}

