package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssuesUrl = "https://api.github.com/search/issues"

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

func main() {
	result, err := searchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	var monthOld, yearOld, older []*Issue
	for _, issue := range result.Items {
		switch {
		case time.Since(issue.CreatedAt).Hours() < (24 * 30):
			monthOld = append(monthOld, issue)
		case time.Since(issue.CreatedAt).Hours() > (24 * 365):
			older = append(older, issue)
		default:
			yearOld = append(yearOld, issue)
		}
	}

	fmt.Printf("Month old issues: %d\n", len(monthOld))
	for _, issue := range monthOld {
		fmt.Printf("#%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
	}
	fmt.Println()
	fmt.Printf("Year old issues: %d\n", len(yearOld))
	for _, issue := range yearOld {
		fmt.Printf("#%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
	}
	fmt.Println()
	fmt.Printf("Older issues: %d\n", len(older))
	for _, issue := range older {
		fmt.Printf("#%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
	}
	fmt.Println()
}
