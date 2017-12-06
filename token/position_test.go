package token

import (
	"testing"
)

func TestPositionString(t *testing.T) {
	pos := Position{
		Filename: "test.lz",
		Line:     10,
		Column:   40,
	}

	expected := "test.lz:10:40"
	if pos.String() != expected {
		t.Errorf("bad string for position: got %s, expected %s", pos.String(), expected)
	}
}
