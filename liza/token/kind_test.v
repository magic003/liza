module token

fn test_str() {
	// valid cases
	assert int(Kind._end_) == token_str.len
	assert 'eof' == Kind.eof.str()
	assert 'identifier' == Kind.identifier.str()
	assert '::' == Kind.double_colon.str()
	assert 'match' == Kind.key_match.str()

	// invalid cases
	assert 'unknown' == Kind(-1).str()
	assert 'unknown' == Kind._end_.str()
}

fn test_keywords() {
	assert (int(Kind.keyword_end) - int(Kind.keyword_beg) - 1) == keywords.len

	assert Kind.key_enum == keywords['enum']
	assert Kind.key_defer == keywords['defer']
	assert Kind.key_self == keywords['self']
}
