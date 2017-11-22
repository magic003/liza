package ast

import (
	"github.com/magic003/liza/token"
)

// Type is the base type for all type tree node.
type Type interface {
	Node
	typeNode()
}

// BasicType node represents a basic type provided by the language.
type BasicType struct {
	Ident token.Token // identifier for a basic type
}

// Pos implementation for Node.
func (basic *BasicType) Pos() token.Position {
	return basic.Ident.Position
}

// End implementation for Node.
func (basic *BasicType) End() token.Position {
	return token.Position{
		Filename: basic.Ident.Position.Filename,
		Line:     basic.Ident.Position.Line,
		Column:   basic.Ident.Position.Column + len(basic.Ident.Content),
	}
}

func (basic *BasicType) typeNode() {}
