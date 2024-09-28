package interpreter

import (
	"fmt"

	"github.com/anuarkaliyev23/simple-interpreter-go/public/lexer"
)

type BasicInterpreter struct {
	Lexer lexer.BasicLexer	
}

//factor integer | LPAREN expr RPAREN
func (r * BasicInterpreter) factor() (any, error) {
	token := r.Lexer.CurrentToken

	if token.TokenType == lexer.INTEGER {
		err := r.Lexer.Eat(lexer.INTEGER)
		if err != nil {
			return 0, err
		}
		value := token.TokenValue
		return value.(int), nil
	} else if token.TokenType == lexer.LPAREN {
		if err := r.Lexer.Eat(lexer.LPAREN); err != nil {
			return nil, err
		}

		result, err := r.Expr()
		if err != nil {
			return nil, err
		}
		if err := r.Lexer.Eat(lexer.RPAREN); err != nil {
			return nil, err
		}
		return result, err
	}
	return nil, fmt.Errorf("Could not read factor")
}

// factor((MUL | DIV) factor)* 
func (r *BasicInterpreter) term() (int, error) {
	result, err := r.factor()
	if err != nil {
		return 0, err
	}

	for r.isValidToken(*r.Lexer.CurrentToken, lexer.MUL, lexer.DIV){
		token := r.Lexer.CurrentToken
		if token.TokenType == lexer.MUL {
			err = r.Lexer.Eat(lexer.MUL)
			if err != nil {
				return 0, err
			}

			term, err := r.factor()
			if err != nil {
				return 0, err
			}
			result = result.(int) * term.(int)
		}

		if token.TokenType == lexer.DIV {
			err = r.Lexer.Eat(lexer.DIV)
			if err != nil {
				return 0, err
			}

			term, err := r.factor()
			if err != nil {
				return 0, err
			}
			result = result.(int) / term.(int)
		}
	}
	
	return result.(int), nil
}

func (r *BasicInterpreter) isValidToken(token lexer.BasicToken, types ...lexer.TokenType) bool {
	for _, t := range types {
		if (token.TokenType == t) {
			return true
		}
	}
	return false
}


// term ((PLUS|MINUS) term)*
func (r *BasicInterpreter) Expr() (any, error) {
	result, err := r.term()
	if err != nil {
		return nil, err
	}
	
	for r.isValidToken(*r.Lexer.CurrentToken, lexer.PLUS, lexer.MINUS) {
		op := r.Lexer.CurrentToken
		if op.TokenType == lexer.PLUS {
			err = r.Lexer.Eat(lexer.PLUS)
			if err != nil {
				return nil, err
			}

			term, err := r.term()
			if err != nil {
				return nil, err
			}
			result = result + term
		}

		if op.TokenType == lexer.MINUS {
			err = r.Lexer.Eat(lexer.MINUS)
			if err != nil {
				return nil, err
			}

			term, err := r.term()
			if err != nil {
				return nil, err
			}
			result = result - term
		}	
	}

	return result, nil
}

func NewInterpreter(lexer lexer.BasicLexer) (*BasicInterpreter, error) {
	interpreter := BasicInterpreter {
		Lexer: lexer,
	}

	if err := interpreter.Lexer.Initialize(); err != nil {
		return nil, err
	}

	return &interpreter, nil
}
