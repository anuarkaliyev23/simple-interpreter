package lexer

var ReservedKeywords map[string]BasicToken = map[string]BasicToken{
	"PROGRAM":   {TokenType: PROGRAM},
	"VAR":   {TokenType: VAR},
	"DIV":   {TokenType: DIV},
	"INTEGER":   {TokenType: INTEGER_DECLARAION},
	"REAL":   {TokenType: REAL_DECLARATION},
	"BEGIN": {TokenType: BEGIN},
	"END":   {TokenType: END},
}
