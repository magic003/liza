package lexer

import (
	"bytes"
	"testing"
	"unicode/utf8"

	"github.com/magic003/liza/token"
)

func TestNew(t *testing.T) {
	filename := "test_file.liza"
	src := []byte("\uFEFFsource code for testing")
	errHandler := func(pos token.Position, msg string) {}

	lexer := New(filename, src, errHandler, ScanComments)
	if lexer == nil {
		t.Error("new lexer should not be nil")
	}

	if lexer.filename != filename {
		t.Errorf("bad filename for lexer: got %s, expected %s", string(lexer.filename), string(filename))
	}
	if !bytes.Equal(lexer.src, src) {
		t.Errorf("bad src for lexer: got %s, expected %s", string(lexer.src), string(src))
	}
	if lexer.errHandler == nil {
		t.Error("errHandler should not be nil")
	}
	if lexer.mode != ScanComments {
		t.Errorf("bad mode for lexer: got %v, expected %v", lexer.mode, ScanComments)
	}
	if lexer.ch != 's' { // BOM should be ignored
		t.Errorf("bad ch for lexer: got %c, expected %c", lexer.ch, 's')
	}
	if lexer.offset != 3 {
		t.Errorf("bad offset for lexer: got %v, expected %v", lexer.offset, 3)
	}
	if lexer.rdOffset != 4 {
		t.Errorf("bad rdOffset for lexer: got %v, expected %v", lexer.rdOffset, 4)
	}
	if lexer.line != 1 {
		t.Errorf("bad line for lexer: got %v, expected %v", lexer.line, 1)
	}
	if lexer.col != 2 {
		t.Errorf("bad col for lexer: got %v, expected %v", lexer.col, 2)
	}

	if !lexer.ignoreNewline {
		t.Errorf("bad ignoreNewline for lexer: got %v, expected %v", lexer.ignoreNewline, true)
	}
}

func TestNextAscii(t *testing.T) {
	filename := "test_file.liza"
	src := []byte("ab")

	lexer := New(filename, src, nil, 0)

	if lexer.ch != 'a' {
		t.Errorf("bad ch after next(): got %c, expected %c", lexer.ch, 'a')
	}
	if lexer.offset != 0 {
		t.Errorf("bad offset after next(): got %v, expected %v", lexer.offset, 0)
	}
	if lexer.rdOffset != 1 {
		t.Errorf("bad rdOffset after next(): got %v, expected %v", lexer.rdOffset, 1)
	}
	if lexer.line != 1 {
		t.Errorf("bad line after next(): got %v, expected %v", lexer.line, 1)
	}
	if lexer.col != 1 {
		t.Errorf("bad col after next(): got %v, expected %v", lexer.col, 1)
	}

	lexer.next()
	if lexer.ch != 'b' {
		t.Errorf("bad ch after next(): got %c, expected %c", lexer.ch, 'b')
	}
	if lexer.offset != 1 {
		t.Errorf("bad offset after next(): got %v, expected %v", lexer.offset, 1)
	}
	if lexer.rdOffset != 2 {
		t.Errorf("bad rdOffset after next(): got %v, expected %v", lexer.rdOffset, 2)
	}
	if lexer.line != 1 {
		t.Errorf("bad line after next(): got %v, expected %v", lexer.line, 1)
	}
	if lexer.col != 2 {
		t.Errorf("bad col after next(): got %v, expected %v", lexer.col, 2)
	}
}

func TestNextNonAscii(t *testing.T) {
	filename := "test_file.liza"
	src := []byte("测试")

	lexer := New(filename, src, nil, 0)

	if lexer.ch != '测' {
		t.Errorf("bad ch after next(): got %c, expected %c", lexer.ch, '测')
	}
	if lexer.offset != 0 {
		t.Errorf("bad offset after next(): got %v, expected %v", lexer.offset, 0)
	}
	if lexer.rdOffset != len([]byte("测")) {
		t.Errorf("bad rdOffset after next(): got %v, expected %v", lexer.rdOffset, len([]byte("测")))
	}
	if lexer.line != 1 {
		t.Errorf("bad line after next(): got %v, expected %v", lexer.line, 1)
	}
	if lexer.col != 1 {
		t.Errorf("bad col after next(): got %v, expected %v", lexer.col, 1)
	}

	lexer.next()
	if lexer.ch != '试' {
		t.Errorf("bad ch after next(): got %c, expected %c", lexer.ch, '试')
	}
	if lexer.offset != len([]byte("试")) {
		t.Errorf("bad offset after next(): got %v, expected %v", lexer.offset, len([]byte("试")))
	}
	if lexer.rdOffset != len([]byte("测试")) {
		t.Errorf("bad rdOffset after next(): got %v, expected %v", lexer.rdOffset, len([]byte("测试")))
	}
	if lexer.line != 1 {
		t.Errorf("bad line after next(): got %v, expected %v", lexer.line, 1)
	}
	if lexer.col != 2 {
		t.Errorf("bad col after next(): got %v, expected %v", lexer.col, 2)
	}
}

