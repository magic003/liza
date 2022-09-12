module scanner

import os
import v.errors

pub struct Scanner {
	file_path string // file path: '/path/to/file.lz'
	file_name string // file name: 'file.lz'
	text      string // the text of the entire source file
mut:
	idx int // current position in the text
pub mut:
	errors []errors.Error // errors found by the scanner
}

// from_file creates a scanner from a source file.
pub fn from_file(file_path string) ?&Scanner {
	if !os.is_file(file_path) {
		return error('$file_path is not a file')
	}

	text := os.read_file(file_path)?
	return &Scanner{
		file_path: file_path
		file_name: os.base(file_path)
		text: text
		idx: 0
	}
}

// from_text creates a scanner from text.
pub fn from_text(text string) &Scanner {
	return &Scanner{
		file_path: 'memory'
		file_name: 'memory'
		text: text
		idx: 0
	}
}
