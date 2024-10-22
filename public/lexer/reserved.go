package lexer

var ReservedKeywords map[string]BasicToken = map[string]BasicToken{
	"BEGIN": {TokenType: BEGIN},
	"END":   {TokenType: END},
}
