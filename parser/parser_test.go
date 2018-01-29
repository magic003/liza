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

var validTypeTestCases = []struct {
	desc     string
	src      []byte
	expected ast.Type
}{
	{
		desc: "array type",
		src:  []byte("[]int"),
		expected: &ast.ArrayType{
			Lbrack: token.Position{
				Filename: filename,
				Line:     1,
				Column:   1,
			},
			Rbrack: token.Position{
				Filename: filename,
				Line:     1,
				Column:   2,
			},
			Elt: &ast.BasicType{
				Ident: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   3,
					},
					Content: "int",
				},
			},
		},
	},
	{
		desc: "map type",
		src:  []byte("{string : int}"),
		expected: &ast.MapType{
			Lbrace: token.Position{
				Filename: filename,
				Line:     1,
				Column:   1,
			},
			Key: &ast.BasicType{
				Ident: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   2,
					},
					Content: "string",
				},
			},
			Value: &ast.BasicType{
				Ident: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   11,
					},
					Content: "int",
				},
			},
			Rbrace: token.Position{
				Filename: filename,
				Line:     1,
				Column:   14,
			},
		},
	},
	{
		desc: "tuple type",
		src:  []byte("(string, ast.Token)"),
		expected: &ast.TupleType{
			Lparen: token.Position{
				Filename: filename,
				Line:     1,
				Column:   1,
			},
			Elts: []ast.Type{
				&ast.BasicType{
					Ident: &token.Token{
						Type: token.IDENT,
						Position: token.Position{
							Filename: filename,
							Line:     1,
							Column:   2,
						},
						Content: "string",
					},
				},
				&ast.SelectorType{
					Package: &token.Token{
						Type: token.IDENT,
						Position: token.Position{
							Filename: filename,
							Line:     1,
							Column:   10,
						},
						Content: "ast",
					},
					Sel: &token.Token{
						Type: token.IDENT,
						Position: token.Position{
							Filename: filename,
							Line:     1,
							Column:   14,
						},
						Content: "Token",
					},
				},
			},
			Rparen: token.Position{
				Filename: filename,
				Line:     1,
				Column:   19,
			},
		},
	},
	{
		desc: "tuple type without empty element",
		src:  []byte("()"),
		expected: &ast.TupleType{
			Lparen: token.Position{
				Filename: filename,
				Line:     1,
				Column:   1,
			},
			Rparen: token.Position{
				Filename: filename,
				Line:     1,
				Column:   2,
			},
		},
	},
	{
		desc: "basic type",
		src:  []byte("string"),
		expected: &ast.BasicType{
			Ident: &token.Token{
				Type: token.IDENT,
				Position: token.Position{
					Filename: filename,
					Line:     1,
					Column:   1,
				},
				Content: "string",
			},
		},
	},
	{
		desc: "selector type",
		src:  []byte("ast.Token"),
		expected: &ast.SelectorType{
			Package: &token.Token{
				Type: token.IDENT,
				Position: token.Position{
					Filename: filename,
					Line:     1,
					Column:   1,
				},
				Content: "ast",
			},
			Sel: &token.Token{
				Type: token.IDENT,
				Position: token.Position{
					Filename: filename,
					Line:     1,
					Column:   5,
				},
				Content: "Token",
			},
		},
	},
}

func TestValidType(t *testing.T) {
	for _, tc := range validTypeTestCases {
		parser := New(filename, tc.src)
		result := parser.parseType()

		if !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("bad node for %s '%s':\ngot      %+v\nexpected %+v\n", tc.desc, tc.src, result, tc.expected)
		}
	}
}

