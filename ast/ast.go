// ast 抽象语法树代码, Monkey语言由表达式和语句组成
package ast

import (
	local/gointerpreter/token
)

// 语句和表达式都需要实现Node接口
type Node interface {
	TokenLiteral() string
}

// 语句接口
type Statement interface {
	Node
	statementNode()
}

// 表达式接口
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value String
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}