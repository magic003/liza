package ast

import (
	"github.com/magic003/liza/token"
)

// File node represents a source file.
type File struct {
	Package *PackageDecl
	Imports []*ImportDecl
	Decls   []Decl // top level delarations
}

// Pos implementation for Node.
func (f *File) Pos() token.Position {
	return f.Package.Pos()
}

// End implementation for Node.
func (f *File) End() token.Position {
	if n := len(f.Decls); n > 0 {
		return f.Decls[n-1].End()
	}

	return f.Package.End()
}
