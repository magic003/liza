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

var validImportDeclTestCases = []struct {
	desc     string
	src      []byte
	expected ast.Node
}{
	{
		desc: "import declaration with alias and external library",
		src:  []byte("import external::foo/bar as baz"),
		expected: &ast.ImportDecl{
			Import: token.Position{
				Filename: filename,
				Line:     1,
				Column:   1,
			},
			Path: &ast.ImportPath{
				LibraryName: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   8,
					},
					Content: "external",
				},
				Path: []*token.Token{
					{
						Type: token.IDENT,
						Position: token.Position{
							Filename: filename,
							Line:     1,
							Column:   18,
						},
						Content: "foo",
					},
					{
						Type: token.IDENT,
						Position: token.Position{
							Filename: filename,
							Line:     1,
							Column:   22,
						},
						Content: "bar",
					},
				},
			},
			As: &token.Position{
				Filename: filename,
				Line:     1,
				Column:   26,
			},
			Alias: &token.Token{
				Type: token.IDENT,
				Position: token.Position{
					Filename: filename,
					Line:     1,
					Column:   29,
				},
				Content: "baz",
			},
		},
	},
	{
		desc: "import declaration without alias or external library",
		src:  []byte("import foo"),
		expected: &ast.ImportDecl{
			Import: token.Position{
				Filename: filename,
				Line:     1,
				Column:   1,
			},
			Path: &ast.ImportPath{
				Path: []*token.Token{
					{
						Type: token.IDENT,
						Position: token.Position{
							Filename: filename,
							Line:     1,
							Column:   8,
						},
						Content: "foo",
					},
				},
			},
		},
	},
}

func TestValidImportDecl(t *testing.T) {
	for _, tc := range validImportDeclTestCases {
		parser := New(filename, tc.src)
		result := parser.parseImportDecl()

		if !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("bad node for %s '%s':\ngot      %+v\nexpected %+v\n", tc.desc, tc.src, result, tc.expected)
		}
	}
}
