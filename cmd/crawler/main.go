package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	visited := make(map[string]bool)
	startURL := "https://w3id.org/a2a-fair-metrics/"
	crawl(startURL, visited)
}


func getLinks(header http.Header) []string {
	links := header["Link"]
	var urls []string
	for _, link := range links {
		// Link header format: <url>; rel="next", <url>; rel="prev", ...
		parts := strings.Split(link, ",")
		for _, part := range parts {
			bracketIndex := strings.Index(part, ">")
			if bracketIndex != -1 {
				urls = append(urls, strings.Trim(part[1:bracketIndex], " "))
			}
		}
	}
	return urls
}

func crawl(urlStr string, visited map[string]bool) {
	if visited[urlStr] {
		return
	}
	visited[urlStr] = true

	resp, err := http.Get(urlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	urls := getLinks(resp.Header)
	for _, u := range urls {
		parsed, err := url.Parse(u)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if parsed.Host == "" {
			parsed = resp.Request.URL.ResolveReference(parsed)
		}
		crawl(parsed.String(), visited)
	}
}

