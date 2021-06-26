package ast

type Node interface {
	TokenLiteral() string // token字面量
	String() string
}

// Statement 声明
type Statement interface {
	Node
	statementNode() // 声明节点
}

// Expression 表达式
type Expression interface {
	Node
	expressionNode() // 表达式节点
}
