package lexer 

import "fmt"

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

func (r BasicToken) Value() (any, error) {
	if r.HasValue() {
		return r.TokenValue, nil
	} else {
		return nil, fmt.Errorf("Token of type %v cannot have value", r.TokenValue)
	}
}

func (r BasicToken) Type() TokenType {
	return r.TokenType
}