func TestNextEmptyFile(t *testing.T) {
	filename := "empty.liza"
	var src []byte

	lexer := New(filename, src, nil, 0)
	if lexer.ch != -1 {
		t.Errorf("bad ch after next(): got %c, expected %c", lexer.ch, -1)
	}
	if lexer.offset != 0 {
		t.Errorf("bad offset after next(): got %v, expected %v", lexer.offset, 0)
	}
}

func TestNextEof(t *testing.T) {
	filename := "test.liza"
	src := []byte("ab")

	lexer := New(filename, src, nil, 0)
	lexer.next()
	lexer.next()

	if lexer.ch != -1 {
		t.Errorf("bad ch after next(): got %c, expected %c", lexer.ch, -1)
	}
	if lexer.offset != 2 {
		t.Errorf("bad offset after next(): got %v, expected %v", lexer.offset, 2)
	}
}

func TestNextNewLine(t *testing.T) {
	filename := "test.liza"
	src := []byte("a\nb\nc")

	lexer := New(filename, src, nil, 0)
	lexer.next()
	lexer.next()

	if lexer.line != 2 {
		t.Errorf("bad line after next(): got %c, expected %c", lexer.line, 2)
	}
	if lexer.col != 1 {
		t.Errorf("bad col after next(): got %c, expected %c", lexer.col, 1)
	}

	lexer.next()
	lexer.next()

	if lexer.line != 3 {
		t.Errorf("bad line after next(): got %c, expected %c", lexer.line, 3)
	}
	if lexer.col != 1 {
		t.Errorf("bad col after next(): got %c, expected %c", lexer.col, 1)
	}
}

func TestNextIllegalNull(t *testing.T) {
	filename := "test_file.liza"
	src := []byte("\u0000")
	var errPos token.Position
	var errMsg string
	errHandler := func(pos token.Position, msg string) {
		errPos = pos
		errMsg = msg
	}

	lexer := New(filename, src, errHandler, 0)

	if errPos.Filename != filename {
		t.Errorf("bad filename in error handler: got %s, expected %s", errPos.Filename, filename)
	}
	if errPos.Line != 1 {
		t.Errorf("bad line in error handler: got %v, expected %v", errPos.Line, 1)
	}
	if errPos.Column != 1 {
		t.Errorf("bad column in error handler: got %v, expected %v", errPos.Column, 1)
	}
	if errMsg != "illegal character NULL" {
		t.Errorf("bad error msg in error handler: got %s, expected %s", errMsg, "illegal character NULL")
	}

	if lexer.ch != '\u0000' {
		t.Errorf("bad ch after next(): got %c, expected %c", lexer.ch, '\u0000')
	}
	if lexer.offset != 0 {
		t.Errorf("bad offset after next(): got %v, expected %v", lexer.offset, 0)
	}
	if lexer.rdOffset != 1 {
		t.Errorf("bad rdOffset after next(): got %v, expected %v", lexer.rdOffset, 1)
	}
	if lexer.line != 1 {
		t.Errorf("bad line after next(): got %v, expected %v", lexer.line, 1)
	}
	if lexer.col != 1 {
		t.Errorf("bad col after next(): got %v, expected %v", lexer.col, 1)
	}
}

