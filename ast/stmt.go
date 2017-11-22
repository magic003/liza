package ast

import (
	"github.com/magic003/liza/token"
)

// Stmt is the base type for all statment tree nodes.
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
