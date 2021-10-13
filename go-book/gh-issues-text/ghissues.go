package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"
	"time"
)

const IssuesUrl = "https://api.github.com/search/issues"

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: 	{{.Number}}
User: 		{{.User.Login}}
Title: 		{{.Title | printf "%.64s"}}
Age 		{{.CreatedAt | daysAgo}} days
{{end}}`

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func searchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesUrl + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

var report = template.Must(template.New("report").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	result, err := searchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%d issues:\n", result.TotalCount)

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