func TestNextIllegalUtf8Encoding(t *testing.T) {
	filename := "test_file.liza"
	src := []byte{0xff, 0xfe, 0xfd}
	var errPos token.Position
	var errMsg string
	errHandler := func(pos token.Position, msg string) {
		errPos = pos
		errMsg = msg
	}

	lexer := New(filename, src, errHandler, 0)

	if errPos.Filename != filename {
		t.Errorf("bad filename in error handler: got %s, expected %s", errPos.Filename, filename)
	}
	if errPos.Line != 1 {
		t.Errorf("bad line in error handler: got %v, expected %v", errPos.Line, 1)
	}
	if errPos.Column != 1 {
		t.Errorf("bad column in error handler: got %v, expected %v", errPos.Column, 1)
	}
	if errMsg != "illegal UTF-8 encoding" {
		t.Errorf("bad error msg in error handler: got %s, expected %s", errMsg, "illegal UTF-8 encoding")
	}

	if lexer.ch != utf8.RuneError {
		t.Errorf("bad ch after next(): got %c, expected %c", lexer.ch, utf8.RuneError)
	}
	if lexer.offset != 0 {
		t.Errorf("bad offset after next(): got %v, expected %v", lexer.offset, 0)
	}
	if lexer.rdOffset != 1 {
		t.Errorf("bad rdOffset after next(): got %v, expected %v", lexer.rdOffset, 1)
	}
	if lexer.line != 1 {
		t.Errorf("bad line after next(): got %v, expected %v", lexer.line, 1)
	}
	if lexer.col != 1 {
		t.Errorf("bad col after next(): got %v, expected %v", lexer.col, 1)
	}
}

func TestNextIllegalBom(t *testing.T) {
	filename := "test_file.liza"
	src := []byte("a\uFEFF")
	var errPos token.Position
	var errMsg string
	errHandler := func(pos token.Position, msg string) {
		errPos = pos
		errMsg = msg
	}

	lexer := New(filename, src, errHandler, 0)
	lexer.next()

	if errPos.Filename != filename {
		t.Errorf("bad filename in error handler: got %s, expected %s", errPos.Filename, filename)
	}
	if errPos.Line != 1 {
		t.Errorf("bad line in error handler: got %v, expected %v", errPos.Line, 1)
	}
	if errPos.Column != 2 {
		t.Errorf("bad column in error handler: got %v, expected %v", errPos.Column, 2)
	}
	if errMsg != "illegal byte order mark" {
		t.Errorf("bad error msg in error handler: got %s, expected %s", errMsg, "illegal byte order mark")
	}

	if lexer.ch != bom {
		t.Errorf("bad ch after next(): got %c, expected %c", lexer.ch, bom)
	}
	if lexer.offset != 1 {
		t.Errorf("bad offset after next(): got %v, expected %v", lexer.offset, 1)
	}
	if lexer.rdOffset != 4 {
		t.Errorf("bad rdOffset after next(): got %v, expected %v", lexer.rdOffset, 4)
	}
	if lexer.line != 1 {
		t.Errorf("bad line after next(): got %v, expected %v", lexer.line, 1)
	}
	if lexer.col != 2 {
		t.Errorf("bad col after next(): got %v, expected %v", lexer.col, 2)
	}
}

func TestIsLetter(t *testing.T) {
	filename := "test_file.liza"
	src := []byte("")

	lexer := New(filename, src, nil, 0)

	if !lexer.isLetter('a') {
		t.Errorf("bad result for isLetter(): got %v, expected %v", lexer.isLetter('a'), true)
	}
	if !lexer.isLetter('z') {
		t.Errorf("bad result for isLetter(): got %v, expected %v", lexer.isLetter('z'), true)
	}
	if !lexer.isLetter('A') {
		t.Errorf("bad result for isLetter(): got %v, expected %v", lexer.isLetter('A'), true)
	}
	if !lexer.isLetter('Z') {
		t.Errorf("bad result for isLetter(): got %v, expected %v", lexer.isLetter('Z'), true)
	}
	if !lexer.isLetter('_') {
		t.Errorf("bad result for isLetter(): got %v, expected %v", lexer.isLetter('_'), true)
	}
	if !lexer.isLetter('ŝ') {
		t.Errorf("bad result for isLetter(): got %v, expected %v", lexer.isLetter('ŝ'), true)
	}

	if lexer.isLetter('1') {
		t.Errorf("bad result for isLetter(): got %v, expected %v", lexer.isLetter('1'), false)
	}
	if lexer.isLetter('６') {
		t.Errorf("bad result for isLetter(): got %v, expected %v", lexer.isLetter('６'), false)
	}
	if lexer.isLetter('#') {
		t.Errorf("bad result for isLetter(): got %v, expected %v", lexer.isLetter('#'), false)
	}
}

