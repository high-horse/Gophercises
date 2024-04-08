package main

import (
	"fmt"
	"html-link-parser/parser"
	"strings"
)

var exHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>

  <a href="/2-page">A link to 2 page</a>
  </body>
</html>

`

func main() {
	r := strings.NewReader(exHtml)
	links, err := parser.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Println("%+v\n", links)
}
