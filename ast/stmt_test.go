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
				Ident: token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: "test.lz",
						Line:     10,
						Column:   26,
					},
					Content: "testVar",
				},
				Value: &BasicLit{
					token: token.Token{
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