func TestIsDigit(t *testing.T) {
	filename := "test_file.liza"
	src := []byte("")

	lexer := New(filename, src, nil, 0)

	if !lexer.isDigit('0') {
		t.Errorf("bad result for isDigit(): got %v, expected %v", lexer.isDigit('0'), true)
	}
	if !lexer.isDigit('9') {
		t.Errorf("bad result for isDigit(): got %v, expected %v", lexer.isDigit('9'), true)
	}
	if !lexer.isDigit('６') {
		t.Errorf("bad result for isDigit(): got %v, expected %v", lexer.isDigit('６'), true)
	}

	if lexer.isDigit('a') {
		t.Errorf("bad result for isDigit(): got %v, expected %v", lexer.isDigit('a'), false)
	}
	if lexer.isDigit('#') {
		t.Errorf("bad result for isDigit(): got %v, expected %v", lexer.isDigit('#'), false)
	}
	if lexer.isDigit('ŝ') {
		t.Errorf("bad result for isDigit(): got %v, expected %v", lexer.isDigit('ŝ'), false)
	}
}

// test cases for tokens

var tokens = []*token.Token{
	// Special tokens
	{Type: token.COMMENT, Content: "// a comment \n"},
	{Type: token.COMMENT, Content: "//\r\n"},
	{Type: token.COMMENT, Content: "/* a comment */"},
	{Type: token.COMMENT, Content: "/* a multi-line comment\n a comment \n*/"},
	{Type: token.COMMENT, Content: "/*\r*/"},

	// Identifiers
	{Type: token.IDENT, Content: "foobar"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.IDENT, Content: "a۰۱۸"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.IDENT, Content: "foo६४"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.IDENT, Content: "bar９８７６"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.IDENT, Content: "ŝ"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.IDENT, Content: "ŝfoo"},
	{Type: token.NEWLINE, Content: "\n"},

	// Basic type literals
	{Type: token.INT, Content: "0"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.INT, Content: "1"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.INT, Content: "123456789012345678890"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.INT, Content: "0b0"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.INT, Content: "0B1010"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.INT, Content: "01234567"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.INT, Content: "0xcafebabe"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.INT, Content: "0Xcafebabe"},
	{Type: token.NEWLINE, Content: "\n"},

	{Type: token.FLOAT, Content: "0."},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.FLOAT, Content: ".0"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.FLOAT, Content: "3.14159265"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.FLOAT, Content: "1e0"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.FLOAT, Content: "1e+100"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.FLOAT, Content: "1e-100"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.FLOAT, Content: "2.71828e-1000"},
	{Type: token.NEWLINE, Content: "\n"},

	{Type: token.STRING, Content: "`foobar`"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.STRING, Content: `"foobar"`},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.STRING, Content: "`" + `foo
							                        bar` +
		"`",
	},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.STRING, Content: "`\r`"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.STRING, Content: "`foo\r\nbar`"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.STRING, Content: `"\\"`},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.STRING, Content: `"\000foo"`},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.STRING, Content: `"\xFFbar"`},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.STRING, Content: `"\uff16foo"`},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.STRING, Content: `"\U0000ff16bar"`},
	{Type: token.NEWLINE, Content: "\n"},

	// Operators and delimiters
	{Type: token.ADD, Content: "+"},
	{Type: token.SUB, Content: "-"},
	{Type: token.MUL, Content: "*"},
	{Type: token.DIV, Content: "/"},
	{Type: token.REM, Content: "%"},

	{Type: token.AND, Content: "&"},
	{Type: token.OR, Content: "|"},
	{Type: token.XOR, Content: "^"},
	{Type: token.SHL, Content: "<<"},
	{Type: token.SHR, Content: ">>"},

	{Type: token.ADDASSIGN, Content: "+="},
	{Type: token.SUBASSIGN, Content: "-="},
	{Type: token.MULASSIGN, Content: "*="},
	{Type: token.DIVASSIGN, Content: "/="},
	{Type: token.REMASSIGN, Content: "%="},

	{Type: token.ANDASSIGN, Content: "&="},
	{Type: token.ORASSIGN, Content: "|="},
	{Type: token.XORASSIGN, Content: "^="},
	{Type: token.SHLASSIGN, Content: "<<="},
	{Type: token.SHRASSIGN, Content: ">>="},

	{Type: token.LAND, Content: "&&"},
	{Type: token.LOR, Content: "||"},
	{Type: token.INC, Content: "++"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.DEC, Content: "--"},
	{Type: token.NEWLINE, Content: "\n"},

	{Type: token.EQL, Content: "=="},
	{Type: token.LSS, Content: "<"},
	{Type: token.GTR, Content: ">"},
	{Type: token.ASSIGN, Content: "="},
	{Type: token.NOT, Content: "!"},

	{Type: token.NEQ, Content: "!="},
	{Type: token.LEQ, Content: "<="},
	{Type: token.GEQ, Content: ">="},
	{Type: token.DEFINE, Content: ":="},

	{Type: token.LPAREN, Content: "("},
	{Type: token.LBRACK, Content: "["},
	{Type: token.LBRACE, Content: "{"},
	{Type: token.COMMA, Content: ","},
	{Type: token.PERIOD, Content: "."},

	{Type: token.RPAREN, Content: ")"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.RBRACK, Content: "]"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.RBRACE, Content: "}"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.COLON, Content: ":"},
	{Type: token.DOUBLECOLON, Content: "::"},
	{Type: token.SEMICOLON, Content: ";"},

	// Keywords
	{Type: token.AS, Content: "as"},
	{Type: token.BREAK, Content: "break"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.CASE, Content: "case"},
	{Type: token.CLASS, Content: "class"},
	{Type: token.CONST, Content: "const"},
	{Type: token.CONTINUE, Content: "continue"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.DEFAULT, Content: "default"},
	{Type: token.ELSE, Content: "else"},
	{Type: token.ENUM, Content: "enum"},
	{Type: token.FOR, Content: "for"},
	{Type: token.FUN, Content: "fun"},
	{Type: token.IF, Content: "if"},
	{Type: token.IMPLEMENTS, Content: "implements"},
	{Type: token.IMPORT, Content: "import"},
	{Type: token.INTERFACE, Content: "interface"},
	{Type: token.MATCH, Content: "match"},
	{Type: token.PACKAGE, Content: "package"},
	{Type: token.PUBLIC, Content: "public"},
	{Type: token.RETURN, Content: "return"},
	{Type: token.NEWLINE, Content: "\n"},
	{Type: token.VAR, Content: "var"},

	// EOF
	{Type: token.EOF, Content: ""},
}

