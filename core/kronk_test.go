package core

import "testing"
import "strings"

func TestObj(t *testing.T) {
	k := NewKronk([]string{"1:1", "2:2"}, []byte("Hello world!\nHow are you?\nGood?\nGreat to hear!"))
	if len(k.args) != 2 {
		t.Fatalf("2 args passed in, but not set properly")
	}

	if !strings.HasPrefix(string(k.content), "Hello") {
		t.Fatalf("Expected Hello to be the start of content")
	}
}
