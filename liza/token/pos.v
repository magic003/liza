module token

// Pos is the position in a source file.
pub struct Pos {
	filename string
	line     int
	column   int
}

// str returns the Pos in a human readable text.
pub fn (p Pos) str() string {
	return '$p.filename:$p.line,$p.column'
}
