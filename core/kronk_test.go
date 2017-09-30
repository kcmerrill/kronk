package core

import (
	"strings"
	"testing"
)

var sampleText string

func TestNewKronk(t *testing.T) {
	k := NewKronk([]string{"1:1", "2:2"}, []byte("Hello:world!\nHow are you?\nHello:Good?\nGreat to hear!"))
	if len(k.args) != 2 {
		t.Fatalf("2 args passed in, but not set properly")
	}

	if !strings.HasPrefix(string(k.content), "Hello") {
		t.Fatalf("Expected Hello to be the start of content")
	}

	cols := []string{"1", "2"}
	for idx, col := range cols {
		// validate our column orders
		if k.cols[idx] != col {
			t.Fatalf("Order for columns should be intact")
		}
		// validate our regexes
		if k.regexes[col] != col {
			t.Fatalf("The regexes are not being properly set")
		}
	}

	// ok, lets test the kronkin
	/*k = NewKronk([]string{`hello:Hello:(.*?)`}, []byte("Hello:world!\nHow are you?\nHello:Good?\nGreat to hear!"))
	fmt.Println(k.matches)
	if len(k.matches) != 1 {
		fmt.Println(len(k.matches))
		t.Fatalf("Only gave 1 column present, however, 1 was not found. +/-")
	}

	if matches, exists := k.matches["hello"]; exists {
		if matches[0] != "world!" {
			t.Fatalf("The column hello should have had multiple matches")
		}
	} else {
		t.Fatalf("hello, a column provided, should exist")
	}
	*/
}
