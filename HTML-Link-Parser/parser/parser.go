package parser

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type Link struct {
	Href string // link href
	Text string // desc string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	dfs(doc, "")

	return nil, nil
}
func dfs(doc *html.Node, padding string) {
	msg := doc.Data
	if doc.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
