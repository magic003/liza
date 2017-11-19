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
	NamePos token.Position // identifier position
	Name    string         // identifier name
}

// Pos implementation for Node.
func (ident *Ident) Pos() token.Position {
	return ident.NamePos
}

// End implementation for Node.
func (ident *Ident) End() token.Position {
	return token.Position{
		Filename: ident.NamePos.Filename,
		Line:     ident.NamePos.Line,
		Column:   ident.NamePos.Column + len(ident.Name),
	}
}

func (ident *Ident) exprNode() {}
