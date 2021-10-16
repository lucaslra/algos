package main

import (
	"fmt"
	"log"

	"github.com/lucaslra/algos/links"
)

func main() {
	fmt.Println(crawl("https://golang.org"))
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	return list
}
