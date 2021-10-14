package movies

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type MovieSearchResult struct {
	Page         int
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
	Results      []*Movie
}

type Movie struct {
	ID       int
	Title    string `json:"original_title"`
	Info     string `json:"overview"`
	Poster   string `json:"poster_path"`
	Tagline  string
	Budget   float64
	Language string `json:"original_language"`
}

const (
	baseUrl        = "https://api.themoviedb.org/3/"
	movieSearchUrl = baseUrl + "search/movie"
	apiKeyQs       = "?api_key="
	pageQs         = "&page="
	queryQs        = "&query="
)

var apiKey = apiKeyQs + getKey()

func (m Movie) String() string {
	return m.Title
}

func getKey() string {
	data, err := os.ReadFile(".apikey")
	if err != nil {
		log.Fatal("Could not access API Key")
	}

	return string(data)
}

func getMovies(url string, page int) ([]*Movie, error) {
	movies := []*Movie{}

	pagedUrl := url + pageQs + strconv.Itoa(page)
	resp, err := http.Get(pagedUrl)
	if err != nil {
		log.Fatalf("Failed to query list of movies: %q", err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	var moviesSearch MovieSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&moviesSearch); err != nil {
		log.Fatal(err)
	}

	movies = append(movies, moviesSearch.Results...)

	if moviesSearch.Page < moviesSearch.TotalPages {
		nextPage, err := getMovies(url, moviesSearch.Page+1)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, nextPage...)
	}

	return movies, nil
}

func ListMovies(searchTerms ...string) []*Movie {
	searchUrl := movieSearchUrl + apiKey + queryQs + strings.Join(searchTerms, "%20")
	fmt.Println(searchUrl)

	movies, err := getMovies(searchUrl, 1)
	if err != nil {
		log.Fatal(err)
	}

	return movies
}

var testData = []*Movie{
	{
		ID:       1,
		Title:    "Movie 1",
		Info:     "Long Description for Movie 1",
		Poster:   "RelativeUrlForPosterMovie1.jpg",
		Tagline:  "Smart Tagline for this Movie 1",
		Budget:   1000000,
		Language: "pt-br",
	},
	{
		ID:       2,
		Title:    "Movie 2",
		Info:     "Long Description for Movie 2",
		Poster:   "RelativeUrlForPosterMovie2.jpg",
		Tagline:  "Smart Tagline for this Movie 2",
		Budget:   2000000,
		Language: "en-us",
	},
}
