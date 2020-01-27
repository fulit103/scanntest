package scanner

import (
	"net/http"
	"testing"
)

func TestHtmlToRst(t *testing.T) {
	resp, err := http.Get("https://www.rapigo.co")
	//resp, err := http.Get("https://siongui.github.io/zh/2016/03/14/pillow-useful-items-for-me-notes/")
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
