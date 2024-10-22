package interpreter

import (
	"fmt"

	"github.com/anuarkaliyev23/simple-interpreter-go/public/ast"
	"github.com/anuarkaliyev23/simple-interpreter-go/public/lexer"
)

type NodeVisitor interface {
	Visit(node ast.Node) int
}

type Lexer interface {
	Initialize() (error)
	NextToken() (lexer.BasicToken, error)
	Eat(lexer.TokenType) (error)
	GetCurrentToken() *lexer.BasicToken
}

type BasicInterpreter struct {
	Lexer Lexer	
	Visitor NodeVisitor
}

const ErrorCode int = 1

//factor (PLUS|MINUS) integer | LPAREN expr RPAREN
func (r *BasicInterpreter) factor() (ast.Node, error) {
	token := r.Lexer.GetCurrentToken()
	
	if token.TokenType == lexer.PLUS {
		err := r.Lexer.Eat(lexer.PLUS)
		if err != nil {
			return nil, err
		}

		right, err := r.factor()
		if err != nil {
			return nil, err
		}

		return ast.NewUnaryOperation(right, *token), nil
	} else if token.TokenType == lexer.MINUS {
		err := r.Lexer.Eat(lexer.MINUS)
		if err != nil {
			return nil, err
		}

		right, err := r.factor()
		if err != nil {
			return nil, err
		}

		return ast.NewUnaryOperation(right, *token), nil
	} else if token.TokenType == lexer.INTEGER {
		err := r.Lexer.Eat(lexer.INTEGER)
		if err != nil {
			return nil, err
		}
		value := token.TokenValue
		return ast.NewIntNode(*token, value.(int)), nil
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
func (r *BasicInterpreter) term() (ast.Node, error) {
	node, err := r.factor()
	if err != nil {
		return nil, err
	}

	for r.isValidToken(*r.Lexer.GetCurrentToken(), lexer.MUL, lexer.DIV){
		token := r.Lexer.GetCurrentToken()

		if token.TokenType == lexer.MUL {
			err = r.Lexer.Eat(lexer.MUL)
			if err != nil {
				return nil, err
			}
		} else if token.TokenType == lexer.DIV {
			err = r.Lexer.Eat(lexer.DIV)
			if err != nil {
				return nil, err
			}
		}

		right, err := r.factor()
		if err != nil {
			return nil, err
		}

		node = ast.NewBinaryOperation(node, right, *token)
	}
	
	return node, nil
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
func (r *BasicInterpreter) Expr() (ast.Node, error) {
	node , err := r.term()
	if err != nil {
		return nil, err
	}
	
	for r.isValidToken(*r.Lexer.GetCurrentToken(), lexer.PLUS, lexer.MINUS) {
		token := r.Lexer.GetCurrentToken()
		if token.TokenType == lexer.PLUS {
			err = r.Lexer.Eat(lexer.PLUS)
			if err != nil {
				return nil, err
			}
		} else if token.TokenType == lexer.MINUS {
			err = r.Lexer.Eat(lexer.MINUS)
			if err != nil {
				return nil, err
			}
		}
	
		right, err := r.term()
		if err != nil {
			return nil, err
		}

		node = ast.NewBinaryOperation(node, right, *token)
	}

	return node, nil
}

func (r BasicInterpreter) visit(node ast.Node) int {
	return r.Visitor.Visit(node)
}

func (r BasicInterpreter) Interpret() (int, error) {
	astTree, err := r.Expr()
	if err != nil {
		return ErrorCode, err
	}

	result := r.Visitor.Visit(astTree)
	return result, nil
}

func NewInterpreter(lexer lexer.BasicLexer) (*BasicInterpreter, error) {
	interpreter := BasicInterpreter {
		Lexer: &lexer,
		Visitor: AstNodeEvalVisitor{},
	}

	if err := interpreter.Lexer.Initialize(); err != nil {
		return nil, err
	}

	return &interpreter, nil
}
