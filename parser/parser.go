package parser

import (
	"fmt"

	"github.com/dollarkillerx/monkey/ast"
	"github.com/dollarkillerx/monkey/lexer"
	"github.com/dollarkillerx/monkey/token"
)

type Parser struct {
	l      *lexer.Lexer
	errors []error

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []error{},
	}

	// 填充 curToken && peekToken
	p.nextToken()
	p.nextToken()

	return p
}

// nextToken 读取下一个token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram 解析代码
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

// parseStatement 解析具体某个
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

// 判断当前节点token type
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// expectPeek 期待的token类型
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken() // 如果是就进步
		return true
	} else {
		p.peekError(t)
	}
	return false
}

// peekTokenIs 判定下一个token的类型
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// Errors 获取错误信息
func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Errorf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

type (
	// 前缀解析
	prefixParseFn func() ast.ExpressionStatement
	// 中缀解析
	infixParseFn func(expression ast.Expression) ast.ExpressionStatement
)

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}
