module token

fn test_str() {
	p := Pos{
		filename: 'test.v'
		line: 10
		column: 123
	}

	assert 'test.v:10,123' == p.str()
}
