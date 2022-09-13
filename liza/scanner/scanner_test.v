module scanner

fn test_from_file() ? {
	scanner := from_file('liza/scanner/testdata/dummy.lz')?
	assert 'liza/scanner/testdata/dummy.lz' == scanner.file_path
	assert 'dummy.lz' == scanner.file_name
	assert 'module dummy' == scanner.text.trim_space()
	assert 0 == scanner.pos
	assert 0 == scanner.line
	assert 0 == scanner.col
	assert 0 == scanner.errors.len
}

fn test_from_file_not_file() {
	if scanner := from_file('liza/scanner/testdata') {
		assert false, 'an error should be returned when the file_path is a dir'
	} else {
		assert 'liza/scanner/testdata is not a file' == err.msg()
	}
}

fn test_from_text() {
	scanner := from_text('dummy text')
	assert 'memory' == scanner.file_path
	assert 'memory' == scanner.file_name
	assert 'dummy text' == scanner.text
	assert 0 == scanner.pos
	assert 0 == scanner.line
	assert 0 == scanner.col
	assert 0 == scanner.errors.len
}
