package lexer

import (
	"bytes"
	"testing"

	"github.com/magic003/liza/token"
)

func TestNew(t *testing.T) {
	src := []byte("source code for testing")
	errHandler := func(pos token.Position, msg string) {}

	lexer := New(src, errHandler)
	if lexer == nil {
		t.Error("new lexer should not be nil")
	}

	if !bytes.Equal(lexer.src, src) {
		t.Errorf("bad src for lexer: got %s, expected %s", string(lexer.src), string(src))
	}
	if lexer.errHandler == nil {
		t.Error("errHandler should not be nil")
	}
	if lexer.ch != ' ' {
		t.Errorf("bad ch for lexer: got %v, expected %v", lexer.ch, ' ')
	}
	if lexer.offset != 0 {
		t.Errorf("bad offset for lexer: got %v, expected %v", lexer.offset, 0)
	}
	if lexer.row != 0 {
		t.Errorf("bad row for lexer: got %v, expected %v", lexer.row, 0)
	}
	if lexer.col != 0 {
		t.Errorf("bad col for lexer: got %v, expected %v", lexer.col, 0)
	}
}
