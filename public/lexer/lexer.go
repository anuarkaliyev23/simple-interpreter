package lexer

import (
	"fmt"
	"unicode"
)

type BasicLexer struct {
	Text         string
	Position     int
	CurrentToken *BasicToken
	IsReachedEOF bool
}

func (r *BasicLexer) currentChar() *byte {
	char := r.Text[r.Position]
	return &char
}

func (r *BasicLexer) advance() {
	if !r.IsReachedEOF {
		r.Position++
		if r.Position >= len(r.Text) {
			r.IsReachedEOF = true
		}
	}
}

func (r *BasicLexer) Initialize() error {
	token, err := r.NextToken()
	if err != nil {
		return err
	}
	r.CurrentToken = &token
	return nil
}

func (r *BasicLexer) currentRune() rune {
	return rune(*r.currentChar())
}


func (r *BasicLexer) peekRune() rune {
	return rune(*r.peek())
}

func (r *BasicLexer) isOnSpace() bool {
	return unicode.IsSpace(r.currentRune())
}

func (r *BasicLexer) isOnDigit() bool {
	return unicode.IsDigit(r.currentRune())
}

func (r *BasicLexer) parseNumber() BasicToken {
	result := ""
	for !r.IsReachedEOF && r.isOnDigit() {
		result += string(*r.currentChar())
		r.advance()
	}

	if r.currentRune() == '.' {
		result += string(*r.currentChar())
		r.advance()

		for r.currentChar() != nil || unicode.IsDigit(r.currentRune()) {
			result += string(*r.currentChar())
			r.advance()
		}
		
		return BasicToken{
			TokenType: REAL,
			TokenValue: result,
		}
	}

	return BasicToken{
		TokenType: INTEGER,
		TokenValue: result,
	}
}

func (r *BasicLexer) skipWhitespace() {
	if !r.IsReachedEOF && r.isOnSpace() {
		r.advance()
	}
}

func (r *BasicLexer) skipComment() {
	for r.currentRune() != '}' {
		r.advance()
	}
	r.advance()
}

func (r *BasicLexer) handleNoValueToken(symbol rune, token BasicToken) (BasicToken, error) {
	if r.currentRune() == symbol {
		r.advance()
		return token, nil
	}
	return BasicToken{}, fmt.Errorf("Got rune %v, expected %v", r.currentRune(), symbol)
}

func (r *BasicLexer) peek() *byte {
	peekPosition := r.Position + 1
	if peekPosition > len(r.Text) - 1 {
		return nil
	} else {
		c := r.Text[peekPosition]
		return &c
	}
}

func (r *BasicLexer) identifier() BasicToken {
	result := ""
	for !r.IsReachedEOF && (unicode.IsLetter(r.currentRune()) || unicode.IsDigit(r.currentRune())) {
		result += string(*r.currentChar())
		r.advance()
	}

	reserved, ok := ReservedKeywords[result]
	if ok {
		return reserved
	} else {
		return BasicToken{
			TokenType: ID,
			TokenValue: result,
		}
	}
}

func (r *BasicLexer) NextToken() (BasicToken, error) {
	for !r.IsReachedEOF {
		if r.isOnSpace() {
			r.skipWhitespace()
			continue
		}



		currentRune := r.currentRune()

		if currentRune == '{' {
			r.advance()
			r.skipComment()
			continue
		} 
		
		if r.isOnDigit() {
			return r.parseNumber(), nil
		} else if currentRune == '+' {
			token := BasicToken{TokenType: PLUS}
			r.advance()
			return token, nil
		} else if currentRune == '-' {
			token := BasicToken{TokenType: MINUS}
			r.advance()
			return token, nil
		} else if currentRune == '*' {
			token := BasicToken {TokenType: MUL }
			r.advance()
			return token, nil
		} else if currentRune == '/' {
			token := BasicToken { TokenType: FLOAT_DIV}
			r.advance()
			return token, nil
		} else if currentRune == '(' {
			return r.handleNoValueToken('(', BasicToken{TokenType: LPAREN})
		} else if currentRune == ')' {
			return r.handleNoValueToken(')', BasicToken{TokenType: RPAREN})
		} else if currentRune == ':' && r.peekRune() == '=' {
			r.advance()
			r.advance()
			token := BasicToken {TokenType: ASSIGN}
			return token, nil
		} else if currentRune == ';' {
			r.advance()
			token := BasicToken {TokenType: SEMICOLON}
			return token, nil
		} else if currentRune == '.' {
			r.advance()
			token := BasicToken{TokenType: DOT}
			return token, nil
		} else if unicode.IsLetter(currentRune) {
			return r.identifier(), nil
		} else if currentRune == ':' {
			r.advance()
			token := BasicToken{ TokenType: COLON }
			return token, nil
		} else if currentRune == ',' {
			r.advance()
			token := BasicToken{ TokenType: COMMA}
			return token, nil
		} else if currentRune == '/' {
			r.advance()
			token := BasicToken{ TokenType: FLOAT_DIV }
			return token, nil
		}

	}
	token := BasicToken{TokenType: EOF}
	return token, nil
}

func (r *BasicLexer) Eat(tokenType TokenType) (error) {
	if r.CurrentToken.TokenType == tokenType {
		token, err := r.NextToken()
		if err != nil {
			return err
		}

		r.CurrentToken = &token
		return nil
	} 
	return fmt.Errorf("Cannot eat token of type: %v, current token type: %v", tokenType, r.CurrentToken.TokenType)
}


func (r *BasicLexer) GetCurrentToken() *BasicToken {
	return r.CurrentToken
}

func NewLexer(text string) BasicLexer {
	eof := false
	if len(text) == 0 {
		eof = true
	} 

	lexer := BasicLexer {
		Text: text,
		Position: 0,
		CurrentToken: nil,
		IsReachedEOF: eof,
	}

	return lexer
}
