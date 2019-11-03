// token token.go
package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string // 字面值
}

const (
	// 非法的标识符和eof
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 标识符, 标识符比较特殊, 源码中可以出现任意字母组合的标识符
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

	// 关键字, 这里的值并不是源码中的关键字, 关键字字面值单独定义在keywords中
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// 单独维护一份关键字, 这里的关键字就是源码里面的关键字
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
