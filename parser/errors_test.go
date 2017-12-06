package parser

import (
	"testing"

	"github.com/magic003/liza/token"
)

func TestError(t *testing.T) {
	err := Error{
		Pos: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   40,
		},
		Msg: "unexpected comma",
	}

	expected := "test.lz:10:40: unexpected comma"
	if err.Error() != expected {
		t.Errorf("bad string for error: got %s, expected %s", err.Error(), expected)
	}
}
