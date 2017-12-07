package parser

import (
	"reflect"
	"testing"

	"github.com/magic003/liza/ast"
	"github.com/magic003/liza/token"
)

const filename = "test.lz"

var validPackageDeclTestCases = []struct {
	desc     string
	src      []byte
	expected ast.Node
}{
	{
		desc: "package declaration with newline",
		src:  []byte("package test\n"),
		expected: &ast.PackageDecl{
			Package: token.Position{
				Filename: filename,
				Line:     1,
				Column:   1,
			},
			Name: &token.Token{
				Type: token.IDENT,
				Position: token.Position{
					Filename: filename,
					Line:     1,
					Column:   9,
				},
				Content: "test",
			},
		},
	},
	{
		desc: "package declaration without newline",
		src:  []byte("package test"),
		expected: &ast.PackageDecl{
			Package: token.Position{
				Filename: filename,
				Line:     1,
				Column:   1,
			},
			Name: &token.Token{
				Type: token.IDENT,
				Position: token.Position{
					Filename: filename,
					Line:     1,
					Column:   9,
				},
				Content: "test",
			},
		},
	},
}

func TestValidPackageDecl(t *testing.T) {
	for _, tc := range validPackageDeclTestCases {
		parser := New(filename, tc.src)
		result := parser.parsePackageDecl()

		if !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("bad node for %s '%s':\ngot      %+v\nexpected %+v\n", tc.desc, tc.src, result, tc.expected)
		}
	}
}
