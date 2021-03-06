package ast

import (
	"testing"

	"github.com/magic003/liza/token"
)

var typeTestCases = []struct {
	desc        string
	typeNode    Type
	expectedPos token.Position
	expectedEnd token.Position
}{
	{
		desc: "BasicType",
		typeNode: &BasicType{
			Ident: &token.Token{
				Type: token.IDENT,
				Position: token.Position{
					Filename: "test.lz",
					Line:     10,
					Column:   26,
				},
				Content: "int",
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
			Column:   26 + 3,
		},
	},
	{
		desc: "SelectorType",
		typeNode: &SelectorType{
			Package: &token.Token{
				Type: token.IDENT,
				Position: token.Position{
					Filename: "test.lz",
					Line:     10,
					Column:   26,
				},
				Content: "testpackage",
			},
			Sel: &token.Token{
				Type: token.IDENT,
				Position: token.Position{
					Filename: "test.lz",
					Line:     10,
					Column:   38,
				},
				Content: "testtype",
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
			Column:   46,
		},
	},
	{
		desc: "ArrayType",
		typeNode: &ArrayType{
			Lbrack: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   26,
			},
			Elt: &BasicType{
				Ident: &token.Token{
					Type: token.IDENT,
					Position: token.Position{
						Filename: "test.lz",
						Line:     10,
						Column:   28,
					},
					Content: "int",
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
			Column:   31,
		},
	},
	{
		desc: "MapType",
		typeNode: &MapType{
			Lbrace: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   26,
			},
			Rbrace: token.Position{
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
		desc: "TupleType",
		typeNode: &TupleType{
			Lparen: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   26,
			},
			Rparen: token.Position{
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
}

func TestType(t *testing.T) {
	for _, tc := range typeTestCases {
		results := []token.Position{tc.typeNode.Pos(), tc.typeNode.End()}
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
