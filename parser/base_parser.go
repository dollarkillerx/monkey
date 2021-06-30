package parser

import (
	"fmt"
	"strconv"

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

	stmt.Value = p.parseExpression(LOWEST)

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

	stmt.ReturnValue = p.parseExpression(LOWEST)

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseExpressionStatement 解析表达式
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	defer unTrace(trace("parseExpressionStatement"))

	stmt := &ast.ExpressionStatement{
		Token: p.curToken,
	}

	stmt.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseExpression 解析表达式
func (p *Parser) parseExpression(precedence int) ast.Expression {
	defer unTrace(trace("parseExpression"))

	prefix, ex := p.prefixParseFns[p.curToken.Type]
	if !ex {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}

	leftExp := prefix()
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix, ex := p.infixParseFns[p.peekToken.Type]
		if !ex {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

// parseIdentifier 解析识别符
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

// parseIntegerLiteral 解析int
func (p *Parser) parseIntegerLiteral() ast.Expression {
	defer unTrace(trace("parseIntegerLiteral"))

	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Errorf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value
	return lit
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Errorf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

// parsePrefixExpression 前置解析
func (p *Parser) parsePrefixExpression() ast.Expression {
	defer unTrace(trace("parsePrefixExpression"))

	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

// peekPrecedence 向下获取运算符优先级
func (p *Parser) peekPrecedence() int {
	if r, ok := precedences[p.peekToken.Type]; ok {
		return r
	}

	return LOWEST
}

// curPrecedence 向上获取运算符优先级
func (p *Parser) curPrecedence() int {
	if r, ok := precedences[p.curToken.Type]; ok {
		return r
	}

	return LOWEST
}

// parseInfixExpression 解析中缀运算符
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	defer unTrace(trace("parseInfixExpression"))

	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedences := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedences)
	return expression
}

// parseBoolean 解析bool
func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{
		Token: p.curToken,
		Value: p.curTokenIs(token.TRUE),
	}
}

// parseGroupedExpression 解析括号
func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)
	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}
