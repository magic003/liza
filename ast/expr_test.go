package ast

import (
	"testing"

	"github.com/magic003/liza/token"
)

var exprTestCases = []struct {
	desc        string
	expr        Expr
	expectedPos token.Position
	expectedEnd token.Position
}{
	{
		desc: "Ident",
		expr: &Ident{
			token: token.Token{
				Type: token.IDENT,
				Position: token.Position{
					Filename: "test.lz",
					Line:     10,
					Column:   26,
				},
				Content: "testVar",
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
			Column:   26 + 7,
		},
	},
	{
		desc: "BasicLit",
		expr: &BasicLit{
			token: token.Token{
				Type: token.INT,
				Position: token.Position{
					Filename: "test.lz",
					Line:     10,
					Column:   26,
				},
				Content: "12345",
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
			Column:   26 + 5,
		},
	},
	{
		desc: "ArrayLit",
		expr: &ArrayLit{
			Lbrack: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   26,
			},
			Rbrack: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   50,
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
			Column:   51,
		},
	},
	{
		desc: "KeyValueExpr",
		expr: &KeyValueExpr{
			Key: &Ident{
				token: token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: "test.lz",
						Line:     10,
						Column:   26,
					},
					Content: "testValue",
				},
			},
			Colon: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   36,
			},
			Value: &Ident{
				token: token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: "test.lz",
						Line:     11,
						Column:   10,
					},
					Content: "testKey",
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
			Line:     11,
			Column:   17,
		},
	},
	{
		desc: "MapLit",
		expr: &MapLit{
			Lbrace: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   26,
			},
			Rbrace: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   50,
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
			Column:   51,
		},
	},
	{
		desc: "TupleLit",
		expr: &TupleLit{
			Lparen: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   26,
			},
			Rparen: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   50,
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
			Column:   51,
		},
	},
	{
		desc: "ParenExpr",
		expr: &ParenExpr{
			Lparen: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   26,
			},
			Rparen: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   50,
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
			Column:   51,
		},
	},
	{
		desc: "SelectorExpr",
		expr: &SelectorExpr{
			X: &Ident{
				token: token.Token{
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
				token: token.Token{
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
		expectedPos: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   26,
		},
		expectedEnd: token.Position{
			Filename: "test.lz",
			Line:     10,
			Column:   35 + 10,
		},
	},
	{
		desc: "IndexExpr",
		expr: &IndexExpr{
			X: &Ident{
				token: token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: "test.lz",
						Line:     10,
						Column:   26,
					},
					Content: "testVar",
				},
			},
			Lbrack: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   33,
			},
			Index: &BasicLit{
				token: token.Token{
					Type: token.INT,
					Position: token.Position{
						Filename: "test.lz",
						Line:     10,
						Column:   34,
					},
					Content: "1",
				},
			},
			Rbrack: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   36,
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
			Column:   37,
		},
	},
	{
		desc: "CallExpr",
		expr: &CallExpr{
			Fun: &SelectorExpr{
				X: &Ident{
					token: token.Token{
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
					token: token.Token{
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
		desc: "UnaryExpr",
		expr: &UnaryExpr{
			Op: token.Token{
				Type: token.SUB,
				Position: token.Position{
					Filename: "test.lz",
					Line:     10,
					Column:   26,
				},
				Content: "-",
			},
			X: &Ident{
				token: token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: "test.lz",
						Line:     10,
						Column:   27,
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
			Column:   34,
		},
	},
	{
		desc: "BinaryExpr",
		expr: &BinaryExpr{
			X: &Ident{
				token: token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: "test.lz",
						Line:     10,
						Column:   26,
					},
					Content: "testVar",
				},
			},
			Op: token.Token{
				Type: token.ADD,
				Position: token.Position{
					Filename: "test.lz",
					Line:     10,
					Column:   33,
				},
				Content: "+",
			},
			Y: &Ident{
				token: token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: "test.lz",
						Line:     10,
						Column:   34,
					},
					Content: "y",
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
			Column:   35,
		},
	},
}

func TestExpr(t *testing.T) {
	for _, tc := range exprTestCases {
		results := []token.Position{tc.expr.Pos(), tc.expr.End()}
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
