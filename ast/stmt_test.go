package ast

import (
	"testing"

	"github.com/magic003/liza/token"
)

var stmtTestCases = []struct {
	desc        string
	stmt        Stmt
	expectedPos token.Position
	expectedEnd token.Position
}{
	{
		desc: "DeclStmt",
		stmt: &DeclStmt{
			Decl: &VarDecl{
				Var: token.Position{
					Filename: "test.lz",
					Line:     10,
					Column:   26,
				},
				Value: &BasicLit{
					Token: &token.Token{
						Type: token.INT,
						Position: token.Position{
							Filename: "test.lz",
							Line:     10,
							Column:   46,
						},
						Content: "12345",
					},
				},
			},
		},
		expectedPos: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   26,
		},
		expectedEnd: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   46 + 5,
		},
	},
	{
		desc: "ExprStmt",
		stmt: &ExprStmt{
			Expr: &CallExpr{
				Fun: &SelectorExpr{
					X: &Ident{
						Token: &token.Token{
							Type: token.IDENT,
							Position: token.Position{
								Filename: "test.lz",
								Line:     10,
								Column:   26,
							},
							Content: "testVar",
						},
					},
					Sel: &Ident{
						Token: &token.Token{
							Type: token.IDENT,
							Position: token.Position{
								Filename: "test.lz",
								Line:     10,
								Column:   35,
							},
							Content: "testMethod",
						},
					},
				},
				Lparen: token.Position{
					Filename: "test.lz",
					Line:     10,
					Column:   46,
				},
				Rparen: token.Position{
					Filename: "test.lz",
					Line:     10,
					Column:   47,
				},
			},
		},
		expectedPos: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   26,
		},
		expectedEnd: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   48,
		},
	},
	{
		desc: "IncDecStmt",
		stmt: &IncDecStmt{
			Expr: &Ident{
				Token: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: "test.lz",
						Line:     10,
						Column:   26,
					},
					Content: "testVar",
				},
			},
			Op: &token.Token{
				Type: token.INC,
				Position: token.Position{
					Filename: "test.lz",
					Line:     10,
					Column:   34,
				},
				Content: "++",
			},
		},
		expectedPos: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   26,
		},
		expectedEnd: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   36,
		},
	},
	{
		desc: "AssignStmt",
		stmt: &AssignStmt{
			LHS: &Ident{
				Token: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: "test.lz",
						Line:     10,
						Column:   26,
					},
					Content: "testVar",
				},
			},
			Assign: &token.Token{
				Type: token.ASSIGN,
				Position: token.Position{
					Filename: "test.lz",
					Line:     10,
					Column:   34,
				},
				Content: "=",
			},
			RHS: &Ident{
				Token: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: "test.lz",
						Line:     10,
						Column:   36,
					},
					Content: "testVar2",
				},
			},
		},
		expectedPos: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   26,
		},
		expectedEnd: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   44,
		},
	},
	{
		desc: "ReturnStmt",
		stmt: &ReturnStmt{
			Return: &token.Token{
				Type: token.RETURN,
				Position: token.Position{
					Filename: "test.lz",
					Line:     10,
					Column:   26,
				},
				Content: "return",
			},
			Value: &Ident{
				Token: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: "test.lz",
						Line:     10,
						Column:   33,
					},
					Content: "testVar",
				},
			},
		},
		expectedPos: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   26,
		},
		expectedEnd: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   40,
		},
	},
	{
		desc: "ReturnStmt",
		stmt: &ReturnStmt{
			Return: &token.Token{
				Type: token.RETURN,
				Position: token.Position{
					Filename: "test.lz",
					Line:     10,
					Column:   26,
				},
				Content: "return",
			},
		},
		expectedPos: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   26,
		},
		expectedEnd: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   32,
		},
	},
	{
		desc: "BranchStmt",
		stmt: &BranchStmt{
			Tok: &token.Token{
				Type: token.BREAK,
				Position: token.Position{
					Filename: "test.lz",
					Line:     10,
					Column:   26,
				},
				Content: "break",
			},
		},
		expectedPos: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   26,
		},
		expectedEnd: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   31,
		},
	},
	{
		desc: "BlockStmt",
		stmt: &BlockStmt{
			Lbrace: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   26,
			},
			Rbrace: token.Position{
				Filename: "test.lz",
				Line:     13,
				Column:   10,
			},
		},
		expectedPos: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   26,
		},
		expectedEnd: token.Position{
			Filename: "test.lz",
			Line:     13,
			Column:   11,
		},
	},
	{
		desc: "IfStmt",
		stmt: &IfStmt{
			If: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   26,
			},
			Body: &BlockStmt{
				Lbrace: token.Position{
					Filename: "test.lz",
					Line:     10,
					Column:   46,
				},
				Rbrace: token.Position{
					Filename: "test.lz",
					Line:     13,
					Column:   10,
				},
			},
		},
		expectedPos: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   26,
		},
		expectedEnd: token.Position{
			Filename: "test.lz",
			Line:     13,
			Column:   11,
		},
	},
	{
		desc: "IfStmt",
		stmt: &IfStmt{
			If: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   26,
			},
			Body: &BlockStmt{
				Lbrace: token.Position{
					Filename: "test.lz",
					Line:     10,
					Column:   46,
				},
				Rbrace: token.Position{
					Filename: "test.lz",
					Line:     13,
					Column:   10,
				},
			},
			Else: &ElseStmt{
				Else: token.Position{
					Filename: "test.lz",
					Line:     13,
					Column:   11,
				},
				Body: &BlockStmt{
					Lbrace: token.Position{
						Filename: "test.lz",
						Line:     13,
						Column:   13,
					},
					Rbrace: token.Position{
						Filename: "test.lz",
						Line:     18,
						Column:   10,
					},
				},
			},
		},
		expectedPos: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   26,
		},
		expectedEnd: token.Position{
			Filename: "test.lz",
			Line:     18,
			Column:   11,
		},
	},
	{
		desc: "ElseStmt",
		stmt: &ElseStmt{
			Else: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   26,
			},
			Body: &BlockStmt{
				Lbrace: token.Position{
					Filename: "test.lz",
					Line:     10,
					Column:   27,
				},
				Rbrace: token.Position{
					Filename: "test.lz",
					Line:     18,
					Column:   10,
				},
			},
		},
		expectedPos: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   26,
		},
		expectedEnd: token.Position{
			Filename: "test.lz",
			Line:     18,
			Column:   11,
		},
	},
	{
		desc: "ElseStmt",
		stmt: &ElseStmt{
			Else: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   26,
			},
			If: &IfStmt{
				If: token.Position{
					Filename: "test.lz",
					Line:     10,
					Column:   32,
				},
				Body: &BlockStmt{
					Lbrace: token.Position{
						Filename: "test.lz",
						Line:     10,
						Column:   46,
					},
					Rbrace: token.Position{
						Filename: "test.lz",
						Line:     13,
						Column:   10,
					},
				},
			},
		},
		expectedPos: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   26,
		},
		expectedEnd: token.Position{
			Filename: "test.lz",
			Line:     13,
			Column:   11,
		},
	},
	{
		desc: "MatchStmt",
		stmt: &MatchStmt{
			Match: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   26,
			},
			Rbrace: token.Position{
				Filename: "test.lz",
				Line:     20,
				Column:   26,
			},
		},
		expectedPos: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   26,
		},
		expectedEnd: token.Position{
			Filename: "test.lz",
			Line:     20,
			Column:   27,
		},
	},
	{
		desc: "CaseClause",
		stmt: &CaseClause{
			Pattern: &Ident{
				Token: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: "test.lz",
						Line:     10,
						Column:   26,
					},
					Content: "testVar",
				},
			},
			Colon: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   34,
			},
		},
		expectedPos: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   26,
		},
		expectedEnd: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   34,
		},
	},
	{
		desc: "CaseClause",
		stmt: &CaseClause{
			Pattern: &Ident{
				Token: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: "test.lz",
						Line:     10,
						Column:   26,
					},
					Content: "testVar",
				},
			},
			Colon: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   34,
			},
			Body: []Stmt{
				&ReturnStmt{
					Return: &token.Token{
						Type: token.RETURN,
						Position: token.Position{
							Filename: "test.lz",
							Line:     10,
							Column:   36,
						},
						Content: "return",
					},
					Value: &Ident{
						Token: &token.Token{
							Type: token.IDENT,
							Position: token.Position{
								Filename: "test.lz",
								Line:     10,
								Column:   43,
							},
							Content: "testVar",
						},
					},
				},
			},
		},
		expectedPos: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   26,
		},
		expectedEnd: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   50,
		},
	},
	{
		desc: "ForStmt",
		stmt: &ForStmt{
			For: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   26,
			},
			Body: &BlockStmt{
				Lbrace: token.Position{
					Filename: "test.lz",
					Line:     10,
					Column:   46,
				},
				Rbrace: token.Position{
					Filename: "test.lz",
					Line:     13,
					Column:   10,
				},
			},
		},
		expectedPos: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   26,
		},
		expectedEnd: token.Position{
			Filename: "test.lz",
			Line:     13,
			Column:   11,
		},
	},
}

func TestStmt(t *testing.T) {
	for _, tc := range stmtTestCases {
		results := []token.Position{tc.stmt.Pos(), tc.stmt.End()}
		expectedResults := []token.Position{tc.expectedPos, tc.expectedEnd}
		names := []string{"Pos()", "End()"}

		for i, result := range results {
			if result.Filename != expectedResults[i].Filename {
				t.Errorf("bad filename for %s %s: got %s, expected %s",
					tc.desc, names[i], result.Filename, expectedResults[i].Filename)
			}

			if result.Line != expectedResults[i].Line {
				t.Errorf("bad line for %s %s: got %v, expected %v",
					tc.desc, names[i], result.Line, expectedResults[i].Line)
			}

			if result.Column != expectedResults[i].Column {
				t.Errorf("bad column for %s %s: got %v, expected %v",
					tc.desc, names[i], result.Column, expectedResults[i].Column)
			}
		}
	}
}
