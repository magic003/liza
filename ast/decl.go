package ast

// Decl is the base type for all declaration tree nodes.
type Decl interface {
	Node
	declNode()
}
