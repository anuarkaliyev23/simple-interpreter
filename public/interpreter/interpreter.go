package interpreter

import (
	"fmt"

	"github.com/anuarkaliyev23/simple-interpreter-go/public/lexer"
)

type BasicInterpreter struct {
	Lexer lexer.BasicLexer	
}


func (r *BasicInterpreter) Expr() (any, error) {
	err := r.Lexer.Initialize()
	if err != nil {
		return nil, err
	}

	left := r.Lexer.CurrentToken
	err = r.Lexer.Eat(lexer.INTEGER)
	if err != nil {
		return nil, err
	}

	op := r.Lexer.CurrentToken
	if op.TokenType == lexer.PLUS {
		err = r.Lexer.Eat(lexer.PLUS)
		if err != nil {
			return nil, err
		}
	}

	if op.TokenType == lexer.MINUS {
		err = r.Lexer.Eat(lexer.MINUS)
		if err != nil {
			return nil, err
		}
	}

	right := r.Lexer.CurrentToken
	err = r.Lexer.Eat(lexer.INTEGER)

	leftValue := left.TokenValue
	if err != nil {
		return nil, err
	}

	rightValue := right.TokenValue
	if err != nil {
		return nil, err
	}

	if op.TokenType == lexer.PLUS {
		return leftValue.(int) + rightValue.(int), nil
	} else if op.TokenType == lexer.MINUS {
		return leftValue.(int) - rightValue.(int), nil
	} else {
		return nil, fmt.Errorf("Operation is not permitted")
	}
}
