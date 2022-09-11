module token

// Kind is the token type.
pub enum Kind {
	// special tokens
	unknown
	eof
	newline
	comment
	// identifier and literals
	identifier // main
	number // 10, 10.0
	string // "s", 's', r"s", r's'
	char // `C` - rune
	// operators
	add // +
	minus // -
	mul // *
	div // /
	mod // %
	and // &
	pipe // |
	xor // ^
	left_shift // <<
	right_shift // >>
	unsigned_right_shift // >>>
	lt // <
	le // <=
	gt // >
	ge // >=
	eq // ==
	neq // !=
	logical_and // &&
	logical_or // ||
	assign // =
	add_assign // +=
	minus_assign // -=
	mul_assign // *=
	div_assign // /=
	mod_assign // %=
	and_assign // &=
	or_assign // |=
	xor_assign // ^=
	left_shift_assign // <<=
	right_shift_assign // >>=
	unsigned_right_shift_assign // >>>=
	decl_assign // :=
	colon // :
	double_colon // ::
	dot // .
	double_dot // ..
	comma // ,
	lparen // (
	rparen // )
	lbrack // [
	rbrack // ]
	lbrace // {
	rbrace // }
	not_in // !in
	// keywords
	keyword_beg
	key_module
	key_import
	key_as
	key_pub
	key_const
	key_fun
	key_void
	key_mut
	key_interface
	key_enum
	key_class
	key_override
	key_if
	key_else
	key_match
	key_break
	key_continue
	key_return
	key_for
	key_in
	key_defer
	key_new
	key_self
	keyword_end
	_end_
}

// str returns the token kind in a human readable text.
pub fn (k Kind) str() string {
	idx := int(k)
	if idx < 0 || idx >= token.token_str.len {
		return 'unknown'
	}
	return token.token_str[idx]
}

const token_str = build_token_str()

const keywords = build_keywords()

fn build_token_str() []string {
	mut s := []string{len: int(Kind._end_)}
	s[Kind.unknown] = 'unknown'
	s[Kind.eof] = 'eof'
	s[Kind.newline] = 'newline'
	s[Kind.comment] = 'comment'
	s[Kind.identifier] = 'identifier'
	s[Kind.number] = 'number'
	s[Kind.string] = 'string'
	s[Kind.char] = 'char'
	s[Kind.add] = '+'
	s[Kind.minus] = '-'
	s[Kind.mul] = '*'
	s[Kind.div] = '/'
	s[Kind.mod] = '%'
	s[Kind.and] = '&'
	s[Kind.pipe] = '|'
	s[Kind.xor] = 'xor'
	s[Kind.left_shift] = '<<'
	s[Kind.right_shift] = '>>'
	s[Kind.unsigned_right_shift] = '>>>'
	s[Kind.lt] = '<'
	s[Kind.le] = '<='
	s[Kind.gt] = '>'
	s[Kind.ge] = '>='
	s[Kind.eq] = '=='
	s[Kind.neq] = '!='
	s[Kind.logical_and] = '&&'
	s[Kind.logical_or] = '||'
	s[Kind.assign] = '='
	s[Kind.add_assign] = '+='
	s[Kind.minus_assign] = '-='
	s[Kind.mul_assign] = '*='
	s[Kind.div_assign] = '/='
	s[Kind.mod_assign] = '%='
	s[Kind.and_assign] = '&='
	s[Kind.or_assign] = '|='
	s[Kind.xor_assign] = '^='
	s[Kind.left_shift_assign] = '<<='
	s[Kind.right_shift_assign] = '>>='
	s[Kind.unsigned_right_shift_assign] = '>>>='
	s[Kind.decl_assign] = ':='
	s[Kind.colon] = ':'
	s[Kind.double_colon] = '::'
	s[Kind.dot] = '.'
	s[Kind.double_dot] = '..'
	s[Kind.comma] = ','
	s[Kind.lparen] = '('
	s[Kind.rparen] = ')'
	s[Kind.lbrack] = '['
	s[Kind.rbrack] = ']'
	s[Kind.lbrace] = '{'
	s[Kind.rbrace] = '}'
	s[Kind.not_in] = '!in'
	s[Kind.key_module] = 'module'
	s[Kind.key_import] = 'import'
	s[Kind.key_as] = 'as'
	s[Kind.key_pub] = 'pub'
	s[Kind.key_const] = 'const'
	s[Kind.key_fun] = 'fun'
	s[Kind.key_void] = 'void'
	s[Kind.key_mut] = 'mut'
	s[Kind.key_interface] = 'interface'
	s[Kind.key_enum] = 'enum'
	s[Kind.key_class] = 'class'
	s[Kind.key_override] = 'override'
	s[Kind.key_if] = 'if'
	s[Kind.key_else] = 'else'
	s[Kind.key_match] = 'match'
	s[Kind.key_break] = 'break'
	s[Kind.key_continue] = 'continue'
	s[Kind.key_return] = 'return'
	s[Kind.key_for] = 'for'
	s[Kind.key_in] = 'in'
	s[Kind.key_defer] = 'defer'
	s[Kind.key_new] = 'new'
	s[Kind.key_self] = 'self'

	// the following tokens are not returned by scanner.
	s[Kind.keyword_beg] = 'keyword_beg'
	s[Kind.keyword_end] = 'keyword_end'

	return s
}

fn build_keywords() map[string]Kind {
	mut res := map[string]Kind{}
	for t in int(Kind.keyword_beg) + 1 .. int(Kind.keyword_end) {
		key := token.token_str[t]
		res[key] = Kind(t)
	}
	return res
}
