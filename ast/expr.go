package ast

// Expr is the base type for all expression tree node.
type Expr interface {
	Node
	exprNode()
}
