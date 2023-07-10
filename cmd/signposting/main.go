package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	 "github.com/gleanerio/signposting/internal/splinks"
)

func main() {
	urls := testURLS()
	
	fs := flag.NewFlagSet("", flag.ExitOnError) 
	var positionToKeep = fs.Int("position", 1000, "Position of the element to keep in the array")
	var showTests = fs.Bool("testset", false, "Show the set of test URLs and exit (overrides other flags)")
	var runAll = fs.Bool("all", false, "Run all known test URLs")
	var testURL = fs.String("url", "", "URL to test")
	
	err := fs.Parse(os.Args[1:])  // Parse the flags
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	if fs.NFlag() == 0 {
		fmt.Println("no flags provided")
		fs.Usage()
		os.Exit(2)
	} 

	// show the test URLs exit
	if *showTests {
		fmt.Println("Show the test URLs and exist")
		for u := range urls {
			fmt.Printf("%d -> %s \n", u, urls[u])
		}
		fmt.Println("\nCaution, list index number does not align to the number in the URL")
		os.Exit(0)
	}
	
	if !*runAll {
		if *testURL != "" {
			// try the URL
			urls = []string{*testURL}
		} else {
			// try the URL at the index in the slice of test URLs
			if *positionToKeep < len(urls) {
				urls = []string{urls[*positionToKeep]}
			} else {
				fmt.Println("No testing URL index supplied or out of range")
				os.Exit(0)
			}	
		}
	}
	

	
	relValues := []string{"author", "collection", "describedby", "item", "cite-as", "type", "license", "linkset"}

	for _, url := range urls {
		fmt.Println("\n-----------------------------------------------------")
		fmt.Printf("HTTP HEAD Results for %s\n\n", url)
		links, err := splinks.ExtractLinks(url, relValues)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
		
		for l := range links {
			_, err = fmt.Fprintf(writer, "%s \t to %s \t with type %s\n", links[l].Rel, links[l].Url, links[l].Type)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		}
		err = writer.Flush()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		
		fmt.Printf("\nHTML HEAD Results for %s\n\n", url)
		htmllinks, err := splinks.ExtractLinkFromHTMLHead(url)
		if err != nil {
			log.Println(err)
			return
		}
	
		for l := range htmllinks {
			_, err = fmt.Fprintf(writer, "%s \t to %s \t with type %s\n", htmllinks[l].Rel, htmllinks[l].Url, htmllinks[l].Type)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		}
		err = writer.Flush()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		
		
	}
}

func testURLS() []string {
	urls := []string{
		"https://w3id.org/a2a-fair-metrics/01-http-describedby-only/",
		"https://w3id.org/a2a-fair-metrics/02-html-full/",
		"https://w3id.org/a2a-fair-metrics/03-http-citeas-only/",
		"https://w3id.org/a2a-fair-metrics/04-http-describedby-iri/",
		"https://w3id.org/a2a-fair-metrics/05-http-describedby-citeas/",
		"https://w3id.org/a2a-fair-metrics/06-http-citeas-describedby-item/",
		"https://w3id.org/a2a-fair-metrics/07-http-describedby-citeas-linkset-json/",
		"https://w3id.org/a2a-fair-metrics/08-http-describedby-citeas-linkset-txt/",
		"https://w3id.org/a2a-fair-metrics/09-http-describedby-citeas-linkset-json-txt/",
		"https://w3id.org/a2a-fair-metrics/10-http-citeas-not-perma/",
		"https://w3id.org/a2a-fair-metrics/11-http-describedby-iri-wrong-type/",
		"https://w3id.org/a2a-fair-metrics/12-http-item-does-not-resolve/",
		"https://w3id.org/a2a-fair-metrics/13-http-describedby-with-type/",
		"https://w3id.org/a2a-fair-metrics/14-http-describedby-citeas-linkset-json-txt-conneg/",
		"https://w3id.org/a2a-fair-metrics/15-http-describedby-no-conneg/",
		"https://w3id.org/a2a-fair-metrics/16-http-describedby-conneg/",
		"https://w3id.org/a2a-fair-metrics/17-http-citeas-multiple-rels/",
		"https://w3id.org/a2a-fair-metrics/18-html-citeas-only/",
		"https://w3id.org/a2a-fair-metrics/19-html-citeas-multiple-rels/",
		"https://w3id.org/a2a-fair-metrics/20-http-html-citeas-same/",
		"https://w3id.org/a2a-fair-metrics/21-http-html-citeas-differ/",
		"https://w3id.org/a2a-fair-metrics/22-http-html-citeas-describedby-mixed/",
		"https://w3id.org/a2a-fair-metrics/23-http-citeas-describedby-item-license-type-author/",
		"https://w3id.org/a2a-fair-metrics/24-http-citeas-204-no-content/",
		"https://w3id.org/a2a-fair-metrics/26-http-citeas-203-non-authorative/",
		"https://w3id.org/a2a-fair-metrics/27-http-linkset-json-only/",
		"https://w3id.org/a2a-fair-metrics/28-http-linkset-txt-only/",
		"https://w3id.org/a2a-fair-metrics/30-http-citeas-describedby-item-license-type-author-joint/",
		"https://w3id.org/a2a-fair-metrics/31-http-describedby-profile/",
		"https://w3id.org/a2a-fair-metrics/32-http-describedby-profile-conneg/",
		"https://w3id.org/a2a-fair-metrics/33-http-item-profile/",
		"https://w3id.org/a2a-fair-metrics/34-http-item-rocrate/",
		}
		
		return urls
}
