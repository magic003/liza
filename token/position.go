package token

// Position describes a location in the source code.
type Position struct {
	Filename string
	Line     int
	Column   int
}
