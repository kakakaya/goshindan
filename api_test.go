package goshindan

import (
	"fmt"
	"testing"
)

const TestShindanID = 662881

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
