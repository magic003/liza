package lexer

import (
	"github.com/magic003/liza/token"
)

// ErrorHandler is provided to lexer for error handling. If a syntax error is encountered and a handler is provider,
// the handler is called with a position and message.
type ErrorHandler func(pos token.Position, msg string)

// New returns a new instance of lexer.
func New(src []byte, errHandler ErrorHandler) *Lexer {
	lexer := &Lexer{
		src:        src,
		errHandler: errHandler,
		ch:         ' ',
		offset:     0,
		lineOffset: 0,
	}

	return lexer
}

// Lexer holds the insternal state of a lexer.
type Lexer struct {
	// immutable state
	src        []byte // source code
	errHandler ErrorHandler

	// lexing state
	ch         rune // current character
	offset     int  // character offset
	lineOffset int  // current line offset
}
