module token

pub struct Pos {
	filename string
	line     int
	column   int
}

pub fn (p Pos) str() string {
	return '$p.filename:$p.line,$p.column'
}
