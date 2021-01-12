// ast 抽象语法树代码, Monkey语言由表达式和语句组成
package ast

import (
	"gointerpreter/token"
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

// root node
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

// let语句实现的是语句，语句不生成值
type LetStatement struct {
	Token token.Token // 特指token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// 标识符实现的是表达式，表达式生成值，注意与语句的区别
type Identifier struct {
	Token token.Token // 特指token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
