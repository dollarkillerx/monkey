package parser

import (
	"github.com/dollarkillerx/monkey/ast"
	"github.com/dollarkillerx/monkey/token"
)

// parseLetStatement let 解析
// let a = 16;
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{
		Token: p.curToken,
	}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// 设定变量名称
	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	// 如果不是 `=` 就跳过
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// 我们跳过表达式，直到我们 遇到一个分号为止
	if !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseReturnStatement return 解析
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
