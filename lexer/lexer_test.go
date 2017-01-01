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
	if lexer.isLetter('#') {
		t.Errorf("bad result for isLetter(): got %v, expected %v", lexer.isLetter('#'), false)
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

	for _, etk := range tokens {
		tk := lexer.NextToken()

		// check token type
		if tk.Type != etk.Type {
			t.Errorf("bad token for %q: got %s, expected %s", tk.Content, tk.Type, etk.Type)
		}

		// check token position
		if tk.Type == token.EOF {
			// correct for EOF: it is last line plus 1
			epos.Line = newlineCount(string(source)) + 1
		}
		checkPos(t, tk.Content, tk.Position, epos)

		// check content
		var eContent string
		switch etk.Type {
		case token.COMMENT:
			// no CRs in comments
			eContent = string(lexer.stripCR([]byte(etk.Content)))
			//-style comment doesn't content newline
			if etk.Content[1] == '/' {
				eContent = eContent[0 : len(eContent)-1]
			}
		default:
			eContent = etk.Content
		}

		if tk.Content != eContent {
			t.Errorf("bad content for %q: got %q, expected %q", tk.Content, tk.Content, eContent)
		}

		// update position for next token
		epos.Line += newlineCount(etk.Content) + whitespacesLinecount
	}
}
