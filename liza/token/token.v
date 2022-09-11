module token

// Token is the definition of a lexical token.
pub struct Token {
pub:
	kind    Kind
	pos     Pos
	content string
}
