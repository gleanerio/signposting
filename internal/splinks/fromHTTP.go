package splinks

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"strings"
)

type Links struct {
	Url       string
	Rel       string
	Type      string
	Code      int
	MediaType string
}

func ExtractLinks(url string, relValues []string) ([]Links, error) {
	var links []Links

	resp, err := http.Head(url) // note we only do a HEAD call, so don't go looking for the body in this function

	code := resp.StatusCode
	//	fmt.Println(code)

	contentType := resp.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Println("Media type:", mediaType)

	if err != nil {
		return links, err
	}
	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			// Handle close error
			fmt.Println("Error closing the response body: ", closeErr)
		}
	}()
	//	links := make(map[string]string)
	linkHeaders := resp.Header["Link"]
	for _, linkHeader := range linkHeaders {
		parts := strings.Split(linkHeader, ";")
		var link Links
		// set the code and mediatype set above
		link.Code = code
		link.MediaType = mediaType

		if StringInSlice(parts[1], relValues) {
			for i, part := range parts {
				part = strings.TrimSpace(part)
				// THIS IS A HIDEOUS HACK! parse and convert to a map and work from that.
				// This will likely (or should) fail at least one of the tests
				switch i {
				case 0:
					part = strings.Trim(part, `<`)
					part = strings.Trim(part, `>`)
					//					fmt.Println("Is the URL absolute? ", isAbsoluteURL(part))
					//					fmt.Println("Is the URL relative? ", isRelativeURL(part))
					link.Url = part
				case 1:
					part = strings.ReplaceAll(part, "rel=", "")
					part = strings.Trim(part, `"`)
					link.Rel = part
				case 2:
					part = strings.ReplaceAll(part, "type=", "")
					part = strings.Trim(part, `"`)
					link.Type = part
				}
			}
			links = append(links, link)

		}
	}
	return links, err
}
