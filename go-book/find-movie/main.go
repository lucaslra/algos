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
	cmd := parseArgs()

	foundMovies := movies.ListMovies(cmd.Terms...)

	for i, movie := range foundMovies {
		fmt.Printf("%d\t%s\n", i, movie.Title)
	}
}

func parseArgs() cmd.PosterCmd {
	var cmd cmd.PosterCmd

	set := flag.NewFlagSet(flag.Flag{})
	err := set.ParseStruct(&cmd)
	if err != nil {
		log.Fatal(err)
	}
	if len(cmd.Terms) < 1 {
		fmt.Fprint(os.Stderr, "Invalid usage! Inform the search terms\n")
		set.Help()

		os.Exit(1)
	}
	return cmd
}
