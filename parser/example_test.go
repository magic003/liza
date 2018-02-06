package parser

import (
	"reflect"
	"testing"

	"github.com/magic003/liza/ast"
	"github.com/magic003/liza/token"
)

func TestExample(t *testing.T) {
	exampleFile := "example.lz"
	src := `package example

import io
import lib::simple/beautydate as date

public class Example {
	public fun main(args []string) {
		io.print(date.now())
	}
}`

	parser := New(exampleFile, []byte(src))
	result := parser.parseFile()
	expected := &ast.File{
		Package: &ast.PackageDecl{
			Package: token.Position{
				Filename: exampleFile,
				Line:     1,
				Column:   1,
			},
			Name: &token.Token{
				Type: token.IDENT,
				Position: token.Position{
					Filename: exampleFile,
					Line:     1,
					Column:   9,
				},
				Content: "example",
			},
		},
		Imports: []*ast.ImportDecl{
			{
				Import: token.Position{
					Filename: exampleFile,
					Line:     3,
					Column:   1,
				},
				Path: &ast.ImportPath{
					Path: []*token.Token{
						{
							Type: token.IDENT,
							Position: token.Position{
								Filename: exampleFile,
								Line:     3,
								Column:   8,
							},
							Content: "io",
						},
					},
				},
			},
			{
				Import: token.Position{
					Filename: exampleFile,
					Line:     4,
					Column:   1,
				},
				Path: &ast.ImportPath{
					LibraryName: &token.Token{
						Type: token.IDENT,
						Position: token.Position{
							Filename: exampleFile,
							Line:     4,
							Column:   8,
						},
						Content: "lib",
					},
					Path: []*token.Token{
						{
							Type: token.IDENT,
							Position: token.Position{
								Filename: exampleFile,
								Line:     4,
								Column:   13,
							},
							Content: "simple",
						},
						{
							Type: token.IDENT,
							Position: token.Position{
								Filename: exampleFile,
								Line:     4,
								Column:   20,
							},
							Content: "beautydate",
						},
					},
				},
				As: &token.Position{
					Filename: exampleFile,
					Line:     4,
					Column:   31,
				},
				Alias: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: exampleFile,
						Line:     4,
						Column:   34,
					},
					Content: "date",
				},
			},
		},
		Decls: []ast.Decl{
			&ast.ClassDecl{
				Visibility: &token.Token{
					Type: token.PUBLIC,
					Position: token.Position{
						Filename: exampleFile,
						Line:     6,
						Column:   1,
					},
					Content: "public",
				},
				Class: token.Position{
					Filename: exampleFile,
					Line:     6,
					Column:   8,
				},
				Name: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: exampleFile,
						Line:     6,
						Column:   14,
					},
					Content: "Example",
				},
				Lbrace: token.Position{
					Filename: exampleFile,
					Line:     6,
					Column:   22,
				},
				Methods: []*ast.FuncDecl{
					{
						Visibility: &token.Token{
							Type: token.PUBLIC,
							Position: token.Position{
								Filename: exampleFile,
								Line:     7,
								Column:   2,
							},
							Content: "public",
						},
						Fun: token.Position{
							Filename: exampleFile,
							Line:     7,
							Column:   9,
						},
						Name: &token.Token{
							Type: token.IDENT,
							Position: token.Position{
								Filename: exampleFile,
								Line:     7,
								Column:   13,
							},
							Content: "main",
						},
						Params: []*ast.ParameterDef{
							{
								Name: &token.Token{
									Type: token.IDENT,
									Position: token.Position{
										Filename: exampleFile,
										Line:     7,
										Column:   18,
									},
									Content: "args",
								},
								Type: &ast.ArrayType{
									Lbrack: token.Position{
										Filename: exampleFile,
										Line:     7,
										Column:   23,
									},
									Rbrack: token.Position{
										Filename: exampleFile,
										Line:     7,
										Column:   24,
									},
									Elt: &ast.BasicType{
										Ident: &token.Token{
											Type: token.IDENT,
											Position: token.Position{
												Filename: exampleFile,
												Line:     7,
												Column:   25,
											},
											Content: "string",
										},
									},
								},
							},
						},
						Body: &ast.BlockStmt{
							Lbrace: token.Position{
								Filename: exampleFile,
								Line:     7,
								Column:   33,
							},
							Stmts: []ast.Stmt{
								&ast.ExprStmt{
									Expr: &ast.CallExpr{
										Fun: &ast.SelectorExpr{
											X: &ast.Ident{
												Token: &token.Token{
													Type: token.IDENT,
													Position: token.Position{
														Filename: exampleFile,
														Line:     8,
														Column:   3,
													},
													Content: "io",
												},
											},
											Sel: &ast.Ident{
												Token: &token.Token{
													Type: token.IDENT,
													Position: token.Position{
														Filename: exampleFile,
														Line:     8,
														Column:   6,
													},
													Content: "print",
												},
											},
										},
										Lparen: token.Position{
											Filename: exampleFile,
											Line:     8,
											Column:   11,
										},
										Args: []ast.Expr{
											&ast.CallExpr{
												Fun: &ast.SelectorExpr{
													X: &ast.Ident{
														Token: &token.Token{
															Type: token.IDENT,
															Position: token.Position{
																Filename: exampleFile,
																Line:     8,
																Column:   12,
															},
															Content: "date",
														},
													},
													Sel: &ast.Ident{
														Token: &token.Token{
															Type: token.IDENT,
															Position: token.Position{
																Filename: exampleFile,
																Line:     8,
																Column:   17,
															},
															Content: "now",
														},
													},
												},
												Lparen: token.Position{
													Filename: exampleFile,
													Line:     8,
													Column:   20,
												},
												Rparen: token.Position{
													Filename: exampleFile,
													Line:     8,
													Column:   21,
												},
											},
										},
										Rparen: token.Position{
											Filename: exampleFile,
											Line:     8,
											Column:   22,
										},
									},
								},
							},
							Rbrace: token.Position{
								Filename: exampleFile,
								Line:     9,
								Column:   2,
							},
						},
					},
				},
				Rbrace: token.Position{
					Filename: exampleFile,
					Line:     10,
					Column:   1,
				},
			},
		},
	}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("bad node for example.lz '%s':\ngot      %+v\nexpected %+v\n", src, result, expected)
	}
}
