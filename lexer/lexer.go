package lexer

import (
	"fmt"
	"unicode"
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
	ignoreNewline := true
	defer func() {
		l.ignoreNewline = ignoreNewline
	}()

scanAgain:
	l.skipWhitespace()

	pos := l.currentPosition()
	startOffset := l.offset

	ch := l.ch
	switch {
	case l.isLetter(ch):
		content := l.scanIdentifier()
		ty := token.LookupKeyword(content)
		if ty == token.IDENT || ty == token.BREAK || ty == token.CONTINUE || ty == token.RETURN {
			ignoreNewline = false
		}
		return &token.Token{Type: ty, Position: pos, Content: content}
	case '0' <= ch && ch <= '9':
		ignoreNewline = false
		ty, content := l.scanNumber(false)
		return &token.Token{Type: ty, Position: pos, Content: content}
	default:
		l.next()
		switch ch {
		case -1:
			if !l.ignoreNewline {
				return &token.Token{Type: token.NEWLINE, Position: pos, Content: "\n"}
			}

			return &token.Token{Type: token.EOF, Position: pos, Content: ""}
		case '\n':
			// only reach here if ignoreNewline was false and exited from skipWhitespace()
			return &token.Token{Type: token.NEWLINE, Position: pos, Content: "\n"}
		case '/':
			if l.ch == '/' || l.ch == '*' { // comment
				offset := l.offset
				// if any newline is found in the comments, it should be treated as a NEWLINE token and returned first
				if !l.ignoreNewline && l.findNewlineInComments() {
					// reset position to the beginning of comment
					l.ch = '/'
					l.offset = offset - 1
					l.rdOffset = offset
					return &token.Token{Type: token.NEWLINE, Position: pos, Content: "\n"}
				}
				comment := l.scanComment()
				if l.mode&ScanComments == 0 {
					// skip comment and return next token
					l.ignoreNewline = true // if newline needs to be returned, it should be returned before this
					goto scanAgain
				}
				return &token.Token{Type: token.COMMENT, Position: pos, Content: comment}
			}
			ty := l.switch2(token.DIV, token.DIVASSIGN)
			return &token.Token{Type: ty, Position: pos, Content: string(l.src[startOffset:l.offset])}
		case '.':
			if '0' <= l.ch && l.ch <= '9' {
				ignoreNewline = false
				ty, content := l.scanNumber(true)
				return &token.Token{Type: ty, Position: pos, Content: content}
			}
			return &token.Token{Type: token.PERIOD, Position: pos, Content: "."}
		case '"':
			ignoreNewline = false
			content := l.scanString()
			return &token.Token{Type: token.STRING, Position: pos, Content: content}
		case '`':
			ignoreNewline = false
			content := l.scanRawString()
			return &token.Token{Type: token.STRING, Position: pos, Content: content}
		case '+':
			ty := l.switch3(token.ADD, token.ADDASSIGN, '+', token.INC)
			if ty == token.INC {
				ignoreNewline = false
			}
			return &token.Token{Type: ty, Position: pos, Content: string(l.src[startOffset:l.offset])}
		case '-':
			ty := l.switch3(token.SUB, token.SUBASSIGN, '-', token.DEC)
			if ty == token.DEC {
				ignoreNewline = false
			}
			return &token.Token{Type: ty, Position: pos, Content: string(l.src[startOffset:l.offset])}
		case '*':
			ty := l.switch2(token.MUL, token.MULASSIGN)
			return &token.Token{Type: ty, Position: pos, Content: string(l.src[startOffset:l.offset])}
		case '%':
			ty := l.switch2(token.REM, token.REMASSIGN)
			return &token.Token{Type: ty, Position: pos, Content: string(l.src[startOffset:l.offset])}
		case '&':
			ty := l.switch3(token.AND, token.ANDASSIGN, '&', token.LAND)
			return &token.Token{Type: ty, Position: pos, Content: string(l.src[startOffset:l.offset])}
		case '|':
			ty := l.switch3(token.OR, token.ORASSIGN, '|', token.LOR)
			return &token.Token{Type: ty, Position: pos, Content: string(l.src[startOffset:l.offset])}
		case '^':
			ty := l.switch2(token.XOR, token.XORASSIGN)
			return &token.Token{Type: ty, Position: pos, Content: string(l.src[startOffset:l.offset])}
		case '<':
			ty := l.switch4(token.LSS, token.LEQ, '<', token.SHL, token.SHLASSIGN)
			return &token.Token{Type: ty, Position: pos, Content: string(l.src[startOffset:l.offset])}
		case '>':
			ty := l.switch4(token.GTR, token.GEQ, '>', token.SHR, token.SHRASSIGN)
			return &token.Token{Type: ty, Position: pos, Content: string(l.src[startOffset:l.offset])}
		case '=':
			ty := l.switch2(token.ASSIGN, token.EQL)
			return &token.Token{Type: ty, Position: pos, Content: string(l.src[startOffset:l.offset])}
		case '!':
			ty := l.switch2(token.NOT, token.NEQ)
			return &token.Token{Type: ty, Position: pos, Content: string(l.src[startOffset:l.offset])}
		case ':':
			if l.ch == '=' {
				l.next()
				return &token.Token{Type: token.DEFINE, Position: pos, Content: string(l.src[startOffset:l.offset])}
			} else if l.ch == ':' {
				l.next()
				return &token.Token{Type: token.DOUBLECOLON, Position: pos, Content: string(l.src[startOffset:l.offset])}
			} else {
				return &token.Token{Type: token.COLON, Position: pos, Content: ":"}
			}
		case '(':
			return &token.Token{Type: token.LPAREN, Position: pos, Content: string(l.src[startOffset:l.offset])}
		case ')':
			ignoreNewline = false
			return &token.Token{Type: token.RPAREN, Position: pos, Content: string(l.src[startOffset:l.offset])}
		case '[':
			return &token.Token{Type: token.LBRACK, Position: pos, Content: string(l.src[startOffset:l.offset])}
		case ']':
			ignoreNewline = false
			return &token.Token{Type: token.RBRACK, Position: pos, Content: string(l.src[startOffset:l.offset])}
		case '{':
			return &token.Token{Type: token.LBRACE, Position: pos, Content: string(l.src[startOffset:l.offset])}
		case '}':
			ignoreNewline = false
			return &token.Token{Type: token.RBRACE, Position: pos, Content: string(l.src[startOffset:l.offset])}
		case ',':
			return &token.Token{Type: token.COMMA, Position: pos, Content: ","}
		case ';':
			return &token.Token{Type: token.SEMICOLON, Position: pos, Content: ";"}
		}
	}

	// unexpected bom is already reported in next(), don't repeat it here
	if ch != bom {
		l.error(l.line, l.col-1, fmt.Sprintf("illegal character %#U", ch))
	}
	ignoreNewline = l.ignoreNewline // reserver ignoreNewline info
	return &token.Token{Type: token.ILLEGAL, Position: pos, Content: string(ch)}
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
		l.col++
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
			Line:     line,
			Column:   col,
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

func (l *Lexer) isLetter(ch rune) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_' ||
		(ch >= utf8.RuneSelf && unicode.IsLetter(ch))
}

