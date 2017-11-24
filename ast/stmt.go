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
	Expr Expr        // expression
	Op   token.Token // INC or DEC
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
	Assign token.Token // assignment token
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
	Return token.Token // position of "return"
	Value  Expr        // returned value expression, optional
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
	Tok token.Token // keyword token (break, continue)
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
