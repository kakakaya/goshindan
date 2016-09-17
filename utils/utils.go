package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// shindan ...
func Shindan(shindanID int, userName string) (result string, err error) {
	shindanURL := fmt.Sprintf("https://shindanmaker.com/%d", shindanID)

	values := url.Values{}
	values.Add("u", userName)

	resp, err := http.PostForm(shindanURL, values)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return parseHtml(resp.Body), nil
}

// parseHtml ...
func parseHtml(r io.Reader) (result string) {
	doc, _ := goquery.NewDocumentFromReader(r)
	result = doc.Find("div.result2 > div").Text()

	return strings.TrimSpace(result)
}
