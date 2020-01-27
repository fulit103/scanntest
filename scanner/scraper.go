package scanner

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func isLogoElement(n *html.Node) bool {
	isLink := n.Type == html.ElementNode && n.Data == "link"
	if isLink {
		//fmt.Println("element: ", n.Data)
		bandera := 0
		for i := range n.Attr {
			a := n.Attr[i]
			// s := fmt.Sprintf("Attr: %s=%s", a.Key, a.Val)
			// fmt.Println(s)
			if a.Key == "rel" && (a.Val == "shortcut icon" || a.Val == "icon") {
				bandera++
			}

			if a.Key == "type" && strings.Contains(a.Val, "image/") {
				bandera++
			}
		}
		// fmt.Println("----")
		if bandera == 2 {
			return true
		}
		return false
	}

	return isLink
}

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func getHref(n *html.Node) string {
	for i := range n.Attr {
		a := n.Attr[i]
		if a.Key == "href" {
			return a.Val
		}
	}
	return ""
}

func traverse(n *html.Node, mode string) (string, bool, error) {
	if mode == "logo" && isLogoElement(n) {
		return getHref(n), true, nil
	}
	if mode == "title" && isTitleElement(n) {
		return n.FirstChild.Data, true, nil
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok, err := traverse(c, mode)
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

	return traverse(doc, "title")
}

// GetHTMLLogo get logo fro page
func GetHTMLLogo(r io.Reader) (string, bool, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return "", false, err
	}

	return traverse(doc, "logo")
}

// GetHTMLTitleFromURL get title from url
func GetHTMLTitleFromURL(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	title, ok, err := GetHTMLTitle(resp.Body)
	if ok {
		return title, nil
	}

	return "", err
}

// GetHTMLLogoFromURL get title from url
func GetHTMLLogoFromURL(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	title, ok, err := GetHTMLLogo(resp.Body)
	if ok {
		return title, nil
	}

	return "", err
}
