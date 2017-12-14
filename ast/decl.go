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
	Ident    *token.Token   // constant identifier
	Type     Type           // optional constant type
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

// VarDecl node represents a variable declaration.
type VarDecl struct {
	Ident *token.Token // variable identifier
	Type  Type         // optional constant type
	Value Expr         // variable initial value
}

// Pos implementation for Node.
func (decl *VarDecl) Pos() token.Position {
	return decl.Ident.Position
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
