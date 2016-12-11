package token

// Type is the set of lexical token types of Liza programming language.
type Type int

// List of token types.
const (
	// Special tokens
	ILLEGAL Type = iota
	EOF
	COMMENT

	literalBeg
	// Identifiers and basic type literals
	IDENT  // main
	INT    // 12345
	FLOAT  // 123.45
	STRING // "abc"
	literalEnd

	operatorBeg
	// Operators and delimiters
	ADD // +
	SUB // -
	MUL // *
	DIV // /
	REM // &

	AND // &
	OR  // |
	XOR // ^
	SHL // <<
	SHR // >>

	ADDASSIGN // +=
	SUBASSIGN // -=
	MULASSIGN // *=
	DIVASSIGN // /=
	REMASSIGN // %=

	ANDASSIGN // &=
	ORASSIGN  // |=
	XORASSIGN // ^=
	SHLASSIGN // <<=
	SHRASSIGN // >>=

	LAND // &&
	LOR  // ||
	INC  // ++
	DEC  // --

	EQL    // ==
	LSS    // <
	GTR    // >
	ASSIGN // =
	NOT    // !

	NEQ    // !=
	LEQ    // <=
	GEQ    // >=
	DEFINE // :=

	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,
	PERIOD // .

	RPAREN // )
	RBRACK // ]
	RBRACE // }
	operatorEnd

	keywordBeg
	// Keywords
	BREAK
	CLASS
	CONST
	CONTINUE
	ELSE

	ENUM
	FOR
	FUN
	IF
	IMPLEMENTS

	IMPORT
	INTERFACE
	MATCH
	PACKAGE
	PUBLIC

	RETURN
	keywordEnd
)
