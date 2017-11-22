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
			Ident: token.Token{
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
