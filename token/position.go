package token

import (
	"fmt"
)

// Position describes a location in the source code.
type Position struct {
	Filename string
	Line     int
	Column   int
}

// String returns the string format of a position.
func (pos Position) String() string {
	return fmt.Sprintf("%s:%d:%d", pos.Filename, pos.Line, pos.Column)
}
