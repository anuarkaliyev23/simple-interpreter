package lexer

import (
	"fmt"
	"strconv"
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

func (r *BasicLexer) isOnSpace() bool {
	return unicode.IsSpace(r.currentRune())
}

func (r *BasicLexer) isOnDigit() bool {
	return unicode.IsDigit(r.currentRune())
}

func (r *BasicLexer) parseInteger() (int, error) {
	result := ""
	for !r.IsReachedEOF && r.isOnDigit() {
		result += string(*r.currentChar())
		r.advance()
	}

	i, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func (r *BasicLexer) skipWhitespace() {
	if !r.IsReachedEOF && r.isOnSpace() {
		r.advance()
	}
}

func (r *BasicLexer) handleNoValueToken(symbol rune, token BasicToken) (BasicToken, error) {
	if r.currentRune() == symbol {
		r.advance()
		return token, nil
	}
	return BasicToken{}, fmt.Errorf("Got rune %v, expected %v", r.currentRune(), symbol)
}

func (r *BasicLexer) NextToken() (BasicToken, error) {
	for !r.IsReachedEOF {
		if r.isOnSpace() {
			r.skipWhitespace()
			continue
		}

		if r.isOnDigit() {
			result, err := r.parseInteger()
			if err != nil {
				return BasicToken{}, err
			}
			token := BasicToken{TokenType: INTEGER, TokenValue: result}
			return token,  nil
		} else if r.currentRune() == '+' {
			token := BasicToken{TokenType: PLUS}
			r.advance()
			return token, nil
		} else if r.currentRune() == '-' {
			token := BasicToken{TokenType: MINUS}
			r.advance()
			return token, nil
		} else if r.currentRune() == '*' {
			token := BasicToken {TokenType: MUL }
			r.advance()
			return token, nil
		} else if r.currentRune() == '/' {
			token := BasicToken { TokenType: DIV }
			r.advance()
			return token, nil
		} else if r.currentRune() == '(' {
			return r.handleNoValueToken('(', BasicToken{TokenType: LPAREN})
		} else if r.currentRune() == ')' {
			return r.handleNoValueToken(')', BasicToken{TokenType: RPAREN})
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
