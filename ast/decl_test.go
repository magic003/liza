package ast

import (
	"testing"

	"github.com/magic003/liza/token"
)

var declTestCases = []struct {
	desc        string
	decl        Decl
	expectedPos token.Position
	expectedEnd token.Position
}{
	{
		desc: "ConstDecl",
		decl: &ConstDecl{
			ConstPos: token.Position{
				Filename: "test.lz",
				Line:     10,
				Column:   26,
			},
			Value: &BasicLit{
				token: &token.Token{
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
		desc: "VarDecl",
		decl: &VarDecl{
			Ident: &token.Token{
				Type: token.IDENT,
				Position: token.Position{
					Filename: "test.lz",
					Line:     10,
					Column:   26,
				},
				Content: "testVar",
			},
			Value: &BasicLit{
				token: &token.Token{
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
		desc: "PackageDecl",
		decl: &PackageDecl{
			Package: token.Position{
				Filename: "test.lz",
				Line:     1,
				Column:   0,
			},
			Name: &token.Token{
				Type: token.IDENT,
				Position: token.Position{
					Filename: "test.lz",
					Line:     1,
					Column:   8,
				},
				Content: "hello",
			},
		},
		expectedPos: token.Position{
			Filename: "test.lz",
			Line:     1,
			Column:   0,
		},
		expectedEnd: token.Position{
			Filename: "test.lz",
			Line:     1,
			Column:   8 + 5,
		},
	},
	{
		desc: "ImportDecl with alias",
		decl: &ImportDecl{
			Import: token.Position{
				Filename: "test.lz",
				Line:     3,
				Column:   0,
			},
			Alias: &token.Token{
				Type: token.IDENT,
				Position: token.Position{
					Filename: "test.lz",
					Line:     3,
					Column:   25,
				},
				Content: "test",
			},
		},
		expectedPos: token.Position{
			Filename: "test.lz",
			Line:     3,
			Column:   0,
		},
		expectedEnd: token.Position{
			Filename: "test.lz",
			Line:     3,
			Column:   25 + 4,
		},
	},
	{
		desc: "ImportDecl without alias",
		decl: &ImportDecl{
			Import: token.Position{
				Filename: "test.lz",
				Line:     3,
				Column:   0,
			},
			Path: &ImportPath{
				Path: []*token.Token{
					{
						Type: token.IDENT,
						Position: token.Position{
							Filename: "test.lz",
							Line:     3,
							Column:   10,
						},
						Content: "test",
					},
				},
			},
		},
		expectedPos: token.Position{
			Filename: "test.lz",
			Line:     3,
			Column:   0,
		},
		expectedEnd: token.Position{
			Filename: "test.lz",
			Line:     3,
			Column:   10 + 4,
		},
	},
	{
		desc: "ImportPath with library name",
		decl: &ImportPath{
			LibraryName: &token.Token{
				Type: token.IDENT,
				Position: token.Position{
					Filename: "test.lz",
					Line:     3,
					Column:   3,
				},
				Content: "external",
			},
			Path: []*token.Token{
				{
					Type: token.IDENT,
					Position: token.Position{
						Filename: "test.lz",
						Line:     3,
						Column:   14,
					},
					Content: "test",
				},
				{
					Type: token.IDENT,
					Position: token.Position{
						Filename: "test.lz",
						Line:     3,
						Column:   19,
					},
					Content: "test2",
				},
			},
		},
		expectedPos: token.Position{
			Filename: "test.lz",
			Line:     3,
			Column:   3,
		},
		expectedEnd: token.Position{
			Filename: "test.lz",
			Line:     3,
			Column:   19 + 5,
		},
	},
	{
		desc: "ImportPath without library name",
		decl: &ImportPath{
			Path: []*token.Token{
				{
					Type: token.IDENT,
					Position: token.Position{
						Filename: "test.lz",
						Line:     3,
						Column:   14,
					},
					Content: "test",
				},
				{
					Type: token.IDENT,
					Position: token.Position{
						Filename: "test.lz",
						Line:     3,
						Column:   19,
					},
					Content: "test2",
				},
			},
		},
		expectedPos: token.Position{
			Filename: "test.lz",
			Line:     3,
			Column:   14,
		},
		expectedEnd: token.Position{
			Filename: "test.lz",
			Line:     3,
			Column:   19 + 5,
		},
	},
}

func TestDecl(t *testing.T) {
	for _, tc := range declTestCases {
		results := []token.Position{tc.decl.Pos(), tc.decl.End()}
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
