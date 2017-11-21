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

// BasicLit is a node represents a literal of basic type.
type BasicLit struct {
	token token.Token // basic literal token
}

// Pos implementation for Node.
func (lit *BasicLit) Pos() token.Position {
	return lit.token.Position
}

// End implementation for Node.
func (lit *BasicLit) End() token.Position {
	return token.Position{
		Filename: lit.token.Position.Filename,
		Line:     lit.token.Position.Line,
		Column:   lit.token.Position.Column + len(lit.token.Content),
	}
}

func (lit *BasicLit) exprNode() {}

// ArrayLit is a node represents array literal.
type ArrayLit struct {
	Lbrack token.Position // position of "["
	Elts   []Expr         // list of array elements
	Rbrack token.Position // position of "]"
}

// Pos implementation for Node.
func (lit *ArrayLit) Pos() token.Position {
	return lit.Lbrack
}

// End implementation for Node.
func (lit *ArrayLit) End() token.Position {
	return token.Position{
		Filename: lit.Rbrack.Filename,
		Line:     lit.Rbrack.Line,
		Column:   lit.Rbrack.Column + 1,
	}
}

func (lit *ArrayLit) exprNode() {}

// KeyValueExpr is a node represents (key : value) pairs.
type KeyValueExpr struct {
	Key   Expr
	Colon token.Position // position of ":"
	Value Expr
}

// Pos implementation for Node.
func (kv *KeyValueExpr) Pos() token.Position {
	return kv.Key.Pos()
}

// End implementation for Node.
func (kv *KeyValueExpr) End() token.Position {
	return kv.Value.End()
}

func (kv *KeyValueExpr) exprNode() {}