var validInterfaceDeclTestCases = []struct {
	desc     string
	src      []byte
	expected *ast.InterfaceDecl
}{
	{
		desc: "interface declaration",
		src: []byte("interface Test {\n" +
			"fun foo(): {string : int}\n" +
			"fun bar(a int, b int): int\n" +
			"fun baz(str string)\n" +
			"}"),
		expected: &ast.InterfaceDecl{
			Interface: token.Position{
				Filename: filename,
				Line:     1,
				Column:   1,
			},
			Name: &token.Token{
				Type: token.IDENT,
				Position: token.Position{
					Filename: filename,
					Line:     1,
					Column:   11,
				},
				Content: "Test",
			},
			Lbrace: token.Position{
				Filename: filename,
				Line:     1,
				Column:   16,
			},
			Methods: []*ast.FuncDef{
				{
					Fun: token.Position{
						Filename: filename,
						Line:     2,
						Column:   1,
					},
					Name: &token.Token{
						Type: token.IDENT,
						Position: token.Position{
							Filename: filename,
							Line:     2,
							Column:   5,
						},
						Content: "foo",
					},
					Lparen: token.Position{
						Filename: filename,
						Line:     2,
						Column:   8,
					},
					Rparen: token.Position{
						Filename: filename,
						Line:     2,
						Column:   9,
					},
					ReturnType: &ast.MapType{
						Lbrace: token.Position{
							Filename: filename,
							Line:     2,
							Column:   12,
						},
						Key: &ast.BasicType{
							Ident: &token.Token{
								Type: token.IDENT,
								Position: token.Position{
									Filename: filename,
									Line:     2,
									Column:   13,
								},
								Content: "string",
							},
						},
						Value: &ast.BasicType{
							Ident: &token.Token{
								Type: token.IDENT,
								Position: token.Position{
									Filename: filename,
									Line:     2,
									Column:   22,
								},
								Content: "int",
							},
						},
						Rbrace: token.Position{
							Filename: filename,
							Line:     2,
							Column:   25,
						},
					},
				},
				{
					Fun: token.Position{
						Filename: filename,
						Line:     3,
						Column:   1,
					},
					Name: &token.Token{
						Type: token.IDENT,
						Position: token.Position{
							Filename: filename,
							Line:     3,
							Column:   5,
						},
						Content: "bar",
					},
					Lparen: token.Position{
						Filename: filename,
						Line:     3,
						Column:   8,
					},
					Params: []*ast.ParameterDef{
						{
							Name: &token.Token{
								Type: token.IDENT,
								Position: token.Position{
									Filename: filename,
									Line:     3,
									Column:   9,
								},
								Content: "a",
							},
							Type: &ast.BasicType{
								Ident: &token.Token{
									Type: token.IDENT,
									Position: token.Position{
										Filename: filename,
										Line:     3,
										Column:   11,
									},
									Content: "int",
								},
							},
						},
						{
							Name: &token.Token{
								Type: token.IDENT,
								Position: token.Position{
									Filename: filename,
									Line:     3,
									Column:   16,
								},
								Content: "b",
							},
							Type: &ast.BasicType{
								Ident: &token.Token{
									Type: token.IDENT,
									Position: token.Position{
										Filename: filename,
										Line:     3,
										Column:   18,
									},
									Content: "int",
								},
							},
						},
					},
					Rparen: token.Position{
						Filename: filename,
						Line:     3,
						Column:   21,
					},
					ReturnType: &ast.BasicType{
						Ident: &token.Token{
							Type: token.IDENT,
							Position: token.Position{
								Filename: filename,
								Line:     3,
								Column:   24,
							},
							Content: "int",
						},
					},
				},
				{
					Fun: token.Position{
						Filename: filename,
						Line:     4,
						Column:   1,
					},
					Name: &token.Token{
						Type: token.IDENT,
						Position: token.Position{
							Filename: filename,
							Line:     4,
							Column:   5,
						},
						Content: "baz",
					},
					Lparen: token.Position{
						Filename: filename,
						Line:     4,
						Column:   8,
					},
					Params: []*ast.ParameterDef{
						{
							Name: &token.Token{
								Type: token.IDENT,
								Position: token.Position{
									Filename: filename,
									Line:     4,
									Column:   9,
								},
								Content: "str",
							},
							Type: &ast.BasicType{
								Ident: &token.Token{
									Type: token.IDENT,
									Position: token.Position{
										Filename: filename,
										Line:     4,
										Column:   13,
									},
									Content: "string",
								},
							},
						},
					},
					Rparen: token.Position{
						Filename: filename,
						Line:     4,
						Column:   19,
					},
				},
			},
			Rbrace: token.Position{
				Filename: filename,
				Line:     5,
				Column:   1,
			},
		},
	},
	{
		desc: "interface declaration without methods",
		src:  []byte("interface Test {}"),
		expected: &ast.InterfaceDecl{
			Interface: token.Position{
				Filename: filename,
				Line:     1,
				Column:   1,
			},
			Name: &token.Token{
				Type: token.IDENT,
				Position: token.Position{
					Filename: filename,
					Line:     1,
					Column:   11,
				},
				Content: "Test",
			},
			Lbrace: token.Position{
				Filename: filename,
				Line:     1,
				Column:   16,
			},
			Rbrace: token.Position{
				Filename: filename,
				Line:     1,
				Column:   17,
			},
		},
	},
}

