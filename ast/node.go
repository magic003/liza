package ast

import (
	"github.com/magic003/liza/token"
)

// Node is the base type for all abstract tree nodes.
type Node interface {
	// Pos returns the position of first character belonging to the node.
	Pos() token.Position
	// End returns the position of first character immedietely after the node.
	End() token.Position
}
