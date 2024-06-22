package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	// "io"
	"net/http"
	"net/url"

	// "os"
	link "sitemap-buider/parser"
)

/*
	1. Get the webpage
	2. Parse all the liks
	3. Build proper urls with our link
	4. Filter out links from different domain
	5. Find all pages
	6. Print out XML
*/

const (
	xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"
)
type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	urlFlag := flag.String("url", "http://gophercises.com/", "url tat you want to create sitemap for")
	maxDepth := flag.Int("depth", 50, "max depth of links to follow")
	flag.Parse()
	
	// pages := get(*urlFlag)
	pages := bfs(*urlFlag, *maxDepth)

	toXml := urlset {
		Xmlns: xmlns,
	}
	for _, page := range pages {
		toXml.Urls = append(toXml.Urls, loc{page})
	}

	// fmt.Print(xml.Header)
	// enc := xml.NewEncoder(os.Stdout)
	// enc.Indent("", "  ")
	// if err := enc.Encode(toXml); err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(1)
	// }
	// println()

	printXml(toXml)
	createFile("sitemap.xml", toXml)
}

func printXml(toXml urlset) error {

	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(toXml); err != nil {
		fmt.Fprintln(os.Stderr, "Error encoding XML to console:", err)
	}
	println()

	return nil

}

func createFile(filename string, toXml urlset) error{
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(xml.Header)
	if err != nil {
		return err
	}

	enc := xml.NewEncoder(file)
	enc.Indent("", "  ")
	if err := enc.Encode(toXml); err != nil {
		return err
	}

	return nil
}

func bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStr: struct{}{},
	}

	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		
		for url,_ := range q {
			if _, ok := seen[url]; ok {
				continue
			}

			seen[url] = struct{}{}
			links := get(url)
			// for _, link := range get(url) {
			for _, link := range links{
				nq[link] = struct{}{}
			}
		}
	}

	ret := make([]string, 0, len(seen))
	for url, _ := range seen {
		ret = append(ret, url)
	}
	return ret
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		return []string{}
	}
	defer resp.Body.Close()


	reqUrl := resp.Request.URL

	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	links := hrefs(resp.Body, base)
	
	return filter(links, withPrefix(base))
}


func hrefs(r io.Reader, base string) []string {
	links, _ := link.Parse(r)

	var ret []string
	for _, l := range links {
		switch{
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base + l.Href)

		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)

		default:
			continue
		}
	}

	return ret

}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}

	return ret
}

func withPrefix(pfx string) func (string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}