func (l *Lexer) isDigit(ch rune) bool {
	return ('0' <= ch && ch <= '9') || (ch >= utf8.RuneSelf && unicode.IsDigit(ch))
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || (l.ch == '\n' && l.ignoreNewline) || l.ch == '\r' {
		l.next()
	}
}

func (l *Lexer) scanComment() string {
	// initial '/' is already consumed; l.ch == '/' || l.ch == '*'
	offset := l.offset - 1
	line := l.line
	col := l.col - 1
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
	l.error(line, col, "comment not terminated")

exit:
	comment := l.src[offset:l.offset]
	if hasCR {
		comment = l.stripCR(comment)
	}
	return string(comment)
}

func (l *Lexer) scanIdentifier() string {
	offset := l.offset
	for l.isLetter(l.ch) || l.isDigit(l.ch) {
		l.next()
	}
	return string(l.src[offset:l.offset])
}

func (l *Lexer) digitValue(ch rune) int {
	switch {
	case '0' <= ch && ch <= '9':
		return int(ch - '0')
	case 'a' <= ch && ch <= 'z':
		return int(ch - 'a' + 10)
	case 'A' <= ch && ch <= 'Z':
		return int(ch - 'A' + 10)
	}
	return 16 // larger than any legal digit val
}

func (l *Lexer) scanMantissa(base int) {
	for l.digitValue(l.ch) < base {
		l.next()
	}
}