const whitespaces = " \t \n\n\r\n"

var source = func() []byte {
	var src []byte
	for _, t := range tokens {
		src = append(src, t.Content...)
		src = append(src, whitespaces...)
	}
	return src
}()

func newlineCount(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			n++
		}
	}
	return n
}

func firstNewlineColumn(s string) int {
	n := 0
	for _, runeValue := range s {
		n++
		if runeValue == '\n' {
			return n
		}
	}
	// should never happen
	return -1
}

func lengthOfLastLine(s string) int {
	runes := []rune(s)
	len := len(runes)
	for i := len - 1; i >= 0; i-- {
		if runes[i] == '\n' {
			return len - 1 - i
		}
	}

	return len
}

func checkPos(t *testing.T, content string, p token.Position, expected token.Position) {
	if p.Filename != expected.Filename {
		t.Errorf("bad filename for %q: got %s, expected %s", content, p.Filename, expected.Filename)
	}
	if p.Line != expected.Line {
		t.Errorf("bad line for %q: got %d, expected %d", content, p.Line, expected.Line)
	}
	if p.Column != expected.Column {
		t.Errorf("bad column for %q: got %d, expected %d", content, p.Column, expected.Column)
	}
}

func TestNextToken(t *testing.T) {
	whitespacesLinecount := newlineCount(whitespaces)
	whitespacesFirstNewlineCol := firstNewlineColumn(whitespaces)

	filename := "test_file.liza"
	errHandler := func(pos token.Position, msg string) {
		t.Errorf("error handler called (msg = %s", msg)
	}

	lexer := New(filename, source, errHandler, ScanComments)

	epos := token.Position{
		Filename: filename,
		Line:     1,
		Column:   1,
	}

	for i, etk := range tokens {
		tk := lexer.NextToken()

		// check token type
		if tk.Type != etk.Type {
			t.Errorf("bad token for %q: got %s, expected %s", tk.Content, tk.Type, etk.Type)
		}

		// check token position
		if tk.Type == token.NEWLINE {
			// NEWLINE is actually in last token or the appended whitespaces
			pos := epos
			pos.Line -= whitespacesLinecount
			pos.Column = lengthOfLastLine(tokens[i-1].Content) + whitespacesFirstNewlineCol

			checkPos(t, tk.Content, tk.Position, pos)
		} else {
			if tk.Type == token.EOF {
				// correct for EOF: it is last line plus 1
				epos.Line = newlineCount(string(source)) + 1
			}
			checkPos(t, tk.Content, tk.Position, epos)
		}

		// check content
		eContent := etk.Content
		switch etk.Type {
		case token.COMMENT:
			// no CRs in comments
			eContent = string(lexer.stripCR([]byte(etk.Content)))
			//-style comment doesn't content newline
			if etk.Content[1] == '/' {
				eContent = eContent[0 : len(eContent)-1]
			}
		case token.STRING:
			// no CRs in raw string literals
			if eContent[0] == '`' {
				eContent = string(lexer.stripCR([]byte(eContent)))
			}
		}

		if tk.Content != eContent {
			t.Errorf("bad content for %q: got %q, expected %q", tk.Content, tk.Content, eContent)
		}

		// update position for next token
		epos.Line += newlineCount(etk.Content) + whitespacesLinecount
	}
}

