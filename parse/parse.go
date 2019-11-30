package parse

import (
	local/gointerpreter/token
	local/gointerpreter/lexer
	local/gointerpreter/ast
)

type Parse struct {
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parse {
	p := &Parse{l: l}

	// 调用两次是为了将curToken和peekToken都赋值，然后用来判断后续的操作（结束或者还需要其他信息）
	p.nextToken()
	p.nextToken()

	return p
}

func nextToken(p *Parse) {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parse) ParseProgram() *ast.Program {
	return nil
}
