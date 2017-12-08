package parser

import (
	"github.com/magic003/liza/ast"
	"github.com/magic003/liza/lexer"
	"github.com/magic003/liza/token"
)

// New returns a new instance of parser.
func New(filename string, src []byte) *Parser {
	parser := &Parser{}

	lexer := lexer.New(filename, src, parser.handleErr, lexer.ScanComments)
	parser.lexer = lexer

	parser.next()

	return parser
}

// Parser holds the internal state of a parser.
type Parser struct {
	lexer *lexer.Lexer

	tok *token.Token // current token

	errors []Error
}

// ---------------------------------------------------------------------------
// Parsing utilities

// next advances to the next non-comment token.
func (p *Parser) next() {
	p.next0()

	for p.tok.Type == token.COMMENT {
		p.next0()
	}
}

// next0 advances to the next token.
func (p *Parser) next0() {
	p.tok = p.lexer.NextToken()
}

func (p *Parser) handleErr(pos token.Position, msg string) {
	err := Error{
		Pos: pos,
		Msg: msg,
	}
	p.errors = append(p.errors, err)
}

func (p *Parser) expect(tt token.Type) *token.Token {
	currentTok := p.tok
	if currentTok.Type != tt {
		p.handleErr(currentTok.Position, "expected "+"'"+tt.String()+"' found '"+currentTok.Type.String()+"'")
	}
	p.next() // make progress
	return currentTok
}

// ---------------------------------------------------------------------------
// Declarations

func (p *Parser) parsePackageDecl() *ast.PackageDecl {
	packagePos := p.expect(token.PACKAGE).Position
	name := p.expect(token.IDENT)
	p.expect(token.NEWLINE)
	return &ast.PackageDecl{
		Package: packagePos,
		Name:    name,
	}
}

func (p *Parser) parseImportDecl() *ast.ImportDecl {
	importPos := p.expect(token.IMPORT).Position
	path := p.parseImportPath()

	node := &ast.ImportDecl{
		Import: importPos,
		Path:   path,
	}
	if p.tok.Type == token.AS {
		asPos := p.expect(token.AS).Position
		alias := p.expect(token.IDENT)

		node.As = &asPos
		node.Alias = alias
	}
	p.expect(token.NEWLINE)

	return node
}

func (p *Parser) parseImportPath() *ast.ImportPath {
	var (
		libraryName *token.Token
		path        []*token.Token
	)

	ident := p.expect(token.IDENT)
	if p.tok.Type == token.DOUBLECOLON {
		libraryName = ident
		p.expect(token.DOUBLECOLON)
		ident = p.expect(token.IDENT)
	}
	path = append(path, ident)

	for p.tok.Type == token.DIV {
		p.expect(token.DIV)
		ident = p.expect(token.IDENT)
		path = append(path, ident)
	}

	return &ast.ImportPath{
		LibraryName: libraryName,
		Path:        path,
	}
}
