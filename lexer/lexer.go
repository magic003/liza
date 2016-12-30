package lexer

import (
	"unicode/utf8"

	"github.com/magic003/liza/token"
)

// ErrorHandler is provided to lexer for error handling. If a syntax error is encountered and a handler is provider,
// the handler is called with a position and message.
type ErrorHandler func(pos token.Position, msg string)

// New returns a new instance of lexer.
func New(filename string, src []byte, errHandler ErrorHandler) *Lexer {
	lexer := &Lexer{
		filename:   filename,
		src:        src,
		errHandler: errHandler,
		offset:     0,
		rdOffset:   0,
		line:       1,
		col:        0,
	}

	// read in the first character
	lexer.next()
	if lexer.ch == bom {
		lexer.next() // ignore BOM at file beginning
	}

	return lexer
}

// Lexer holds the insternal state of a lexer.
type Lexer struct {
	// immutable state
	filename   string
	src        []byte // source code
	errHandler ErrorHandler

	// lexing state
	ch       rune // current character, -1 means end-of-file
	offset   int  // character offset
	rdOffset int  // reading offset (position after current character)
	line     int  // current line, starts from 1
	col      int  // column in current line, starts from 1
}

// NextToken returns the next token from the source.
func (l *Lexer) NextToken() *token.Token {
	return nil
}

const bom = 0xFEFF // byte order mark, only permitted as very first character

// next reads the next unicode char into Lexer.ch. Lexer.ch < 0 means end-of-file.
func (l *Lexer) next() {
	// update line and col if current character is newline
	if l.ch == '\n' {
		l.increaseLineNumber()
	}

	if l.rdOffset == len(l.src) { // reach to eof
		l.ch = -1
		l.offset = len(l.src)
		return
	}

	l.offset = l.rdOffset
	l.col++
	r, w := rune(l.src[l.rdOffset]), 1
	switch {
	case r == 0:
		l.error(l.line, l.col, "illegal character NULL")
	case r >= utf8.RuneSelf: // not ASCII
		r, w = utf8.DecodeRune(l.src[l.rdOffset:])
		if r == utf8.RuneError && w == 1 {
			l.error(l.line, l.col, "illegal UTF-8 encoding")
		} else if r == bom && l.offset > 0 {
			l.error(l.line, l.col, "illegal byte order mark")
		}
	}
	l.ch = r
	l.rdOffset += w
}

func (l *Lexer) error(line int, col int, msg string) {
	if l.errHandler != nil {
		pos := token.Position{
			Filename: l.filename,
			Line:     l.line,
			Column:   l.col,
		}
		l.errHandler(pos, msg)
	}
}

func (l *Lexer) increaseLineNumber() {
	l.line++
	l.col = 0
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
		l.next()
	}
}
