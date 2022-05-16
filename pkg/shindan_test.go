package goshindan

import (
	"fmt"
	"testing"
)

const (
	TestShindanID  = 662881
	TestShindanID2 = 226142
	TestShindanID3 = 509717
)

func TestShindan(t *testing.T) {
	var names = []string{
		"hoge",
		"fuga",
		"spam",
		"eggs",
	}

	for _, name := range names {
		result, err := Shindan(TestShindanID, name)
		if err != nil {
			t.Error(err)
		}
		expected := fmt.Sprintf("u=%s&1", name)
		if result != expected {
			t.Errorf("Result: \"%s\", Expected:\"%s\"", result, expected)
		}
	}
}

func TestGetShindanInfo(t *testing.T) {
	// my shindan
	si, err := GetShindanInfo(TestShindanID2)
	if err != nil {
		t.Error(err)
	}
	expected := ShindanInfo{
		Title:        "帰ってきたチャーハンツクリゲーム",
		Description:  "チャーハンつくるよ！",
		URL:          "https://shindanmaker.com/226142",
		ShindanTimes: 3687,
		Pattern:      12,
		Star:         2,
		Author:       "@kakakaya",
		AuthorPage:   "https://shindanmaker.com/author/kakakaya",
		Keywords:     []string{"メシヨソイ"},
	}
	if si.Title != expected.Title ||
		si.Description != expected.Description ||
		si.URL != expected.URL ||
		si.ShindanTimes < expected.ShindanTimes ||
		si.Pattern != expected.Pattern ||
		si.Star < expected.Star ||
		si.Author != expected.Author ||
		si.AuthorPage != expected.AuthorPage ||
		len(si.Keywords) != len(expected.Keywords) {
		t.Errorf("Result:\n%+v\nExpected:\n%+v\n", si, expected)
	}

	// a famous shindan
	si2, err := GetShindanInfo(TestShindanID3)
	if err != nil {
		t.Error(err)
	}
	expected = ShindanInfo{
		Title:        "ねどこ",
		Description:  "寝るところ",
		URL:          "https://shindanmaker.com/509717",
		ShindanTimes: 57837,
		Pattern:      25232,
		Star:         161,
		Keywords:     []string{"寝る"},
	}
	if si2.Title != expected.Title ||
		si2.Description != expected.Description ||
		si2.URL != expected.URL ||
		si2.ShindanTimes < expected.ShindanTimes ||
		si2.Pattern != expected.Pattern ||
		si2.Star < expected.Star ||
		len(si2.Keywords) != len(expected.Keywords) {
		t.Errorf("Result:\n%+v\nExpected:\n%+v\n", si2, expected)
	}
}
