package scaffold

import (
	"strings"
	"testing"
)

func TestNewPage(t *testing.T) {
	html, err := NewPage("Hello World")
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(html, "<title>Hello World</title>") {
		t.Errorf("got: %q", html)
	}
}
