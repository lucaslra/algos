package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string `json:"title"`
	Year  int    `json:"released"`
	Color bool   `json:"color,omitempty"`
	Cast  []string
}

var movies = []Movie{
	{
		Title: "Casablanca",
		Year:  1942,
		Color: false,
		Cast:  []string{"Humphrey Bogart", "Ingrid Bergman"},
	},
	{
		Title: "Cool Hand Luke",
		Year:  1967,
		Color: true,
		Cast:  []string{"Paul Newman"},
	},
	{
		Title: "Bullitt",
		Year:  1968,
		Color: true,
		Cast:  []string{"Steve McQueen", "Jacqueline Bisset"},
	},
}

func main() {
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	var moreMovies []struct {
		Title string `json:"title"`
	}
	err = json.Unmarshal(data, &moreMovies)
	if err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}

	fmt.Printf("%v\n", moreMovies)
}
