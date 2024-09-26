package interpreter

import (
	"github.com/anuarkaliyev23/simple-interpreter-go/public/lexer"
)

type BasicInterpreter struct {
	Lexer lexer.BasicLexer	
}

//factor integer
func (r * BasicInterpreter) factor() (int, error) {
	return r.term()
}

// For simplicity, term is an integer
func (r *BasicInterpreter) term() (int, error) {
	value := r.Lexer.GetCurrentToken().TokenValue
	err := r.Lexer.Eat(lexer.INTEGER)
	if err != nil {
		return 0, err
	}
	return value.(int), nil
}

func (r *BasicInterpreter) isValidToken(token lexer.BasicToken, types ...lexer.TokenType) bool {
	for _, t := range types {
		if (token.TokenType == t) {
			return true
		}
	}
	return false
}

// factor((MUL | DIV) factor)* 
func (r *BasicInterpreter) Expr() (any, error) {
	err := r.Lexer.Initialize()
	if err != nil {
		return nil, err
	}

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

		if op.TokenType == lexer.MUL {
			err = r.Lexer.Eat(lexer.MUL)
			if err != nil {
				return nil, err
			}

			term, err := r.term()
			if err != nil {
				return nil, err
			}
			result = result * term
		}

		if op.TokenType == lexer.DIV {
			err = r.Lexer.Eat(lexer.DIV)
			if err != nil {
				return nil, err
			}

			term, err := r.term()
			if err != nil {
				return nil, err
			}
			result = result / term
		}
	}
	return result, nil
}
