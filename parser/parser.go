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
		p.errorExpected(currentTok.Position, "<"+tt.String()+">")
	}
	p.next() // make progress
	return currentTok
}

func (p *Parser) errorExpected(pos token.Position, expected string) {
	msg := "expected " + expected
	if pos == p.tok.Position {
		// error happens at the current position. make it more specific
		msg += ", found <" + p.tok.Type.String() + "> " + p.tok.Content
	}

	p.handleErr(pos, msg)
}

func (p *Parser) syncTopLevelDecl() {
	for {
		switch p.tok.Type {
		case token.PUBLIC, token.CONST, token.CLASS, token.INTERFACE:
			return
		case token.EOF:
			return
		}
		p.next()
	}
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

func (p *Parser) parseTopLevelDecl() ast.Decl {
	var visibility *token.Token
	if p.tok.Type == token.PUBLIC {
		visibility = p.expect(token.PUBLIC)
	}

	switch p.tok.Type {
	case token.CONST:
		return p.parseConstDecl(visibility)
	case token.CLASS:
		return p.parseClassDecl(visibility)
	case token.INTERFACE:
		return p.parseInterfaceDecl(visibility)
	default:
		pos := p.tok.Position
		p.errorExpected(p.tok.Position, "declaration")
		p.syncTopLevelDecl()
		return &ast.BadDecl{
			From: pos,
			To:   p.tok.Position,
		}
	}
}

func (p *Parser) parseConstDecl(visibility *token.Token) *ast.ConstDecl {
	constPos := p.expect(token.CONST).Position
	ident := p.expect(token.IDENT)
	var tp ast.Type
	if p.tok.Type != token.DEFINE {
		tp = p.parseType()
	}
	p.expect(token.DEFINE)
	value := p.parseExpr()
	p.expect(token.NEWLINE)
	return &ast.ConstDecl{
		Visibility: visibility,
		Const:      constPos,
		Ident:      ident,
		Type:       tp,
		Value:      value,
	}
}

func (p *Parser) parseClassDecl(visibility *token.Token) *ast.ClassDecl {
	return nil
}

func (p *Parser) parseInterfaceDecl(visibility *token.Token) *ast.InterfaceDecl {
	interfacePos := p.expect(token.INTERFACE).Position
	name := p.expect(token.IDENT)
	lbrace := p.expect(token.LBRACE).Position
	var methods []*ast.FuncDef
	for p.tok.Type == token.FUN {
		methods = append(methods, p.parseFuncDef())
	}
	rbrace := p.expect(token.RBRACE).Position
	p.expect(token.NEWLINE)
	return &ast.InterfaceDecl{
		Visibility: visibility,
		Interface:  interfacePos,
		Name:       name,
		Lbrace:     lbrace,
		Methods:    methods,
		Rbrace:     rbrace,
	}
}

func (p *Parser) parseFuncDef() *ast.FuncDef {
	funcDef := p.parseFuncSignature()
	p.expect(token.NEWLINE)
	return funcDef
}

func (p *Parser) parseFuncSignature() *ast.FuncDef {
	fun := p.expect(token.FUN).Position
	name := p.expect(token.IDENT)
	lparen := p.expect(token.LPAREN).Position

	var params []*ast.ParameterDef
	for p.tok.Type != token.RPAREN && p.tok.Type != token.EOF {
		params = append(params, p.parseParameterDef())
		if p.tok.Type != token.RPAREN {
			p.expect(token.COMMA)
		}
	}
	rparen := p.expect(token.RPAREN).Position
	var tp ast.Type
	if p.tok.Type == token.COLON {
		p.expect(token.COLON)
		tp = p.parseType()
	}
	return &ast.FuncDef{
		Fun:        fun,
		Name:       name,
		Lparen:     lparen,
		Params:     params,
		Rparen:     rparen,
		ReturnType: tp,
	}
}

func (p *Parser) parseParameterDef() *ast.ParameterDef {
	ident := p.expect(token.IDENT)
	tp := p.parseType()
	return &ast.ParameterDef{
		Name: ident,
		Type: tp,
	}
}

func (p *Parser) parseVarDecl() *ast.VarDecl {
	ident := p.expect(token.IDENT)
	var tp ast.Type
	if p.tok.Type != token.DEFINE {
		tp = p.parseType()
	}
	p.expect(token.DEFINE)
	value := p.parseExpr()
	p.expect(token.NEWLINE)
	return &ast.VarDecl{
		Ident: ident,
		Type:  tp,
		Value: value,
	}
}

// ---------------------------------------------------------------------------
// Expression

func (p *Parser) parseExpr() ast.Expr {
	return p.parseBinaryExpr(lowestPrec)
}

func (p *Parser) parseBinaryExpr(prec int) ast.Expr {
	x := p.parseUnaryExpr()
	for {
		currPrec := precedence(p.tok.Type)
		if currPrec <= prec {
			return x
		}
		op := p.expect(p.tok.Type)
		y := p.parseBinaryExpr(currPrec)
		x = &ast.BinaryExpr{X: x, Op: op, Y: y}
	}
}

func (p *Parser) parseUnaryExpr() ast.Expr {
	tp := p.tok.Type
	if tp == token.SUB || tp == token.XOR || tp == token.NOT {
		op := p.expect(tp)
		x := p.parseUnaryExpr()
		return &ast.UnaryExpr{Op: op, X: x}
	}

	return p.parsePrimaryExpr()
}

func (p *Parser) parsePrimaryExpr() ast.Expr {
	x := p.parseOperand()

	for {
		switch p.tok.Type {
		case token.PERIOD:
			x = p.parseSelectorExpr(x)
		case token.LBRACK:
			x = p.parseIndexExpr(x)
		case token.LPAREN:
			x = p.parseCallExpr(x)
		default:
			return x
		}
	}
}

func (p *Parser) parseOperand() ast.Expr {
	switch p.tok.Type {
	case token.IDENT:
		return p.parseIdent()
	case token.INT, token.FLOAT, token.STRING:
		tok := p.expect(p.tok.Type)
		return &ast.BasicLit{Token: tok}
	case token.LBRACK:
		return p.parseArrayLit()
	case token.LBRACE:
		return p.parseMapLit()
	case token.LPAREN:
		// It could not tell it is a TupleLit or ParenExpr, so it always treats it as a TupleLit.
		// The actual check will happen in the semantic analysis phase.
		return p.parseTupleLit()
	}

	// TODO record error, sync and return BadExpr
	return nil
}

func (p *Parser) parseArrayLit() *ast.ArrayLit {
	lbrack := p.expect(token.LBRACK).Position
	var elts []ast.Expr
	for p.tok.Type != token.RBRACK && p.tok.Type != token.EOF {
		elts = append(elts, p.parseExpr())
		if p.tok.Type != token.RBRACK {
			p.expect(token.COMMA)
		}
	}
	rbrack := p.expect(token.RBRACK).Position
	return &ast.ArrayLit{
		Lbrack: lbrack,
		Elts:   elts,
		Rbrack: rbrack,
	}
}

func (p *Parser) parseMapLit() *ast.MapLit {
	lbrace := p.expect(token.LBRACE).Position
	var elts []*ast.KeyValueExpr
	for p.tok.Type != token.RBRACE && p.tok.Type != token.EOF {
		elts = append(elts, p.parseKeyValueExpr())
		if p.tok.Type != token.RBRACE {
			p.expect(token.COMMA)
		}
	}
	rbrace := p.expect(token.RBRACE).Position
	return &ast.MapLit{
		Lbrace: lbrace,
		Elts:   elts,
		Rbrace: rbrace,
	}
}

func (p *Parser) parseKeyValueExpr() *ast.KeyValueExpr {
	key := p.parseExpr()
	colon := p.expect(token.COLON).Position
	value := p.parseExpr()
	return &ast.KeyValueExpr{
		Key:   key,
		Colon: colon,
		Value: value,
	}
}

func (p *Parser) parseTupleLit() *ast.TupleLit {
	lparen := p.expect(token.LPAREN).Position
	var elts []ast.Expr
	for p.tok.Type != token.RPAREN && p.tok.Type != token.EOF {
		elts = append(elts, p.parseExpr())
		if p.tok.Type != token.RPAREN {
			p.expect(token.COMMA)
		}
	}
	rparen := p.expect(token.RPAREN).Position
	return &ast.TupleLit{
		Lparen: lparen,
		Elts:   elts,
		Rparen: rparen,
	}
}

func (p *Parser) parseSelectorExpr(x ast.Expr) *ast.SelectorExpr {
	p.expect(token.PERIOD)
	sel := p.parseIdent()
	return &ast.SelectorExpr{X: x, Sel: sel}
}

func (p *Parser) parseIdent() *ast.Ident {
	ident := p.expect(token.IDENT)
	return &ast.Ident{Token: ident}
}

func (p *Parser) parseIndexExpr(x ast.Expr) *ast.IndexExpr {
	lbrack := p.expect(token.LBRACK).Position
	index := p.parseExpr()
	rbrack := p.expect(token.RBRACK).Position
	return &ast.IndexExpr{
		X:      x,
		Lbrack: lbrack,
		Index:  index,
		Rbrack: rbrack,
	}
}

func (p *Parser) parseCallExpr(x ast.Expr) *ast.CallExpr {
	lparen := p.expect(token.LPAREN).Position
	var args []ast.Expr
	for p.tok.Type != token.RPAREN && p.tok.Type != token.EOF {
		args = append(args, p.parseExpr())
		if p.tok.Type != token.RPAREN {
			p.expect(token.COMMA)
		}
	}
	rparen := p.expect(token.RPAREN).Position
	return &ast.CallExpr{
		Fun:    x,
		Lparen: lparen,
		Args:   args,
		Rparen: rparen,
	}
}

// ---------------------------------------------------------------------------
// Type

func (p *Parser) parseType() ast.Type {
	switch p.tok.Type {
	case token.LBRACK:
		return p.parseArrayType()
	case token.LBRACE:
		return p.parseMapType()
	case token.LPAREN:
		return p.parseTupleType()
	case token.IDENT:
		return p.parseBasicOrSelectorType()
	default:
		// TODO record error, sync and return bad type node.
		return nil
	}
}

func (p *Parser) parseArrayType() *ast.ArrayType {
	lbrack := p.expect(token.LBRACK).Position
	rbrack := p.expect(token.RBRACK).Position
	elt := p.parseType()
	return &ast.ArrayType{
		Lbrack: lbrack,
		Rbrack: rbrack,
		Elt:    elt,
	}
}

func (p *Parser) parseMapType() *ast.MapType {
	lbrace := p.expect(token.LBRACE).Position
	key := p.parseType()
	p.expect(token.COLON)
	value := p.parseType()
	rbrace := p.expect(token.RBRACE).Position
	return &ast.MapType{
		Lbrace: lbrace,
		Key:    key,
		Value:  value,
		Rbrace: rbrace,
	}
}

func (p *Parser) parseTupleType() *ast.TupleType {
	lparen := p.expect(token.LPAREN).Position
	var elts []ast.Type
	for p.tok.Type != token.RPAREN && p.tok.Type != token.EOF {
		elts = append(elts, p.parseType())
		if p.tok.Type != token.RPAREN {
			p.expect(token.COMMA)
		}
	}
	rparen := p.expect(token.RPAREN).Position
	return &ast.TupleType{
		Lparen: lparen,
		Elts:   elts,
		Rparen: rparen,
	}
}

func (p *Parser) parseBasicOrSelectorType() ast.Type {
	ident1 := p.expect(token.IDENT)
	if p.tok.Type == token.PERIOD {
		p.expect(token.PERIOD)
		ident2 := p.expect(token.IDENT)
		return &ast.SelectorType{
			Package: ident1,
			Sel:     ident2,
		}
	}

	return &ast.BasicType{
		Ident: ident1,
	}
}
