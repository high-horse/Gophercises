package main

import (
	"flag"
	"fmt"
	"html-link-parser/parser"
	"io"
	"net/http"
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
	url := flag.String("url", "", "URL to be parsed")
	flag.Parse()

	var text string
	if *url != "" {
		var err error
		text, err = getHtml(*url)
		println(text)
		checkErr(err, "getting html data")
	} else if *filename != "" {
		var err error
		text, err = readFromFile(*filename)
		checkErr(err, "reading file contents")
	} else {
		text = exHtml
	}
	r := strings.NewReader(text)
	links, err := parser.Parse(r)
	checkErr(err, "parse file and generate links")
	fmt.Printf("%+v\n", links)
}

func readFromFile(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func getHtml(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusForbidden {
		return "", fmt.Errorf("403 Forbidden: Access to %s is not allowed", url)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}
	return string(bytes), nil
}

func checkErr(err error, message string) {
	if err != nil {
		fmt.Printf("%s: %v\n", message, err)
		os.Exit(1)
	}
}
