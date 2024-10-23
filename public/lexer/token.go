package lexer 


type TokenType int

const (
	INTEGER TokenType = iota
	MINUS
	PLUS
	MUL
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
	COLON
	COMMA
	FLOAT_DIV
	INTEGER_DIV
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


