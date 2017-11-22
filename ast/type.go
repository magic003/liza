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

// SelectorType node represents a type defined in a package.
type SelectorType struct {
	Package token.Token // package identifier
	Sel     token.Token //  type selector
}

// Pos implementation for Node.
func (selector *SelectorType) Pos() token.Position {
	return selector.Package.Position
}

// End implementation for Node.
func (selector *SelectorType) End() token.Position {
	return token.Position{
		Filename: selector.Sel.Position.Filename,
		Line:     selector.Sel.Position.Line,
		Column:   selector.Sel.Position.Column + len(selector.Sel.Content),
	}
}

func (selector *SelectorType) typeNode() {}

// ArrayType node represents an array type.
type ArrayType struct {
	Lbrack  token.Position // position of "["
	Rbrack  token.Position // position of "]"
	EltType Type           // element type
}

// Pos implementation for Node.
func (array *ArrayType) Pos() token.Position {
	return array.Lbrack
}

// End implementation for Node.
func (array *ArrayType) End() token.Position {
	return array.EltType.End()
}

func (array *ArrayType) typeNode() {}
