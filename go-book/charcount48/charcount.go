package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"unicode"
	"unicode/utf8"
)

var counts = make(map[rune]map[rune]int)

func assignTo(cat, v rune) {
	if _, ok := counts[cat]; !ok {
		counts[cat] = make(map[rune]int)
	}

	counts[cat][v]++
}

func main() {
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		switch {
		case unicode.IsLetter(r):
			assignTo('L', r)
		case unicode.IsNumber(r):
			assignTo('N', r)
		case unicode.IsControl(r):
			assignTo('C', r)
		case unicode.IsMark(r):
			assignTo('M', r)
		case unicode.IsPunct(r):
			assignTo('P', r)
		case unicode.IsSpace(r):
			assignTo('Z', r)
		}
		utflen[n]++
	}

	fmt.Println("rune\tcount")
	keys := make([]rune, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}
	for _, catLetter := range keys {
		cat := counts[catLetter]
		keys := make([]int, 0, len(cat))
		for k := range cat {
			keys = append(keys, int(k))
		}
		sort.Ints(keys)

		fmt.Printf("\n%v\tcount\n", string(catLetter))
		for _, k := range keys {
			fmt.Printf("%q\t%d\n", k, cat[rune(k)])
		}
	}

	fmt.Println("\nlen\tcount")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 chracters\n", invalid)
	}
}
