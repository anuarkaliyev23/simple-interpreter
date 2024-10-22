package interpreter

import (
	"fmt"

	"github.com/anuarkaliyev23/simple-interpreter-go/public/ast"
	"github.com/anuarkaliyev23/simple-interpreter-go/public/lexer"
)

type NodeVisitor interface {
	Visit(node ast.Node) (int, error)
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

//factor (PLUS|MINUS) integer | LPAREN expr RPAREN | variable
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

func (r *BasicInterpreter) empty() (ast.Node) {
	return ast.NewNoOp()
}

// var: ID
func (r *BasicInterpreter) variable() (ast.Node) {
	token := r.Lexer.GetCurrentToken()
	node := ast.NewVar(token.TokenValue.(string), *token)
	return node
}

// assignment: variable ASSIGN expr
func (r *BasicInterpreter) assignment() (ast.Node, error) {
	left := r.variable()
	token := r.Lexer.GetCurrentToken()
	r.Lexer.Eat(lexer.ASSIGN)
	right, err := r.Expr()
	if err != nil {
		return nil, err
	}
	node := ast.NewAssignt(left, right, *token)
	return node, nil
}

// statementList: statement | statement SEMI statementList
func (r *BasicInterpreter) statementList() ([]ast.Node, error) {
	node, err := r.statement()
	if err != nil {
		return nil, err 
	}

	var results []ast.Node
	results = append(results, node)

	for r.Lexer.GetCurrentToken().TokenType == lexer.SEMICOLON {
		r.Lexer.Eat(lexer.SEMICOLON)
		statement, err := r.statement()
		if err != nil {
			return nil, err
		}
		results = append(results, statement)
	}

	if r.Lexer.GetCurrentToken().TokenType == lexer.ID {
		return nil, fmt.Errorf("Cannot parse variables")
	}

	return results, nil
}


// compound: BEGIN statementList END
func (r *BasicInterpreter) compound() (ast.Node, error) {
	r.Lexer.Eat(lexer.BEGIN)
	nodes, err := r.statementList()
	if err != nil {
		return nil, err
	}

	r.Lexer.Eat(lexer.END)

	token := r.Lexer.GetCurrentToken()
	compound := ast.NewCompound(nodes, *token)
	return compound, nil
}

// statement: compound | assignment | empty
func (r *BasicInterpreter) statement() (ast.Node, error) {
	currentToken := r.Lexer.GetCurrentToken()
	var result ast.Node

	if currentToken.TokenType == lexer.BEGIN {
		node, err := r.compound()
		if err != nil {
			return nil, err
		}
		result = node
	} else if currentToken.TokenType == lexer.ID {
		node, err := r.assignment()
		if err != nil {
			return nil, err
		}
		result = node
	} else {
		node := r.empty()
		result = node
	}

	return result, nil 
}

//program: compound DOT
func (r *BasicInterpreter) program() (ast.Node, error) {
	node, err := r.compound()
	if err != nil {
		return nil, err
	}
	r.Lexer.Eat(lexer.DOT)
	return node, nil
}

func (r *BasicInterpreter) parse() (ast.Node, error) {
	node, err := r.program()
	if err != nil {
		return nil, err
	}
	if r.Lexer.GetCurrentToken().TokenType != lexer.EOF {
		return nil, fmt.Errorf("EOF expected, got %v instead", r.Lexer.GetCurrentToken())
	}
	return node, nil
}

func (r *BasicInterpreter) Interpret() (int, error) {
	astTree, err := r.Expr()
	if err != nil {
		return ErrorCode, err
	}

	result, err := r.Visitor.Visit(astTree)
	if err != nil {
		return ErrorCode, err
	}

	return result, nil
}

func NewInterpreter(lexer lexer.BasicLexer) (*BasicInterpreter, error) {
	evaluator := NewEvaluatorVisitor()
	interpreter := BasicInterpreter {
		Lexer: &lexer,
		Visitor: &evaluator,
	}

	if err := interpreter.Lexer.Initialize(); err != nil {
		return nil, err
	}

	return &interpreter, nil
}
