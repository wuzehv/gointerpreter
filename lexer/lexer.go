package lexer

import (
	"local/gointerpreter/token"
)

type Lexer struct {
	input        string
	position     int  // 当前字符的位置（ch）
	readPosition int  // 下一个字符的位置
	ch           byte // 仅支持ASCII，支持Unicode需要把类型改为rune
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// 读到字符串的结尾，ch设置为NUL
		l.ch = 0
	} else {
		// 字符串也可以像数组一样使用
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
