// token token.go
package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// 非法的标识符和eof
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 标识符
	IDENT = "IDENT"
	INT   = "INT"

	// 运算符
	ASSIGN = "="
	PLUS   = "+"

	// 控制字符
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET4"
)