func (l *Lexer) scanNumber(seenDecimalPoint bool) (token.Type, string) {
	offset := l.offset
	line := l.line
	col := l.col
	ty := token.INT

	if seenDecimalPoint {
		offset--
		ty = token.FLOAT
		l.scanMantissa(10)
		goto exponent
	}

	if l.ch == '0' {
		// int or float
		line := l.line
		col := l.col
		l.next()
		if l.ch == 'x' || l.ch == 'X' {
			// hexadecimal int
			l.next()
			l.scanMantissa(16)
			if l.offset-offset <= 2 {
				// only scanned "0x" or "0X"
				l.error(line, col, "illegal hexadecimal number")
			}
		} else if l.ch == 'b' || l.ch == 'B' {
			// binary int
			l.next()
			l.scanMantissa(2)
			if l.offset-offset <= 2 {
				// only scanned "0b" or "0B"
				l.error(line, col, "illegal binary number")
			}
		} else {
			// octal int or float
			seenDecimalDigit := false
			l.scanMantissa(8)
			if l.ch == '8' || l.ch == '9' {
				// illegal octal int or float
				seenDecimalDigit = true
				l.scanMantissa(10)
			}
			if l.ch == '.' || l.ch == 'e' || l.ch == 'E' {
				goto fraction
			}
			// octal int
			if seenDecimalDigit {
				l.error(line, col, "illegal octal number")
			}
		}
		goto exit
	}

	// decimal int or float
	l.scanMantissa(10)

fraction:
	if l.ch == '.' {
		ty = token.FLOAT
		l.next()
		l.scanMantissa(10)
	}

exponent:
	if l.ch == 'e' || l.ch == 'E' {
		ty = token.FLOAT
		l.next()
		if l.ch == '-' || l.ch == '+' {
			l.next()
		}
		if l.digitValue(l.ch) < 10 {
			l.scanMantissa(10)
		} else {
			l.error(line, col, "illegal floating-point exponent")
		}
	}

exit:
	return ty, string(l.src[offset:l.offset])
}

func (l *Lexer) scanEscape(quote rune) bool {
	line := l.line
	col := l.col

	var n, base, max uint32
	switch l.ch {
	case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', quote:
		l.next()
		return true
	case '0', '1', '2', '3', '4', '5', '6', '7':
		n, base, max = 3, 8, 255
	case 'x':
		l.next()
		n, base, max = 2, 16, 255
	case 'u':
		l.next()
		n, base, max = 4, 16, unicode.MaxRune
	case 'U':
		l.next()
		n, base, max = 8, 16, unicode.MaxRune
	default:
		msg := "unknown escape sequence"
		if l.ch < 0 {
			msg = "escape sequence not terminated"
		}
		l.error(line, col, msg)
		return false
	}

	var x uint32
	for n > 0 {
		d := uint32(l.digitValue(l.ch))
		if d >= base {
			msg := fmt.Sprintf("illegal character %#U in escape sequence", l.ch)
			if l.ch < 0 {
				msg = "escape sequence not terminated"
			}
			l.error(line, col, msg)
			return false
		}
		x = x*base + d
		l.next()
		n--
	}

	if x > max || (0xD800 <= x && x < 0xE000) {
		l.error(line, col, "escape sequence is invalid Unicode code point")
		return false
	}

	return true
}

func (l *Lexer) scanString() string {
	// '"' is already consumed
	offset := l.offset - 1
	line := l.line
	col := l.col - 1

	for {
		ch := l.ch
		if ch == '\n' || ch < 0 {
			l.error(line, col, "string literal not terminated")
			break
		}
		l.next()
		if ch == '"' {
			break
		}
		if ch == '\\' {
			l.scanEscape('"')
		}
	}

	return string(l.src[offset:l.offset])
}

func (l *Lexer) scanRawString() string {
	// '`' is already consumed
	offset := l.offset - 1
	line := l.line
	col := l.col - 1

	hasCR := false
	for {
		ch := l.ch
		if ch < 0 {
			l.error(line, col, "raw string literal not terminated")
			break
		}
		l.next()
		if ch == '`' {
			break
		}
		if ch == '\r' {
			hasCR = true
		}
	}

	str := l.src[offset:l.offset]
	if hasCR {
		str = l.stripCR(str)
	}

	return string(str)
}

func (l *Lexer) switch2(ty0 token.Type, ty1 token.Type) token.Type {
	if l.ch == '=' {
		l.next()
		return ty1
	}
	return ty0
}

func (l *Lexer) switch3(ty0 token.Type, ty1 token.Type, ch2 rune, ty2 token.Type) token.Type {
	if l.ch == '=' {
		l.next()
		return ty1
	}
	if l.ch == ch2 {
		l.next()
		return ty2
	}
	return ty0
}

func (l *Lexer) switch4(ty0 token.Type, ty1 token.Type, ch2 rune, ty2 token.Type, ty3 token.Type) token.Type {
	if l.ch == '=' {
		l.next()
		return ty1
	}
	if l.ch == ch2 {
		l.next()
		if l.ch == '=' {
			l.next()
			return ty3
		}
		return ty2
	}
	return ty0
}
