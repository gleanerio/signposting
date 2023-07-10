package splinks

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func ExtractLinkFromHTMLHead(url string) ([]Links, error) {
	var links []Links

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			// Handle close error
			fmt.Println("Error closing the response body: ", closeErr)
		}
	}()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var link Links

	doc.Find("head link").Each(func(_ int, s *goquery.Selection) {
		h, exist := s.Attr("href")
		r, _ := s.Attr("rel")
		t, _ := s.Attr("type")
		if exist {
			link.Url = h
			link.Rel = r
			link.Type = t
		}

		links = append(links, link)
	})

	return links, nil
}