func checkNewline(t *testing.T, line string, mode Mode) {
	filename := "test_file.liza"
	lexer := New(filename, []byte(line), nil, mode)

	tok := lexer.NextToken()
	for tok.Type != token.EOF {
		if tok.Type == token.ILLEGAL { // the illegal token literal indicates there must be a newline following
			// next token must be a newline
			pos := token.Position{
				Filename: filename,
				Line:     tok.Position.Line,
				Column:   tok.Position.Column + 1,
			}
			tok = lexer.NextToken()
			if tok.Type == token.NEWLINE {
				if tok.Content != "\n" {
					t.Errorf(`bad content for %q: got %q, expected %q`, line, tok.Content, "\n")
				}
				checkPos(t, line, tok.Position, pos)
			} else {
				t.Errorf("bad token for %q: got %s, expected NEWLINE", line, tok)
			}
		} else if tok.Type == token.NEWLINE {
			t.Errorf("bad token for %q: got NEWLINE, expected no NEWLINE", line)
		}
		tok = lexer.NextToken()
	}
}

var lines = []string{
	// $ indicates an automatically inserted newline
	"",
	"\ufeff\n", // first BOM is ignored
	"\n",
	"foo$\n",
	"123$\n",
	"1.2$\n",
	`"x"` + "$\n",
	"`x`$\n",

	"+\n",
	"-\n",
	"*\n",
	"/\n",
	"%\n",

	"&\n",
	"|\n",
	"^\n",
	"<<\n",
	">>\n",

	"+=\n",
	"-=\n",
	"*=\n",
	"/=\n",
	"%=\n",

	"&=\n",
	"|=\n",
	"^=\n",
	"<<=\n",
	">>=\n",

	"&&\n",
	"||\n",
	"++$\n",
	"--$\n",

	"==\n",
	"<\n",
	">\n",
	"=\n",
	"!\n",

	"!=\n",
	"<=\n",
	">=\n",
	":=\n",

	"(\n",
	"[\n",
	"{\n",
	",\n",
	".\n",

	")$\n",
	"]$\n",
	"}$\n",
	":\n",
	"::\n",
	";\n",

	"as\n",
	"break$\n",
	"case\n",
	"class\n",
	"const\n",

	"continue$\n",
	"default\n",
	"else\n",
	"enum\n",
	"for\n",

	"fun\n",
	"if\n",
	"implements\n",
	"import\n",
	"interface\n",

	"match\n",
	"package\n",
	"public\n",
	"return$\n",

	"foo /*comment*/ /=",

	"foo$//comment\n",
	"foo$//comment",
	"foo$/*comment*/\n",
	"foo$/*\n*/",
	"foo$/*comment*/    \n",
	"foo$/*\n*/    ",

	"foo    $// comment\n",
	"foo    $// comment",
	"foo    $/*comment*/\n",
	"foo    $/*\n*/",
	"foo    $/*  */ /* \n */ bar$/**/\n",
	"foo    $/*0*/ /*1*/ /*2*/\n",

	"foo    $/*comment*/    \n",
	"foo    $/*0*/ /*1*/ /*2*/    \n",
	"foo	$/**/ /*-------------*/       /*----\n*/bar       $/*  \n*/baa$\n",
	"foo    $/* an EOF terminates a line */",
	"foo    $/* an EOF terminates a line */ /*",
	"foo    $/* an EOF terminates a line */ //",

	"package main$",
	"package main$\n\nfun main() {\n\tif {\n\t\treturn /* */ }$\n}$\n",
}

