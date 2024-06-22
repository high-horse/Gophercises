package main

import (
	"flag"
	"fmt"
	"io"
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

func main() {
	urlFlag := flag.String("url", "http://gophercises.com/", "url tat you want to create sitemap for")
	flag.Parse()

	
	get(*urlFlag)
	pages := get(*urlFlag)

	for _, page := range pages {
		fmt.Println(page)
	}
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
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