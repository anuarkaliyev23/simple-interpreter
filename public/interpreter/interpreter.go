package interpreter

import (
	"github.com/anuarkaliyev23/simple-interpreter-go/public/lexer"
)

type BasicInterpreter struct {
	Lexer lexer.BasicLexer	
}

//factor integer
func (r * BasicInterpreter) factor() (int, error) {
	token := r.Lexer.CurrentToken
	err := r.Lexer.Eat(lexer.INTEGER)
	if err != nil {
		return 0, err
	}
	value := token.TokenValue
	return value.(int), nil
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
			result = result * term
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
			result = result / term
		}
	}
	
	return result, nil
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
	}

	return result, nil
}
