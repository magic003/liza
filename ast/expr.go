package ast

import (
	"github.com/magic003/liza/token"
)

// Expr is the base type for all expression tree node.
type Expr interface {
	Node
	exprNode()
}

// Ident is a node represents an identifier.
type Ident struct {
	token token.Token // identifier token
}

// Pos implementation for Node.
func (ident *Ident) Pos() token.Position {
	return ident.token.Position
}

// End implementation for Node.
func (ident *Ident) End() token.Position {
	return token.Position{
		Filename: ident.token.Position.Filename,
		Line:     ident.token.Position.Line,
		Column:   ident.token.Position.Column + len(ident.token.Content),
	}
}

func (ident *Ident) exprNode() {}
