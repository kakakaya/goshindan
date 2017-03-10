package goshindan

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Shindan executes Shindan for given shindanID and returns result.
func Shindan(shindanID int, userName string) (string, error) {
	shindanURL := fmt.Sprintf("https://shindanmaker.com/%d", shindanID)

	values := url.Values{}
	values.Add("u", userName)

	resp, err := http.PostForm(shindanURL, values)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}
	result := doc.Find("div.result2 > div").Text()

	return strings.TrimSpace(result), nil
}
