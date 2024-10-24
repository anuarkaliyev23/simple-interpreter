package lexer

var ReservedKeywords map[string]BasicToken = map[string]BasicToken{
	"PROGRAM":   {TokenType: PROGRAM},
	"VAR":   {TokenType: VAR},
	"DIV":   {TokenType: INTEGER_DIV},
	"INTEGER":   {TokenType: INTEGER_DECLARAION, TokenValue: "INTEGER" },
	"REAL":   {TokenType: REAL_DECLARATION, TokenValue: "REAL" },
	"BEGIN": {TokenType: BEGIN},
	"END":   {TokenType: END},
}
