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

	var links []Link
	nodes := linkNodes(doc)
	for _, node := range nodes {
		links = append(links, buildLink(node))
		fmt.Println(node)
	}

	return links, nil
}

func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = "manual test"
	return ret
}



func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}
