package scanner

import (
	"errors"
	"io"

	"golang.org/x/net/html"
)

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func traverse(n *html.Node) (string, bool, error) {
	if isTitleElement(n) {
		return n.FirstChild.Data, true, nil
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok, err := traverse(c)
		if ok {
			return result, ok, err
		}
	}

	return "", false, errors.New("Title not found")
}

// GetHTMLTitle get Tittle from page
func GetHTMLTitle(r io.Reader) (string, bool, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return "", false, err
	}

	return traverse(doc)
}
