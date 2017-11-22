package ast

import (
	"github.com/magic003/liza/token"
)

// Decl is the base type for all declaration tree nodes.
type Decl interface {
	Node
	declNode()
}

// ConstDecl node represents a constant declaration.
type ConstDecl struct {
	ConstPos token.Position // position of const
	Ident    token.Token    // constant identifier
	Type     Type           // constant type
	Value    Expr           // constant value
}

// Pos implementation for Node.
func (decl *ConstDecl) Pos() token.Position {
	return decl.ConstPos
}

// End implementation for Node.
func (decl *ConstDecl) End() token.Position {
	return decl.Value.End()
}

func (decl *ConstDecl) declNode() {}
