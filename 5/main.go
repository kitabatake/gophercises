package main

import (
	"fmt"
	"flag"
	"github.com/gophercises/link"
	"net/http"
)

var urlQueue []string
var sameDomainLinks []Link

type Link struct {
	Link link.Link
	NormalizedPath string
}

var domain string
var urls map[string]URL

func main() {
	rootUrl := flag.String("url", "https://www.calhoun.io/", "URL to extracts sitemap")
	flag.Parse()

	// extracts same domain links
	url := URL{original:*rootUrl}
	domain = url.extractDomain()

	urls = map[string]URL{}
	extractURLs(url)
	fmt.Println(urls)
	// output xml

}

func extractURLs(url URL) {
	fmt.Println(url.original)
	urls[url.normalizeURL(domain)] = url

	resp, err := http.Get(url.normalizeURL(domain))
	if err != nil {
		panic(err)
	}

	links, err := link.Parse(resp.Body)
	for _, link := range links {
		u := URL {original:link.Href}
		_, already := urls[u.normalizeURL(domain)]
		if !already && url.isSameDomain(link.Href) {
			extractURLs(u)
		}
	}
}