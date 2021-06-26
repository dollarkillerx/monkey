package leaxer

type TypeMatching func(ch byte) bool

// readIdentifier 获取整段
func (l *Lexer) readIdentifier(fn TypeMatching) string {
	start := l.position
	for fn(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

// isLetter 添加对 字母的支持
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' ||
		'A' <= ch && ch <= 'Z' ||
		ch == '_'
}

// isDigit 添加对数字的自支持
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
