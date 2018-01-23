package parser

import (
	"github.com/magic003/liza/token"
)

const lowestPrec = 0 // non-operators

func precedence(op token.Type) int {
	switch op {
	case token.LOR:
		return 1
	case token.LAND:
		return 2
	case token.EQL, token.LSS, token.GTR, token.NEQ, token.LEQ, token.GEQ:
		return 3
	case token.ADD, token.SUB, token.OR, token.XOR:
		return 4
	case token.MUL, token.DIV, token.REM, token.AND, token.SHL, token.SHR:
		return 5
	}
	return lowestPrec
}
