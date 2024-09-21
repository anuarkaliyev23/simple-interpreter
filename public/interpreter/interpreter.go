package interpreter

import (
	"fmt"

	"github.com/anuarkaliyev23/simple-interpreter-go/public/lexer"
)

type Token interface {
	HasValue() bool
	Value() (any, error)
	Type() lexer.TokenType
}

type Lexer interface {
	NextToken() (Token, error)
	Eat(tokenType lexer.TokenType)
	GetCurrentToken() Token
}

type BasicInterpreter struct {
	Lexer Lexer
}

func (r *BasicInterpreter) loadInitialToken() error {
	token, err := r.Lexer.NextToken()
	if err != nil {
		return err
	}
	r.Lexer.Eat(token.Type())
	return nil
}

func (r *BasicInterpreter) Expr() (any, error) {
	err := r.loadInitialToken()
	if err != nil {
		return nil, err
	}


	token := r.Lexer.CurrentToken()
	left := token
	r.Lexer.Eat(lexer.INTEGER)
	
	op := r.Lexer.CurrentToken()
	if op.Type() == lexer.PLUS {
		r.Lexer.Eat(lexer.PLUS)
	}

	if op.Type() == lexer.MINUS {
		r.Lexer.Eat(lexer.MINUS)
	}

	right := r.Lexer.CurrentToken()
	r.Lexer.Eat(lexer.INTEGER)

	leftValue, err := left.Value()
	if err != nil {
		return nil, err
	}

	rightValue, err := right.Value()
	if err != nil {
		return nil, err
	}

	if op.Type() == lexer.PLUS {
		return leftValue.(int) + rightValue.(int), nil
	} else if op.Type() == lexer.MINUS {
		return leftValue.(int) - rightValue.(int), nil
	} else {
		return nil, fmt.Errorf("Operation is not permitted")
	}
}
