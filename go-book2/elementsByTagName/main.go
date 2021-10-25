package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"azurecoder.dev/go-book/elementsByTagName/elems"
	"golang.org/x/net/html"
)

func main() {
	response, err := http.Get("https://golang.org")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		_ = fmt.Errorf("failed to complete request. %v", response.Status)
		os.Exit(1)
	}

	data, err := html.Parse(response.Body)
	if err != nil {
		panic(err)
	}

	images, err := elems.ElementsTagByName(data, "img")
	if err != nil {
		panic(err)
	}

	for _, img := range images {
		for _, attr := range img.Attr {
			if attr.Key == "src" {
				fmt.Println(attr.Val)
			}
		}
	}

	headings, err := elems.ElementsTagByName(data, "h1", "h2", "h3", "h4")
	if err != nil {
		panic(err)
	}

	for _, header := range headings {
		for c := header.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode {
				headerText := strings.Trim(c.Data, "\n ")
				fmt.Println(headerText)
			}
		}
	}
}
