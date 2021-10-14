package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cosiner/flag"
	"github.com/lucaslra/findmovie/cmd"
	"github.com/lucaslra/findmovie/movies"
)

func main() {
	rootCmd := parseArgs()

	foundMovies := movies.ListMovies(rootCmd.Terms...)

	for i, movie := range foundMovies {
		fmt.Printf("%d\t%s\n", i, movie.Title)
	}
}

func parseArgs() cmd.PosterCmd {
	var posterCmd cmd.PosterCmd

	set := flag.NewFlagSet(flag.Flag{})
	err := set.ParseStruct(&posterCmd)
	if err != nil {
		log.Fatal(err)
	}
	if len(posterCmd.Terms) < 1 {
		_, err = fmt.Fprint(os.Stderr, "Invalid usage! Inform the search terms\n")
		if err != nil {
			panic(err)
		}
		set.Help()

		os.Exit(1)
	}
	return posterCmd
}
