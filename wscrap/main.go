package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

type Fact struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func main() {
	allFacts := make([]Fact, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	)

	collector.OnHTML(".factsList li", func(h *colly.HTMLElement) {
		factId, err := strconv.Atoi(h.Attr("id"))
		if err != nil {
			log.Println(err)
		}
		description := h.Text

		fact := Fact{
			ID:          factId,
			Description: description,
		}

		allFacts = append(allFacts, fact)
	})

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	collector.Visit("https://www.factretriever.com/rhino-facts")

	writeFacts(allFacts)
}

func writeFacts(facts []Fact) {
	file, err := json.MarshalIndent(facts, "", " ")
	if err != nil {
		log.Println("Unable to create JSON file")
		return
	}

	_ = ioutil.WriteFile("rfacts.json", file, 0644)
}
