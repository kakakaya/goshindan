package goshindan

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Shindan executes Shindan for given shindanID and returns result.
func Shindan(shindanID int, userName string) (string, error) {
	shindanURL := fmt.Sprintf("https://shindanmaker.com/%d", shindanID)

	values := url.Values{}
	values.Add("u", userName)

	fmt.Println(shindanURL)
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

// ShindanInfo is a struct to describe shindan.
type Shindan struct {
	Title        string
	Description  string
	URL          string
	ShindanTimes int // Times this shindan was shindaned.
	Pattern      int // Patterns for shindan result.
	Star         int
	Author       string
	AuthorPage   string
	Keywords     []string
}

// GetShindanInfo fetches Shindanmaker's information for given shindanID.
func GetShindanInfo(shindanID int) (ShindanInfo, error) {
	shindanURL := fmt.Sprintf("https://shindanmaker.com/%d", shindanID)
	resp, err := http.Get(shindanURL)
	if err != nil {
		return ShindanInfo{}, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return ShindanInfo{}, err
	}
	title := strings.TrimSpace(doc.Find("title").Text())
	desc := strings.TrimSpace(doc.Find("div.shindantitle_block > div.shindandescription").Text())

	stats := doc.Find("div.shindanstats > ul > li > b")
	shindanTimes, err := strconv.Atoi(strings.Replace(stats.Slice(0, 1).Text(), ",", "", -1))
	if err != nil {
		return ShindanInfo{}, err
	}
	patt, err := strconv.Atoi(strings.Replace(stats.Slice(1, 2).Text(), ",", "", -1))
	if err != nil {
		return ShindanInfo{}, err
	}

	star, err := strconv.Atoi(strings.Replace((doc.Find("a.favlabel").Text()), "★", "", 1))
	if err != nil {
		return ShindanInfo{}, err
	}

	author := doc.Find("span.authorlabel > a").Text()

	ap, _ := doc.Find("span.authorlabel > a").Attr("href")
	authorPage := fmt.Sprintf("https://shindanmaker.com%s", ap)

	keywords := []string{}
	kws := doc.Find("span.shindanlabel > a.themelabel")
	kws.Each(func(i int, kw *goquery.Selection) {
		keywords = append(keywords, strings.TrimSpace(kw.Text()))
	})

	return ShindanInfo{
		Title:        title,
		Description:  desc,
		URL:          shindanURL,
		ShindanTimes: shindanTimes,
		Pattern:      patt,
		Star:         star,
		Author:       author,
		AuthorPage:   authorPage,
		Keywords:     keywords,
	}, nil
}

type User struct {
	ID string
	Name string
	UserName string				// @name
	Following int
	Followers int
	ShindanCount int
}

// ListCreatedShindans ...
func (s *User)ListCreatedShindans() ([]Shindan, error) {
	
}
