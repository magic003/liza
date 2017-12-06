package parser

import (
	"github.com/magic003/liza/lexer"
	"github.com/magic003/liza/token"
)

// New returns a new instance of parser.
func New(filename string, src []byte) *Parser {
	parser := &Parser{}

	errHandler := func(pos token.Position, msg string) {
		err := Error{
			Pos: pos,
			Msg: msg,
		}
		parser.errors = append(parser.errors, err)
	}

	lexer := lexer.New(filename, src, errHandler, lexer.ScanComments)
	parser.lexer = lexer

	return parser
}

// Parser holds the internal state of a parser.
type Parser struct {
	lexer *lexer.Lexer

	errors []Error
}
