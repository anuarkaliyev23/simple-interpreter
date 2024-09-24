package lexer 


type TokenType int

const (
	INTEGER TokenType = iota
	MINUS
	PLUS
	EOF
)

type BasicToken struct {
	TokenType TokenType
	TokenValue any
}

func (r BasicToken) HasValue() bool {
	return r.TokenType == INTEGER
}