func TestNewline(t *testing.T) {
	for _, line := range lines {
		checkNewline(t, line, 0)
		checkNewline(t, line, ScanComments)

		// if the input ended in newlines, the input must tokenize the
		// same with or without those newlines
		for i := len(line) - 1; i >= 0 && line[i] == '\n'; i-- {
			checkNewline(t, line[0:i], 0)
			checkNewline(t, line[0:i], ScanComments)
		}
	}
}

func checkError(t *testing.T, src string, etok token.Type, content string, cols []int, errs []string) {
	filename := "test_file.liza"

	var (
		errCount int
		errMsgs  []string
		errPoses []token.Position
	)
	errHandler := func(pos token.Position, msg string) {
		errCount++
		errPoses = append(errPoses, pos)
		errMsgs = append(errMsgs, msg)
	}

	lexer := New(filename, []byte(src), errHandler, ScanComments)

	tok := lexer.NextToken()
	if tok.Type != etok {
		t.Errorf("%q: got %s, expected %s", src, tok.Type, etok)
	}
	if tok.Type != token.ILLEGAL && tok.Content != content {
		t.Errorf("%q: got content %q, expected %q", src, tok.Content, content)
	}

	if errCount != len(errs) {
		t.Errorf("%q: got errCount %d, expected %d", src, errCount, len(errs))
	} else {
		for i, err := range errs {
			if errMsgs[i] != err {
				t.Errorf("%q: got errMsg %q, expected %q", src, errMsgs[i], err)
			}
			if errPoses[i].Column != cols[i] {
				t.Errorf("%q, got column %d, expected %d", src, errPoses[i].Column, cols[i])
			}
		}
	}
}

