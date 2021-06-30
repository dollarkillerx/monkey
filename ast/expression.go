package ast

import (
	"bytes"

	"github.com/dollarkillerx/monkey/token"
)

// Identifier 识别符
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {

}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

// IntegerLiteral 数字
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (i *IntegerLiteral) expressionNode() {

}

func (i *IntegerLiteral) TokenLiteral() string {
	return i.Token.Literal
}

func (i *IntegerLiteral) String() string {
	return i.Token.Literal
}

// PrefixExpression 前缀表达式
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (p *PrefixExpression) expressionNode() {}

func (p *PrefixExpression) TokenLiteral() string {
	return p.Token.Literal
}

func (p *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(p.Operator)
	out.WriteString(p.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (oe *InfixExpression) expressionNode() {}

func (oe *InfixExpression) TokenLiteral() string {
	return oe.Token.Literal
}

func (oe *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")
	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (bl *Boolean) expressionNode() {}
func (bl *Boolean) TokenLiteral() string {
	return bl.Token.Literal
}

func (bl *Boolean) String() string {
	return bl.Token.Literal
}
