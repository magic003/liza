package parser

import (
	"github.com/magic003/liza/token"
)

// Error defines a parser error.
type Error struct {
	Pos token.Position
	Msg string
}

func (e Error) Error() string {
	return e.Pos.String() + ": " + e.Msg
}