var errors = []struct {
	src     string
	tokType token.Type
	cols    []int
	content string
	errs    []string
}{
	{"\a", token.ILLEGAL, []int{1}, "", []string{"illegal character U+0007"}},
	{`#`, token.ILLEGAL, []int{1}, "", []string{"illegal character U+0023 '#'"}},
	{`…`, token.ILLEGAL, []int{1}, "", []string{"illegal character U+2026 '…'"}},
	{"0x", token.INT, []int{1}, "0x", []string{"illegal hexadecimal number"}},
	{"0X", token.INT, []int{1}, "0X", []string{"illegal hexadecimal number"}},
	{"0b", token.INT, []int{1}, "0b", []string{"illegal binary number"}},
	{"0B", token.INT, []int{1}, "0B", []string{"illegal binary number"}},
	{"078", token.INT, []int{1}, "078", []string{"illegal octal number"}},
	{"07800000009", token.INT, []int{1}, "07800000009", []string{"illegal octal number"}},
	{"0E", token.FLOAT, []int{1}, "0E", []string{"illegal floating-point exponent"}},
	{`"\8"`, token.STRING, []int{3}, `"\8"`, []string{"unknown escape sequence"}},
	{`"\`, token.STRING, []int{3, 1}, `"\`, []string{"escape sequence not terminated", "string literal not terminated"}},

	{`"\0"`, token.STRING, []int{3}, `"\0"`, []string{"illegal character U+0022 '\"' in escape sequence"}},
	{`"\07"`, token.STRING, []int{3}, `"\07"`, []string{"illegal character U+0022 '\"' in escape sequence"}},
	{`"\08"`, token.STRING, []int{3}, `"\08"`, []string{"illegal character U+0038 '8' in escape sequence"}},
	{`"\0`, token.STRING, []int{3, 1}, `"\0`, []string{"escape sequence not terminated", "string literal not terminated"}},

	{`"\x"`, token.STRING, []int{3}, `"\x"`, []string{"illegal character U+0022 '\"' in escape sequence"}},
	{`"\x0"`, token.STRING, []int{3}, `"\x0"`, []string{"illegal character U+0022 '\"' in escape sequence"}},
	{`"\x0g"`, token.STRING, []int{3}, `"\x0g"`, []string{"illegal character U+0067 'g' in escape sequence"}},
	{`"\x`, token.STRING, []int{3, 1}, `"\x`, []string{"escape sequence not terminated", "string literal not terminated"}},

	{`"\u"`, token.STRING, []int{3}, `"\u"`, []string{"illegal character U+0022 '\"' in escape sequence"}},
	{`"\u0"`, token.STRING, []int{3}, `"\u0"`, []string{"illegal character U+0022 '\"' in escape sequence"}},
	{`"\u00"`, token.STRING, []int{3}, `"\u00"`, []string{"illegal character U+0022 '\"' in escape sequence"}},
	{`"\u000"`, token.STRING, []int{3}, `"\u000"`, []string{"illegal character U+0022 '\"' in escape sequence"}},
	{`"\u000`, token.STRING, []int{3, 1}, `"\u000`, []string{"escape sequence not terminated", "string literal not terminated"}},

	{`"\U"`, token.STRING, []int{3}, `"\U"`, []string{"illegal character U+0022 '\"' in escape sequence"}},
	{`"\U0"`, token.STRING, []int{3}, `"\U0"`, []string{"illegal character U+0022 '\"' in escape sequence"}},
	{`"\U00"`, token.STRING, []int{3}, `"\U00"`, []string{"illegal character U+0022 '\"' in escape sequence"}},
	{`"\U000"`, token.STRING, []int{3}, `"\U000"`, []string{"illegal character U+0022 '\"' in escape sequence"}},
	{`"\U0000"`, token.STRING, []int{3}, `"\U0000"`, []string{"illegal character U+0022 '\"' in escape sequence"}},
	{`"\U00000"`, token.STRING, []int{3}, `"\U00000"`, []string{"illegal character U+0022 '\"' in escape sequence"}},
	{`"\U000000"`, token.STRING, []int{3}, `"\U000000"`, []string{"illegal character U+0022 '\"' in escape sequence"}},
	{`"\U0000000"`, token.STRING, []int{3}, `"\U0000000"`, []string{"illegal character U+0022 '\"' in escape sequence"}},
	{`"\U0000000`, token.STRING, []int{3, 1}, `"\U0000000`, []string{"escape sequence not terminated", "string literal not terminated"}},

	{`"\Uffffffff"`, token.STRING, []int{3}, `"\Uffffffff"`, []string{"escape sequence is invalid Unicode code point"}},

	{`"abc`, token.STRING, []int{1}, `"abc`, []string{"string literal not terminated"}},
	{"\"abc\n", token.STRING, []int{1}, `"abc`, []string{"string literal not terminated"}},
	{"\"abc\n   ", token.STRING, []int{1}, `"abc`, []string{"string literal not terminated"}},
	{"`", token.STRING, []int{1}, "`", []string{"raw string literal not terminated"}},
	{"\"abc\x00def\"", token.STRING, []int{5}, "\"abc\x00def\"", []string{"illegal character NULL"}},
	{"\"abc\x80def\"", token.STRING, []int{5}, "\"abc\x80def\"", []string{"illegal UTF-8 encoding"}},

	{"/*", token.COMMENT, []int{1}, "/*", []string{"comment not terminated"}},

	// only first BOM is ignored
	{"\ufeff\ufeff", token.ILLEGAL, []int{2}, "\ufeff\ufeff", []string{"illegal byte order mark"}},
	{"//\ufeff", token.COMMENT, []int{3}, "//\ufeff", []string{"illegal byte order mark"}},
	{`"` + "abc\ufeffdef" + `"`, token.STRING, []int{5}, `"` + "abc\ufeffdef" + `"`, []string{"illegal byte order mark"}},
}

func TestNextTokenErrors(t *testing.T) {
	for _, e := range errors {
		checkError(t, e.src, e.tokType, e.content, e.cols, e.errs)
	}
}
