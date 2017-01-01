package lexer

import (
	"unicode/utf8"

	"github.com/magic003/liza/token"
)

// ErrorHandler is provided to lexer for error handling. If a syntax error is encountered and a handler is provider,
// the handler is called with a position and message.
type ErrorHandler func(pos token.Position, msg string)

// Mode defines the lexer scan behavior.
type Mode uint

const (
	// ScanComments returns comments as COMMENT tokens
	ScanComments Mode = 1 << iota
)

// New returns a new instance of lexer.
func New(filename string, src []byte, errHandler ErrorHandler, mode Mode) *Lexer {
	lexer := &Lexer{
		filename:      filename,
		src:           src,
		errHandler:    errHandler,
		mode:          mode,
		offset:        0,
		rdOffset:      0,
		line:          1,
		col:           0,
		ignoreNewline: true,
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
	mode       Mode

	// lexing state
	ch       rune // current character, -1 means end-of-file
	offset   int  // character offset
	rdOffset int  // reading offset (position after current character)
	line     int  // current line, starts from 1
	col      int  // column in current line, starts from 1

	ignoreNewline bool // whether to ignore the next newline
}

// NextToken returns the next token from the source.
func (l *Lexer) NextToken() *token.Token {
	l.skipWhitespace()

	pos := l.currentPosition()

	ch := l.ch
	switch ch {
	case -1:
		if !l.ignoreNewline {
			l.ignoreNewline = true
			return &token.Token{Type: token.NEWLINE, Position: pos, Content: "\n"}
		}

		return &token.Token{Type: token.EOF, Position: pos, Content: ""}
	case '\n':
		l.next()
		// only reach here if ignoreNewline was false and exited from skipWhitespace()
		l.ignoreNewline = true
		return &token.Token{Type: token.NEWLINE, Position: pos, Content: "\n"}
	case '/':
		l.next()
		if l.ch == '/' || l.ch == '*' { // comment
			offset := l.offset
			// if any newline is found in the comments, it should be treated as a NEWLINE token and returned first
			if !l.ignoreNewline && l.findNewlineInComments() {
				// reset position to the beginning of comment
				l.ch = '/'
				l.offset = offset - 1
				l.rdOffset = offset
				l.ignoreNewline = true
				return &token.Token{Type: token.NEWLINE, Position: pos, Content: "\n"}
			}
			comment := l.scanComment()
			if l.mode&ScanComments == 0 {
				// skip comment and return next token
				l.ignoreNewline = true
				return l.NextToken()
			}
			return &token.Token{Type: token.COMMENT, Position: pos, Content: comment}
		}
		// TODO handle other case
	}
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

func (l *Lexer) currentPosition() token.Position {
	return token.Position{
		Filename: l.filename,
		Line:     l.line,
		Column:   l.col,
	}
}

func (l *Lexer) findNewlineInComments() bool {
	// initial '/' is already consumed

	defer func(ch rune, offset int, line int, col int) {
		// reset lexer state
		l.ch = ch
		l.offset = offset
		l.rdOffset = offset
		l.line = line
		l.col = col
	}(l.ch, l.offset, l.line, l.col)

	// read ahead until a newline, EOF, or non-comment token is found
	for l.ch == '/' || l.ch == '*' {
		if l.ch == '/' {
			//-style comment always contains a newline
			return true
		}

		/*-style comment: look for newline */
		l.next()
		for l.ch >= 0 {
			ch := l.ch
			if ch == '\n' {
				return true
			}
			l.next()
			if ch == '*' && l.ch == '/' {
				// if end of /*-style comment is found, continue searching
				l.next()
				break
			}
		}
		l.skipWhitespace()            // l.ignoreNewline is false
		if l.ch < 0 || l.ch == '\n' { // EOF or newline
			return true
		}
		if l.ch != '/' {
			// non-comment token
			return false
		}
		l.next() // consume '/'
	}

	// non-comment token
	return false
}

func (l *Lexer) stripCR(src []byte) []byte {
	res := make([]byte, len(src))
	i := 0
	for _, ch := range src {
		if ch != '\r' {
			res[i] = ch
			i++
		}
	}
	return res[:i]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || (l.ch == '\n' && l.ignoreNewline) || l.ch == '\r' {
		l.next()
	}
}

func (l *Lexer) scanComment() string {
	// initial '/' is already consumed; l.ch == '/' || l.ch == '*'
	offset := l.offset - 1
	hasCR := false

	if l.ch == '/' { //-style comment
		for l.ch != '\n' && l.ch >= 0 {
			if l.ch == '\r' {
				hasCR = true
			}
			l.next()
		}

		goto exit
	}

	/*-style comment */
	l.next()
	for l.ch >= 0 {
		ch := l.ch
		if ch == '\r' {
			hasCR = true
		}
		l.next()
		if ch == '*' && l.ch == '/' {
			l.next()
			goto exit
		}
	}

	// reach here means the comment is not terminated
	l.error(l.line, l.col, "comment not terminated")

exit:
	comment := l.src[offset:l.offset]
	if hasCR {
		comment = l.stripCR(comment)
	}
	return string(comment)
}