func TestValidInterfaceDecl(t *testing.T) {
	for _, tc := range validInterfaceDeclTestCases {
		parser := New(filename, tc.src)

		visibility := &token.Token{
			Type: token.PUBLIC,
			Position: token.Position{
				Filename: filename,
				Line:     1,
				Column:   1,
			},
			Content: "public",
		}
		result := parser.parseInterfaceDecl(visibility)
		tc.expected.Visibility = visibility

		if !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("bad node for %s '%s':\ngot      %+v\nexpected %+v\n", tc.desc, tc.src, result, tc.expected)
		}
	}
}

var validExprTestCases = []struct {
	desc     string
	src      []byte
	expected ast.Expr
}{
	{
		desc: "ident expression",
		src:  []byte("xyz"),
		expected: &ast.Ident{
			Token: &token.Token{
				Type: token.IDENT,
				Position: token.Position{
					Filename: filename,
					Line:     1,
					Column:   1,
				},
				Content: "xyz",
			},
		},
	},
	{
		desc: "int literal",
		src:  []byte("123"),
		expected: &ast.BasicLit{
			Token: &token.Token{
				Type: token.INT,
				Position: token.Position{
					Filename: filename,
					Line:     1,
					Column:   1,
				},
				Content: "123",
			},
		},
	},
	{
		desc: "float literal",
		src:  []byte("123.345"),
		expected: &ast.BasicLit{
			Token: &token.Token{
				Type: token.FLOAT,
				Position: token.Position{
					Filename: filename,
					Line:     1,
					Column:   1,
				},
				Content: "123.345",
			},
		},
	},
	{
		desc: "string literal",
		src:  []byte("\"abc\""),
		expected: &ast.BasicLit{
			Token: &token.Token{
				Type: token.STRING,
				Position: token.Position{
					Filename: filename,
					Line:     1,
					Column:   1,
				},
				Content: "\"abc\"",
			},
		},
	},
	{
		desc: "array literal",
		src:  []byte("[1, 2, 3]"),
		expected: &ast.ArrayLit{
			Lbrack: token.Position{
				Filename: filename,
				Line:     1,
				Column:   1,
			},
			Elts: []ast.Expr{
				&ast.BasicLit{
					Token: &token.Token{
						Type: token.INT,
						Position: token.Position{
							Filename: filename,
							Line:     1,
							Column:   2,
						},
						Content: "1",
					},
				},
				&ast.BasicLit{
					Token: &token.Token{
						Type: token.INT,
						Position: token.Position{
							Filename: filename,
							Line:     1,
							Column:   5,
						},
						Content: "2",
					},
				},
				&ast.BasicLit{
					Token: &token.Token{
						Type: token.INT,
						Position: token.Position{
							Filename: filename,
							Line:     1,
							Column:   8,
						},
						Content: "3",
					},
				},
			},
			Rbrack: token.Position{
				Filename: filename,
				Line:     1,
				Column:   9,
			},
		},
	},
	{
		desc: "empty array literal",
		src:  []byte("[]"),
		expected: &ast.ArrayLit{
			Lbrack: token.Position{
				Filename: filename,
				Line:     1,
				Column:   1,
			},
			Rbrack: token.Position{
				Filename: filename,
				Line:     1,
				Column:   2,
			},
		},
	},
	{
		desc: "map literal",
		src:  []byte("{1: 20, 2: 30}"),
		expected: &ast.MapLit{
			Lbrace: token.Position{
				Filename: filename,
				Line:     1,
				Column:   1,
			},
			Elts: []*ast.KeyValueExpr{
				{
					Key: &ast.BasicLit{
						Token: &token.Token{
							Type: token.INT,
							Position: token.Position{
								Filename: filename,
								Line:     1,
								Column:   2,
							},
							Content: "1",
						},
					},
					Colon: token.Position{
						Filename: filename,
						Line:     1,
						Column:   3,
					},
					Value: &ast.BasicLit{
						Token: &token.Token{
							Type: token.INT,
							Position: token.Position{
								Filename: filename,
								Line:     1,
								Column:   5,
							},
							Content: "20",
						},
					},
				},
				{
					Key: &ast.BasicLit{
						Token: &token.Token{
							Type: token.INT,
							Position: token.Position{
								Filename: filename,
								Line:     1,
								Column:   9,
							},
							Content: "2",
						},
					},
					Colon: token.Position{
						Filename: filename,
						Line:     1,
						Column:   10,
					},
					Value: &ast.BasicLit{
						Token: &token.Token{
							Type: token.INT,
							Position: token.Position{
								Filename: filename,
								Line:     1,
								Column:   12,
							},
							Content: "30",
						},
					},
				},
			},
			Rbrace: token.Position{
				Filename: filename,
				Line:     1,
				Column:   14,
			},
		},
	},
	{
		desc: "empty map literal",
		src:  []byte("{}"),
		expected: &ast.MapLit{
			Lbrace: token.Position{
				Filename: filename,
				Line:     1,
				Column:   1,
			},
			Rbrace: token.Position{
				Filename: filename,
				Line:     1,
				Column:   2,
			},
		},
	},
	{
		desc: "tuple literal",
		src:  []byte("(1, 2, 3)"),
		expected: &ast.TupleLit{
			Lparen: token.Position{
				Filename: filename,
				Line:     1,
				Column:   1,
			},
			Elts: []ast.Expr{
				&ast.BasicLit{
					Token: &token.Token{
						Type: token.INT,
						Position: token.Position{
							Filename: filename,
							Line:     1,
							Column:   2,
						},
						Content: "1",
					},
				},
				&ast.BasicLit{
					Token: &token.Token{
						Type: token.INT,
						Position: token.Position{
							Filename: filename,
							Line:     1,
							Column:   5,
						},
						Content: "2",
					},
				},
				&ast.BasicLit{
					Token: &token.Token{
						Type: token.INT,
						Position: token.Position{
							Filename: filename,
							Line:     1,
							Column:   8,
						},
						Content: "3",
					},
				},
			},
			Rparen: token.Position{
				Filename: filename,
				Line:     1,
				Column:   9,
			},
		},
	},
	{
		desc: "empty tuple literal",
		src:  []byte("()"),
		expected: &ast.TupleLit{
			Lparen: token.Position{
				Filename: filename,
				Line:     1,
				Column:   1,
			},
			Rparen: token.Position{
				Filename: filename,
				Line:     1,
				Column:   2,
			},
		},
	},
	{
		desc: "selector expression",
		src:  []byte("a.b"),
		expected: &ast.SelectorExpr{
			X: &ast.Ident{
				Token: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   1,
					},
					Content: "a",
				},
			},
			Sel: &ast.Ident{
				Token: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   3,
					},
					Content: "b",
				},
			},
		},
	},
	{
		desc: "selector expression with literal",
		src:  []byte("[1].size"),
		expected: &ast.SelectorExpr{
			X: &ast.ArrayLit{
				Lbrack: token.Position{
					Filename: filename,
					Line:     1,
					Column:   1,
				},
				Elts: []ast.Expr{
					&ast.BasicLit{
						Token: &token.Token{
							Type: token.INT,
							Position: token.Position{
								Filename: filename,
								Line:     1,
								Column:   2,
							},
							Content: "1",
						},
					},
				},
				Rbrack: token.Position{
					Filename: filename,
					Line:     1,
					Column:   3,
				},
			},
			Sel: &ast.Ident{
				Token: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   5,
					},
					Content: "size",
				},
			},
		},
	},
	{
		desc: "index expression",
		src:  []byte("a[0]"),
		expected: &ast.IndexExpr{
			X: &ast.Ident{
				Token: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   1,
					},
					Content: "a",
				},
			},
			Lbrack: token.Position{
				Filename: filename,
				Line:     1,
				Column:   2,
			},
			Index: &ast.BasicLit{
				Token: &token.Token{
					Type: token.INT,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   3,
					},
					Content: "0",
				},
			},
			Rbrack: token.Position{
				Filename: filename,
				Line:     1,
				Column:   4,
			},
		},
	},
	{
		desc: "index expression with array literal",
		src:  []byte("[1,2][0]"),
		expected: &ast.IndexExpr{
			X: &ast.ArrayLit{
				Lbrack: token.Position{
					Filename: filename,
					Line:     1,
					Column:   1,
				},
				Elts: []ast.Expr{
					&ast.BasicLit{
						Token: &token.Token{
							Type: token.INT,
							Position: token.Position{
								Filename: filename,
								Line:     1,
								Column:   2,
							},
							Content: "1",
						},
					},
					&ast.BasicLit{
						Token: &token.Token{
							Type: token.INT,
							Position: token.Position{
								Filename: filename,
								Line:     1,
								Column:   4,
							},
							Content: "2",
						},
					},
				},
				Rbrack: token.Position{
					Filename: filename,
					Line:     1,
					Column:   5,
				},
			},
			Lbrack: token.Position{
				Filename: filename,
				Line:     1,
				Column:   6,
			},
			Index: &ast.BasicLit{
				Token: &token.Token{
					Type: token.INT,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   7,
					},
					Content: "0",
				},
			},
			Rbrack: token.Position{
				Filename: filename,
				Line:     1,
				Column:   8,
			},
		},
	},
	{
		desc: "2 dimentional index expression",
		src:  []byte("a[0][b]"),
		expected: &ast.IndexExpr{
			X: &ast.IndexExpr{
				X: &ast.Ident{
					Token: &token.Token{
						Type: token.IDENT,
						Position: token.Position{
							Filename: filename,
							Line:     1,
							Column:   1,
						},
						Content: "a",
					},
				},
				Lbrack: token.Position{
					Filename: filename,
					Line:     1,
					Column:   2,
				},
				Index: &ast.BasicLit{
					Token: &token.Token{
						Type: token.INT,
						Position: token.Position{
							Filename: filename,
							Line:     1,
							Column:   3,
						},
						Content: "0",
					},
				},
				Rbrack: token.Position{
					Filename: filename,
					Line:     1,
					Column:   4,
				},
			},
			Lbrack: token.Position{
				Filename: filename,
				Line:     1,
				Column:   5,
			},
			Index: &ast.Ident{
				Token: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   6,
					},
					Content: "b",
				},
			},
			Rbrack: token.Position{
				Filename: filename,
				Line:     1,
				Column:   7,
			},
		},
	},
	{
		desc: "call expression",
		src:  []byte("add(1,2)"),
		expected: &ast.CallExpr{
			Fun: &ast.Ident{
				Token: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   1,
					},
					Content: "add",
				},
			},
			Lparen: token.Position{
				Filename: filename,
				Line:     1,
				Column:   4,
			},
			Args: []ast.Expr{
				&ast.BasicLit{
					Token: &token.Token{
						Type: token.INT,
						Position: token.Position{
							Filename: filename,
							Line:     1,
							Column:   5,
						},
						Content: "1",
					},
				},
				&ast.BasicLit{
					Token: &token.Token{
						Type: token.INT,
						Position: token.Position{
							Filename: filename,
							Line:     1,
							Column:   7,
						},
						Content: "2",
					},
				},
			},
			Rparen: token.Position{
				Filename: filename,
				Line:     1,
				Column:   8,
			},
		},
	},
	{
		desc: "call expression with selector expression",
		src:  []byte("x.add()"),
		expected: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X: &ast.Ident{
					Token: &token.Token{
						Type: token.IDENT,
						Position: token.Position{
							Filename: filename,
							Line:     1,
							Column:   1,
						},
						Content: "x",
					},
				},
				Sel: &ast.Ident{
					Token: &token.Token{
						Type: token.IDENT,
						Position: token.Position{
							Filename: filename,
							Line:     1,
							Column:   3,
						},
						Content: "add",
					},
				},
			},
			Lparen: token.Position{
				Filename: filename,
				Line:     1,
				Column:   6,
			},
			Rparen: token.Position{
				Filename: filename,
				Line:     1,
				Column:   7,
			},
		},
	},
	{
		desc: "unary expression",
		src:  []byte("!cond"),
		expected: &ast.UnaryExpr{
			Op: &token.Token{
				Type: token.NOT,
				Position: token.Position{
					Filename: filename,
					Line:     1,
					Column:   1,
				},
				Content: "!",
			},
			X: &ast.Ident{
				Token: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   2,
					},
					Content: "cond",
				},
			},
		},
	},
	{
		desc: "chained unary expression",
		src:  []byte("^-x"),
		expected: &ast.UnaryExpr{
			Op: &token.Token{
				Type: token.XOR,
				Position: token.Position{
					Filename: filename,
					Line:     1,
					Column:   1,
				},
				Content: "^",
			},
			X: &ast.UnaryExpr{
				Op: &token.Token{
					Type: token.SUB,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   2,
					},
					Content: "-",
				},
				X: &ast.Ident{
					Token: &token.Token{
						Type: token.IDENT,
						Position: token.Position{
							Filename: filename,
							Line:     1,
							Column:   3,
						},
						Content: "x",
					},
				},
			},
		},
	},
	{
		desc: "binary expression",
		src:  []byte("a + 1"),
		expected: &ast.BinaryExpr{
			X: &ast.Ident{
				Token: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   1,
					},
					Content: "a",
				},
			},
			Op: &token.Token{
				Type: token.ADD,
				Position: token.Position{
					Filename: filename,
					Line:     1,
					Column:   3,
				},
				Content: "+",
			},
			Y: &ast.BasicLit{
				Token: &token.Token{
					Type: token.INT,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   5,
					},
					Content: "1",
				},
			},
		},
	},
	{
		desc: "binary expression with same precedence",
		src:  []byte("a + 1 - b"),
		expected: &ast.BinaryExpr{
			X: &ast.BinaryExpr{
				X: &ast.Ident{
					Token: &token.Token{
						Type: token.IDENT,
						Position: token.Position{
							Filename: filename,
							Line:     1,
							Column:   1,
						},
						Content: "a",
					},
				},
				Op: &token.Token{
					Type: token.ADD,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   3,
					},
					Content: "+",
				},
				Y: &ast.BasicLit{
					Token: &token.Token{
						Type: token.INT,
						Position: token.Position{
							Filename: filename,
							Line:     1,
							Column:   5,
						},
						Content: "1",
					},
				},
			},
			Op: &token.Token{
				Type: token.SUB,
				Position: token.Position{
					Filename: filename,
					Line:     1,
					Column:   7,
				},
				Content: "-",
			},
			Y: &ast.Ident{
				Token: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   9,
					},
					Content: "b",
				},
			},
		},
	},
	{
		desc: "binary expression with different precedence",
		src:  []byte("a + 1 * b"),
		expected: &ast.BinaryExpr{
			X: &ast.Ident{
				Token: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   1,
					},
					Content: "a",
				},
			},
			Op: &token.Token{
				Type: token.ADD,
				Position: token.Position{
					Filename: filename,
					Line:     1,
					Column:   3,
				},
				Content: "+",
			},
			Y: &ast.BinaryExpr{
				X: &ast.BasicLit{
					Token: &token.Token{
						Type: token.INT,
						Position: token.Position{
							Filename: filename,
							Line:     1,
							Column:   5,
						},
						Content: "1",
					},
				},
				Op: &token.Token{
					Type: token.MUL,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   7,
					},
					Content: "*",
				},
				Y: &ast.Ident{
					Token: &token.Token{
						Type: token.IDENT,
						Position: token.Position{
							Filename: filename,
							Line:     1,
							Column:   9,
						},
						Content: "b",
					},
				},
			},
		},
	},
	{
		desc: "binary expression with parens",
		src:  []byte("(a + 1) * b"),
		expected: &ast.BinaryExpr{
			X: &ast.TupleLit{
				Lparen: token.Position{
					Filename: filename,
					Line:     1,
					Column:   1,
				},
				Elts: []ast.Expr{
					&ast.BinaryExpr{
						X: &ast.Ident{
							Token: &token.Token{
								Type: token.IDENT,
								Position: token.Position{
									Filename: filename,
									Line:     1,
									Column:   2,
								},
								Content: "a",
							},
						},
						Op: &token.Token{
							Type: token.ADD,
							Position: token.Position{
								Filename: filename,
								Line:     1,
								Column:   4,
							},
							Content: "+",
						},
						Y: &ast.BasicLit{
							Token: &token.Token{
								Type: token.INT,
								Position: token.Position{
									Filename: filename,
									Line:     1,
									Column:   6,
								},
								Content: "1",
							},
						},
					},
				},
				Rparen: token.Position{
					Filename: filename,
					Line:     1,
					Column:   7,
				},
			},
			Op: &token.Token{
				Type: token.MUL,
				Position: token.Position{
					Filename: filename,
					Line:     1,
					Column:   9,
				},
				Content: "*",
			},
			Y: &ast.Ident{
				Token: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   11,
					},
					Content: "b",
				},
			},
		},
	},
}

