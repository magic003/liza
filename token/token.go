package token

// Token is the definition of a lexical token.
type Token struct {
	Type     Type
	Position Position
	Content  string
}
