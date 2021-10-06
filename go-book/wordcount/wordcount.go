package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
)

func main() {
	data, err := io.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
	splitData := strings.Split(string(data), " ")

	var wordCount = make(map[string]int)

	for _, word := range splitData {
		lWord := strings.ToLower(word)
		lWord = strings.Trim(lWord, ",.")

		wordCount[lWord]++
	}

	keys := make([]string, 0, len(wordCount))
	for k := range wordCount {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	fmt.Fprintln(writer, "Word\tCount")
	for _, k := range keys {
		v := wordCount[k]
		fmt.Fprintf(writer, "%s\t%d\n", k, v)
	}
	writer.Flush()
}
