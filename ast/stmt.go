package ast

// Stmt is the base type for all statment tree nodes.
type Stmt interface {
	Node
	stmtNode()
}
