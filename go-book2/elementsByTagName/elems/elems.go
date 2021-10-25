package elems

import (
	"errors"

	"golang.org/x/net/html"
)

func ElementsTagByName(doc *html.Node, name ...string) ([]*html.Node, error) {
	if len(name) == 0 {
		return nil, errors.New("inform at least one element name")
	}

	result := []*html.Node{}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		parseElement(c, name, &result)
	}

	return result, nil
}

func parseElement(elem *html.Node, names []string, result *[]*html.Node) {
	for c := elem.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			parseElement(c, names, result)
		}
	}

	if isInSlice(elem.Data, names) {
		(*result) = append((*result), elem)
	}
}

func isInSlice(v string, s []string) bool {
	for _, item := range s {
		if item == v {
			return true
		}
	}
	return false
}
