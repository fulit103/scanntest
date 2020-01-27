package scanner

import (
	"net/http"
	"testing"
)

func TestHtmlTitle(t *testing.T) {
	resp, err := http.Get("https://www.truora.com")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if title, ok, _ := GetHTMLTitle(resp.Body); ok {
		println(title)
	} else {
		println("Fail to get HTML title")
	}
}

func TestHtmlLogo(t *testing.T) {
	resp, err := http.Get("https://www.truora.com")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if title, ok, _ := GetHTMLLogo(resp.Body); ok {
		println(title)
	} else {
		println("Fail to get HTML title")
	}
}
