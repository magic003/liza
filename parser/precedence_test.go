package parser

import (
	"testing"

	"github.com/magic003/liza/token"
)

var precedenceTestCases = []struct {
	desc     string
	op       token.Type
	expected int
}{
	{
		desc:     "||",
		op:       token.LOR,
		expected: 1,
	},
	{
		desc:     "&&",
		op:       token.LAND,
		expected: 2,
	},
	{
		desc:     "==",
		op:       token.EQL,
		expected: 3,
	},
	{
		desc:     "<",
		op:       token.LSS,
		expected: 3,
	},
	{
		desc:     ">",
		op:       token.GTR,
		expected: 3,
	},
	{
		desc:     "!=",
		op:       token.NEQ,
		expected: 3,
	},
	{
		desc:     "<=",
		op:       token.LEQ,
		expected: 3,
	},
	{
		desc:     ">=",
		op:       token.GEQ,
		expected: 3,
	},
	{
		desc:     "+",
		op:       token.ADD,
		expected: 4,
	},
	{
		desc:     "-",
		op:       token.SUB,
		expected: 4,
	},
	{
		desc:     "|",
		op:       token.OR,
		expected: 4,
	},
	{
		desc:     "^",
		op:       token.XOR,
		expected: 4,
	},
	{
		desc:     "*",
		op:       token.MUL,
		expected: 5,
	},
	{
		desc:     "/",
		op:       token.DIV,
		expected: 5,
	},
	{
		desc:     "%",
		op:       token.REM,
		expected: 5,
	},
	{
		desc:     "&",
		op:       token.AND,
		expected: 5,
	},
	{
		desc:     "<<",
		op:       token.SHL,
		expected: 5,
	},
	{
		desc:     ">>",
		op:       token.SHR,
		expected: 5,
	},
	{
		desc:     "++",
		op:       token.INC,
		expected: 0,
	},
}

func TestPrecedence(t *testing.T) {
	for _, tc := range precedenceTestCases {
		result := precedence(tc.op)
		if tc.expected != result {
			t.Errorf("bad precedence for '%s':\ngot      %+v\nexpected %+v\n", tc.desc, result, tc.expected)
		}
	}
}
