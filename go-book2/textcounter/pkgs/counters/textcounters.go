package counters

import (
	"bufio"
	"bytes"
)

type BytesCounter int

func (c *BytesCounter) Write(p []byte) (int, error) {
	*c += BytesCounter(len(p))

	return len(p), nil
}

func NewBytesCounter() BytesCounter {
	return BytesCounter(0)
}

type WordsCounter int

func (c *WordsCounter) Write(p []byte) (int, error) {
	wordScanner := bufio.NewScanner(bytes.NewReader(p))
	count := 0
	wordScanner.Split(bufio.ScanWords)
	for wordScanner.Scan() {
		count++
	}

	*c = WordsCounter(count)

	return len(p), nil
}

func NewWordsCounter() WordsCounter {
	return WordsCounter(0)
}

type LinesCounter int

func (c *LinesCounter) Write(p []byte) (int, error) {
	lineScanner := bufio.NewScanner(bytes.NewReader(p))
	count := 0
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		count++
	}

	*c = LinesCounter(count)

	return len(p), nil
}

func NewLinesCounter() LinesCounter {
	return LinesCounter(0)
}
