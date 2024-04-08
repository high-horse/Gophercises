package main

import (
	"flag"
	"fmt"
	"html-link-parser/parser"
	"os"
	"strings"
)

var exHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>

  <a href="/2-page">A link to 2 page<span> sapan test </span></a>
  </body>
</html>

`

func main() {
	filename := flag.String("file", "", "fle to be parsed")
	flag.Parse()
	
	var text string
	if *filename == "" {
		text = exHtml
	} else {
		var err error
		text, err = readFromFile(*filename)
		if err != nil{
			os.Exit(1)
		}
	}
	r := strings.NewReader(text)
	links, err := parser.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}


func readFromFile(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}