package lexer 


type TokenType int

const (
	INTEGER TokenType = iota
	MINUS
	PLUS
	MUL
	DIV
	LPAREN
	RPAREN
	BEGIN
	END
	DOT
	ASSIGN
	SEMICOLON
	ID
	EOF
	PROGRAM
	VAR
	REAL
	REAL_DECLARATION
	INTEGER_DECLARAION
)

type BasicToken struct {
	TokenType TokenType
	TokenValue string
}

func (r BasicToken) HasValue() bool {
	return r.TokenType == INTEGER ||
		r.TokenType == ID ||
		r.TokenType == REAL
}


