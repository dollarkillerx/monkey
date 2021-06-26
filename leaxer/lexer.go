package leaxer

import (
	"github.com/dollarkillerx/monkey/token"
)

type Lexer struct {
	input        string
	len          int
	position     int  // 当前光标位置
	readPosition int  // 光标位置的下一个字符 current reading position in input (after current char)
	ch           byte // 当前读取字符
}

func New(input string) *Lexer {
	l := &Lexer{input: input, len: len(input)}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= l.len {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace() // 跳过空白符

	switch l.ch {
	case '=':
		tok = token.NewToken(token.ASSIGN, l.ch)
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.ch)
	case ',':
		tok = token.NewToken(token.COMMA, l.ch)
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)
	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)
	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)
	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
	case 0:
		tok = token.TokenEOF
	default:
		switch {
		case isLetter(l.ch):
			tok.Literal = l.readIdentifier(isLetter)
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		case isDigit(l.ch):
			tok.Literal = l.readIdentifier(isDigit)
			tok.Type = token.INT
			return tok
		default:
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}
