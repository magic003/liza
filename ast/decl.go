package ast

import (
	"github.com/magic003/liza/token"
)

// Decl is the base type for all declaration tree nodes.
type Decl interface {
	Node
	declNode()
}

// BadDecl node represents declarations containing syntax errors for which no declaration node can be created.
type BadDecl struct {
	From token.Position
	To   token.Position
}

// Pos implementation for Node.
func (decl *BadDecl) Pos() token.Position {
	return decl.From
}

// End implementation for Node.
func (decl *BadDecl) End() token.Position {
	return decl.To
}

func (decl *BadDecl) declNode() {}

// ConstDecl node represents a constant declaration.
type ConstDecl struct {
	Visibility *token.Token   // optional visibility
	Const      token.Position // position of const
	Ident      *token.Token   // constant identifier
	Type       Type           // optional constant type
	Value      Expr           // constant value
}

// Pos implementation for Node.
func (decl *ConstDecl) Pos() token.Position {
	if decl.Visibility != nil {
		return decl.Visibility.Position
	}
	return decl.Const
}

// End implementation for Node.
func (decl *ConstDecl) End() token.Position {
	return decl.Value.End()
}

func (decl *ConstDecl) declNode() {}

// VarDecl node represents a variable declaration.
type VarDecl struct {
	Visibility *token.Token   // optional visibility
	Var        token.Position // position of var
	Ident      *token.Token   // variable identifier
	Type       Type           // optional constant type
	Value      Expr           // variable initial value
}

// Pos implementation for Node.
func (decl *VarDecl) Pos() token.Position {
	if decl.Visibility != nil {
		return decl.Visibility.Position
	}
	return decl.Var
}

// End implementation for Node.
func (decl *VarDecl) End() token.Position {
	return decl.Value.End()
}

func (decl *VarDecl) declNode() {}

// PackageDecl node represents a package declaration.
type PackageDecl struct {
	Package token.Position // position of "package"
	Name    *token.Token   // token of package name
}

// Pos implementation for Node.
func (decl *PackageDecl) Pos() token.Position {
	return decl.Package
}

// End implementation for Node.
func (decl *PackageDecl) End() token.Position {
	return token.Position{
		Filename: decl.Name.Position.Filename,
		Line:     decl.Name.Position.Line,
		Column:   decl.Name.Position.Column + len(decl.Name.Content),
	}
}

func (decl *PackageDecl) declNode() {}

// ImportDecl node represents an import declaration.
type ImportDecl struct {
	Import token.Position // position of "import"
	Path   *ImportPath
	As     *token.Position // optional position of "as"
	Alias  *token.Token    // optional alias
}

// Pos implementation for Node.
func (decl *ImportDecl) Pos() token.Position {
	return decl.Import
}

// End implementation for Node.
func (decl *ImportDecl) End() token.Position {
	if decl.Alias != nil {
		return token.Position{
			Filename: decl.Alias.Position.Filename,
			Line:     decl.Alias.Position.Line,
			Column:   decl.Alias.Position.Column + len(decl.Alias.Content),
		}
	}

	return decl.Path.End()
}

func (decl *ImportDecl) declNode() {}

// ImportPath node represents an import path in an import declaration.
type ImportPath struct {
	LibraryName *token.Token   // external library name; nil if it's an internal path
	Path        []*token.Token // package name path
}

// Pos implementation for Node.
func (path *ImportPath) Pos() token.Position {
	if path.LibraryName != nil {
		return path.LibraryName.Position
	}

	return path.Path[0].Position // the path must have at least 1 package
}

// End implementation for Node.
func (path *ImportPath) End() token.Position {
	n := len(path.Path)
	lastPkg := path.Path[n-1]
	return token.Position{
		Filename: lastPkg.Position.Filename,
		Line:     lastPkg.Position.Line,
		Column:   lastPkg.Position.Column + len(lastPkg.Content),
	}
}

func (path *ImportPath) declNode() {}

// FuncDecl node represents a function declaration.
type FuncDecl struct {
	Visibility *token.Token    // optional visibility token
	Fun        token.Position  // position of fun
	Name       *token.Token    // function name
	Params     []*ParameterDef // parameters
	ReturnType Type            // return type; nil if it returns nothing
	Body       *BlockStmt      // function body
}

// Pos implementation for Node.
func (decl *FuncDecl) Pos() token.Position {
	if decl.Visibility != nil {
		return decl.Visibility.Position
	}

	return decl.Fun
}

// End implementation for Node.
func (decl *FuncDecl) End() token.Position {
	return decl.Body.End()
}

func (decl *FuncDecl) declNode() {}

// ParameterDef node represents a parameter definition.
type ParameterDef struct {
	Name *token.Token
	Type Type
}

// Pos implementation for Node.
func (param *ParameterDef) Pos() token.Position {
	return param.Name.Position
}

// End implementation for Node.
func (param *ParameterDef) End() token.Position {
	return param.Type.End()
}

func (param *ParameterDef) declNode() {}

// ClassDecl node represents a class declaration.
type ClassDecl struct {
	Visibility *token.Token   // optional visibility
	Class      token.Position // position of class
	Name       *token.Token
	Implements []*SelectorType
	Lbrace     token.Position
	Consts     []*ConstDecl
	Vars       []*VarDecl
	Methods    []*FuncDecl
	Rbrace     token.Position
}

// Pos implementation for Node.
func (decl *ClassDecl) Pos() token.Position {
	if decl.Visibility != nil {
		return decl.Visibility.Position
	}
	return decl.Class
}

// End implementation for Node.
func (decl *ClassDecl) End() token.Position {
	return token.Position{
		Filename: decl.Rbrace.Filename,
		Line:     decl.Rbrace.Line,
		Column:   decl.Rbrace.Column + 1,
	}
}

func (decl *ClassDecl) declNode() {}

// InterfaceDecl node represents an interface declaration.
type InterfaceDecl struct {
	Visibility *token.Token   // optional visibility
	Interface  token.Position // position of interface
	Name       *token.Token
	Lbrace     token.Position
	Methods    []*FuncDef
	Rbrace     token.Position
}

// Pos implementation for Node.
func (decl *InterfaceDecl) Pos() token.Position {
	if decl.Visibility != nil {
		return decl.Visibility.Position
	}
	return decl.Interface
}

// End implementation for Node.
func (decl *InterfaceDecl) End() token.Position {
	return token.Position{
		Filename: decl.Rbrace.Filename,
		Line:     decl.Rbrace.Line,
		Column:   decl.Rbrace.Column + 1,
	}
}

func (decl *InterfaceDecl) declNode() {}

// FuncDef node represents a function definition.
type FuncDef struct {
	Fun        token.Position // position of fun
	Name       *token.Token
	Lparen     token.Position
	Params     []*ParameterDef
	Rparen     token.Position
	ReturnType Type // return type; nil if it returns nothing
}

// Pos implementation for Node.
func (def *FuncDef) Pos() token.Position {
	return def.Fun
}

// End implementation for Node.
func (def *FuncDef) End() token.Position {
	if def.ReturnType != nil {
		return def.ReturnType.End()
	}
	return token.Position{
		Filename: def.Rparen.Filename,
		Line:     def.Rparen.Line,
		Column:   def.Rparen.Column + 1,
	}
}

func (def *FuncDef) declNode() {}