func TestValidExpr(t *testing.T) {
	for _, tc := range validExprTestCases {
		parser := New(filename, tc.src)
		result := parser.parseExpr()

		if !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("bad node for %s '%s':\ngot      %+v\nexpected %+v\n", tc.desc, tc.src, result, tc.expected)
		}
	}
}

var validConstDeclTestCases = []struct {
	desc     string
	src      []byte
	expected *ast.ConstDecl
}{
	{
		desc: "const declaration without type",
		src:  []byte("const x := 1"),
		expected: &ast.ConstDecl{
			Const: token.Position{
				Filename: filename,
				Line:     1,
				Column:   1,
			},
			Ident: &token.Token{
				Type: token.IDENT,
				Position: token.Position{
					Filename: filename,
					Line:     1,
					Column:   7,
				},
				Content: "x",
			},
			Value: &ast.BasicLit{
				Token: &token.Token{
					Type: token.INT,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   12,
					},
					Content: "1",
				},
			},
		},
	},
	{
		desc: "const declaration with type",
		src:  []byte("const x Int := 1"),
		expected: &ast.ConstDecl{
			Const: token.Position{
				Filename: filename,
				Line:     1,
				Column:   1,
			},
			Ident: &token.Token{
				Type: token.IDENT,
				Position: token.Position{
					Filename: filename,
					Line:     1,
					Column:   7,
				},
				Content: "x",
			},
			Type: &ast.BasicType{
				Ident: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   9,
					},
					Content: "Int",
				},
			},
			Value: &ast.BasicLit{
				Token: &token.Token{
					Type: token.INT,
					Position: token.Position{
						Filename: filename,
						Line:     1,
						Column:   16,
					},
					Content: "1",
				},
			},
		},
	},
}

func TestValidConstDecl(t *testing.T) {
	for _, tc := range validConstDeclTestCases {
		parser := New(filename, tc.src)

		visibility := &token.Token{
			Type: token.PUBLIC,
			Position: token.Position{
				Filename: filename,
				Line:     1,
				Column:   1,
			},
			Content: "public",
		}
		result := parser.parseConstDecl(visibility)
		tc.expected.Visibility = visibility

		if !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("bad node for %s '%s':\ngot      %+v\nexpected %+v\n", tc.desc, tc.src, result, tc.expected)
		}
	}
}
