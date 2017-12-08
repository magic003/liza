package token

import (
	"strconv"
)

// Type is the set of lexical token types of Liza programming language.
type Type int

// List of token types.
const (
	// Special tokens
	ILLEGAL Type = iota
	EOF
	COMMENT
	NEWLINE

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
	REM // %

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

	RPAREN      // )
	RBRACK      // ]
	RBRACE      // }
	DOUBLECOLON // ::
	operatorEnd

	keywordBeg
	// Keywords
	AS
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

var types = [...]string{
	ILLEGAL: "ILLEGAL",

	EOF:     "EOF",
	COMMENT: "COMMENT",
	NEWLINE: "NEWLINE",

	IDENT:  "IDENT",
	INT:    "INT",
	FLOAT:  "FLOAT",
	STRING: "STRING",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	DIV: "/",
	REM: "%",

	AND: "&",
	OR:  "|",
	XOR: "^",
	SHL: "<<",
	SHR: ">>",

	ADDASSIGN: "+=",
	SUBASSIGN: "-=",
	MULASSIGN: "*=",
	DIVASSIGN: "/=",
	REMASSIGN: "%=",

	ANDASSIGN: "&=",
	ORASSIGN:  "|=",
	XORASSIGN: "^=",
	SHLASSIGN: "<<=",
	SHRASSIGN: ">>=",

	LAND: "&&",
	LOR:  "||",
	INC:  "++",
	DEC:  "--",

	EQL:    "==",
	LSS:    "<",
	GTR:    ">",
	ASSIGN: "=",
	NOT:    "!",

	NEQ:    "!=",
	LEQ:    "<=",
	GEQ:    ">=",
	DEFINE: ":=",

	LPAREN: "(",
	LBRACK: "[",
	LBRACE: "{",
	COMMA:  ",",
	PERIOD: ".",

	RPAREN:      ")",
	RBRACK:      "]",
	RBRACE:      "}",
	DOUBLECOLON: "::",

	AS:       "as",
	BREAK:    "break",
	CLASS:    "class",
	CONST:    "const",
	CONTINUE: "continue",

	ELSE: "else",
	ENUM: "enum",
	FOR:  "for",
	FUN:  "fun",
	IF:   "if",

	IMPLEMENTS: "implements",
	IMPORT:     "import",
	INTERFACE:  "interface",
	MATCH:      "match",
	PACKAGE:    "package",

	PUBLIC: "public",
	RETURN: "return",
}

func (ty Type) String() string {
	s := ""
	if 0 <= ty && ty < Type(len(types)) {
		s = types[ty]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(ty)) + ")"
	}
	return s
}

var keywords map[string]Type

func init() {
	keywords = make(map[string]Type)
	for i := keywordBeg + 1; i < keywordEnd; i++ {
		keywords[types[i]] = i
	}
}

// LookupKeyword looks if an identifier is a keyword. It returns keyword token or IDENT ( if not a keyword).
func LookupKeyword(ident string) Type {
	if ty, isKeyword := keywords[ident]; isKeyword {
		return ty
	}
	return IDENT
}
