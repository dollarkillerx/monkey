package token

type TokenType string

type Token struct {
	Type    TokenType // token类型
	Literal string    // 字面量
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}

const (
	ILLEGAL = "ILLEGAL" // 非法
	EOF     = "EOF"     // 文件结束

	// 标识符
	IDENT = "IDENT" // 标识符  ADD, foobar, x, y
	INT   = "INT"

	// 操作符
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	EQ     = "=="
	NOT_EQ = "!="

	// 分隔符
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	// 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var (
	TokenEOF = Token{
		Type:    EOF,
		Literal: "",
	}
)

func NewToken(tokenType TokenType, ch byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
