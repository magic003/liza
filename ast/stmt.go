package ast

import (
	"github.com/magic003/liza/token"
)

// Stmt is the base type for all statement tree nodes.
type Stmt interface {
	Node
	stmtNode()
}

// DeclStmt node represents a statment with const/var declaration.
type DeclStmt struct {
	Decl Decl // constant or variable declaration
}

// Pos implementation for Node.
func (stmt *DeclStmt) Pos() token.Position {
	return stmt.Decl.Pos()
}

// End implementation for Node.
func (stmt *DeclStmt) End() token.Position {
	return stmt.Decl.End()
}

func (stmt *DeclStmt) stmtNode() {}

// ExprStmt node represents a standalone expression.
type ExprStmt struct {
	Expr Expr // expression
}

// Pos implementation for Node.
func (stmt *ExprStmt) Pos() token.Position {
	return stmt.Expr.Pos()
}

// End implementation for Node.
func (stmt *ExprStmt) End() token.Position {
	return stmt.Expr.End()
}

func (stmt *ExprStmt) stmtNode() {}

// IncDecStmt node represents an increasement or decreasement statement.
type IncDecStmt struct {
	Expr Expr         // expression
	Op   *token.Token // INC or DEC
}

// Pos implementation for Node.
func (stmt *IncDecStmt) Pos() token.Position {
	return stmt.Expr.Pos()
}

// End implementation for Node.
func (stmt *IncDecStmt) End() token.Position {
	return token.Position{
		Filename: stmt.Op.Position.Filename,
		Line:     stmt.Op.Position.Line,
		Column:   stmt.Op.Position.Column + len(stmt.Op.Content),
	}
}

func (stmt *IncDecStmt) stmtNode() {}

// AssignStmt node represents an assignment statement.
type AssignStmt struct {
	LHS    Expr
	Assign *token.Token // assignment token
	RHS    Expr
}

// Pos implementation for Node.
func (stmt *AssignStmt) Pos() token.Position {
	return stmt.LHS.Pos()
}

// End implementation for Node.
func (stmt *AssignStmt) End() token.Position {
	return stmt.RHS.End()
}

func (stmt *AssignStmt) stmtNode() {}

// ReturnStmt node represents a return statement.
type ReturnStmt struct {
	Return *token.Token // position of "return"
	Value  Expr         // returned value expression, optional
}

// Pos implementation for Node.
func (stmt *ReturnStmt) Pos() token.Position {
	return stmt.Return.Position
}

// End implementation for Node.
func (stmt *ReturnStmt) End() token.Position {
	if stmt.Value != nil {
		return stmt.Value.End()
	}

	return token.Position{
		Filename: stmt.Return.Position.Filename,
		Line:     stmt.Return.Position.Line,
		Column:   stmt.Return.Position.Column + len(stmt.Return.Content),
	}
}

func (stmt *ReturnStmt) stmtNode() {}

// BranchStmt node represents a break or continue statement.
type BranchStmt struct {
	Tok *token.Token // keyword token (break, continue)
}

// Pos implementation for Node.
func (stmt *BranchStmt) Pos() token.Position {
	return stmt.Tok.Position
}

// End implementation for Node.
func (stmt *BranchStmt) End() token.Position {
	return token.Position{
		Filename: stmt.Tok.Position.Filename,
		Line:     stmt.Tok.Position.Line,
		Column:   stmt.Tok.Position.Column + len(stmt.Tok.Content),
	}
}

func (stmt *BranchStmt) stmtNode() {}

// BlockStmt node represents a braced statement list.
type BlockStmt struct {
	Lbrace token.Position // position of "{"
	Stmts  []Stmt
	Rbrace token.Position // position of "}"
}

// Pos implementation for Node.
func (stmt *BlockStmt) Pos() token.Position {
	return stmt.Lbrace
}

// End implementation for Node.
func (stmt *BlockStmt) End() token.Position {
	return token.Position{
		Filename: stmt.Rbrace.Filename,
		Line:     stmt.Rbrace.Line,
		Column:   stmt.Rbrace.Column + 1,
	}
}

func (stmt *BlockStmt) stmtNode() {}

// IfStmt node represents an if statement.
type IfStmt struct {
	If   token.Position // position of "if"
	Cond Expr           // condition
	Body *BlockStmt
	Else *ElseStmt // optional else statement
}

// Pos implementation for Node.
func (stmt *IfStmt) Pos() token.Position {
	return stmt.If
}

// End implementation for Node.
func (stmt *IfStmt) End() token.Position {
	if stmt.Else != nil {
		return stmt.Else.End()
	}

	return stmt.Body.End()
}

func (stmt *IfStmt) stmtNode() {}

// ElseStmt node represents the else clause of an if statement.
type ElseStmt struct {
	Else token.Position // position of "else"
	If   *IfStmt        // optional else if statement
	Body *BlockStmt     // body; nil if it has an if statement
}

// Pos implementation for Node.
func (stmt *ElseStmt) Pos() token.Position {
	return stmt.Else
}

// End implementation for Node.
func (stmt *ElseStmt) End() token.Position {
	if stmt.If != nil {
		return stmt.If.End()
	}

	return stmt.Body.End()
}

func (stmt *ElseStmt) stmtNode() {}

// MatchStmt node represents a match statement.
type MatchStmt struct {
	Match  token.Position // position of "match"
	Expr   Expr
	Lbrace token.Position // position of "{"
	Cases  []*CaseClause
	Rbrace token.Position // position of "}"
}

// Pos implementation for Node.
func (stmt *MatchStmt) Pos() token.Position {
	return stmt.Match
}

// End implementation for Node.
func (stmt *MatchStmt) End() token.Position {
	return token.Position{
		Filename: stmt.Rbrace.Filename,
		Line:     stmt.Rbrace.Line,
		Column:   stmt.Rbrace.Column + 1,
	}
}

func (stmt *MatchStmt) stmtNode() {}

// CaseClause node represents a case clause in a match statement.
type CaseClause struct {
	Case    token.Position // position of case or default
	Pattern Expr           // matched pattern
	Colon   token.Position // position of ":"
	Body    []Stmt         // optional statement list
}

// Pos implementation for Node.
func (stmt *CaseClause) Pos() token.Position {
	return stmt.Case
}

// End implementation for Node.
func (stmt *CaseClause) End() token.Position {
	if n := len(stmt.Body); n > 0 {
		return stmt.Body[n-1].End()
	}

	return stmt.Colon
}

func (stmt *CaseClause) stmtNode() {}

// ForStmt node represents a for loop statement.
type ForStmt struct {
	For   token.Position // position of "for"
	Decls []Decl         // list of var or const declarations
	Cond  Expr           // condition; or nil
	Post  Stmt           // post iteration statement; or nil
	Body  *BlockStmt
}

// Pos implementation for Node.
func (stmt *ForStmt) Pos() token.Position {
	return stmt.For
}

// End implementation for Node.
func (stmt *ForStmt) End() token.Position {
	return stmt.Body.End()
}

func (stmt *ForStmt) stmtNode() {}
