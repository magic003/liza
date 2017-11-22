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
