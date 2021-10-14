package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "hcounter: %v\n", err)
		os.Exit(1)
	}
	writer := tabwriter.NewWriter(os.Stdout, 0, 20, 0, ' ', tabwriter.TabIndent)
	defer writer.Flush()
	entries := counter(nil, doc)
	keys := make([]string, 0, len(entries))
	for k := range entries {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, e := range keys {
		writer.Write([]byte(fmt.Sprintf("%s\t = \t%d\t\n", e, entries[e])))
	}
}

func counter(result map[string]int, n *html.Node) map[string]int {
	if result == nil {
		result = map[string]int{}
	}

	if n.Type == html.ElementNode {
		result[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		counter(result, c)
	}

	return result
}
