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

// MapLit is a node represents map literal.
type MapLit struct {
	Lbrace token.Position  // position of "{"
	Elts   []*KeyValueExpr // list of key value elements
	Rbrace token.Position  // position of "}"
}

// Pos implementation for Node.
func (lit *MapLit) Pos() token.Position {
	return lit.Lbrace
}

// End implementation for Node.
func (lit *MapLit) End() token.Position {
	return token.Position{
		Filename: lit.Rbrace.Filename,
		Line:     lit.Rbrace.Line,
		Column:   lit.Rbrace.Column + 1,
	}
}

func (lit *MapLit) exprNode() {}

// TupleLit is a node represents tuple literal.
type TupleLit struct {
	Lparen token.Position // position of "("
	Elts   []Expr         // list of elements
	Rparen token.Position // positionof ")"
}

// Pos implementation for Node.
func (lit *TupleLit) Pos() token.Position {
	return lit.Lparen
}

// End implementation for Node.
func (lit *TupleLit) End() token.Position {
	return token.Position{
		Filename: lit.Rparen.Filename,
		Line:     lit.Rparen.Line,
		Column:   lit.Rparen.Column + 1,
	}
}

func (lit *TupleLit) exprNode() {}

// ParenExpr is a node represents a parenthesized expression.
type ParenExpr struct {
	Lparen token.Position // position of "("
	X      Expr           // parenthesized expression
	Rparen token.Position // position of ")"
}

// Pos implementation for Node.
func (expr *ParenExpr) Pos() token.Position {
	return expr.Lparen
}

// End implementation for Node.
func (expr *ParenExpr) End() token.Position {
	return token.Position{
		Filename: expr.Rparen.Filename,
		Line:     expr.Rparen.Line,
		Column:   expr.Rparen.Column + 1,
	}
}

func (expr *ParenExpr) exprNode() {}

// SelectorExpr is a node represents a selector expression.
type SelectorExpr struct {
	X   Expr   // expression
	Sel *Ident // selector
}

// Pos implementation for Node.
func (expr *SelectorExpr) Pos() token.Position {
	return expr.X.Pos()
}

// End implementation for Node.
func (expr *SelectorExpr) End() token.Position {
	return expr.Sel.End()
}

func (expr *SelectorExpr) exprNode() {}

// IndexExpr is a node represents an expression followed by an index.
type IndexExpr struct {
	X      Expr           // expression
	Lbrack token.Position // position of "["
	Index  Expr           // index expression
	Rbrack token.Position // position of "]"
}

// Pos implementation for Node.
func (index *IndexExpr) Pos() token.Position {
	return index.X.Pos()
}

// End implementation for Node.
func (index *IndexExpr) End() token.Position {
	return token.Position{
		Filename: index.Rbrack.Filename,
		Line:     index.Rbrack.Line,
		Column:   index.Rbrack.Column + 1,
	}
}

func (index *IndexExpr) exprNode() {}

// CallExpr is a node represents a function call.
type CallExpr struct {
	Fun    Expr           // function expression
	Lparen token.Position // position of "("
	Args   []Expr         // function arguments
	Rparen token.Position // position of ")"
}

// Pos implementation for Node.
func (call *CallExpr) Pos() token.Position {
	return call.Fun.Pos()
}

// End implementation for Node.
func (call *CallExpr) End() token.Position {
	return token.Position{
		Filename: call.Rparen.Filename,
		Line:     call.Rparen.Line,
		Column:   call.Rparen.Column + 1,
	}
}

func (call *CallExpr) exprNode() {}

// UnaryExpr node represents an unary expression.
type UnaryExpr struct {
	Op token.Token // operator
	X  Expr        // operand
}

// Pos implementation for Node.
func (unary *UnaryExpr) Pos() token.Position {
	return unary.Op.Position
}

// End implementation for Node.
func (unary *UnaryExpr) End() token.Position {
	return unary.X.End()
}

func (unary *UnaryExpr) exprNode() {}

// BinaryExpr node represents a binary expression.
type BinaryExpr struct {
	X  Expr        // left operand
	Op token.Token // operator
	Y  Expr        // right operand
}

// Pos implementation for Node.
func (binary *BinaryExpr) Pos() token.Position {
	return binary.X.Pos()
}

// End implementation for Node.
func (binary *BinaryExpr) End() token.Position {
	return binary.Y.End()
}

func (binary *BinaryExpr) exprNode() {}
