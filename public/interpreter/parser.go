package interpreter

import (
	"fmt"

	"github.com/anuarkaliyev23/simple-interpreter-go/public/ast"
	"github.com/anuarkaliyev23/simple-interpreter-go/public/lexer"
)

type Lexer interface {
	Initialize() error
	NextToken() (lexer.BasicToken, error)
	Eat(lexer.TokenType) error
	GetCurrentToken() *lexer.BasicToken
}

type BasicParser struct {
	Lexer Lexer
}

func (r *BasicParser) factor() (ast.Node, error) {
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
		return ast.NewIntNode(*token)
	} else if token.TokenType == lexer.REAL {
		err := r.Lexer.Eat(lexer.REAL)
		if err != nil {
			return nil, err
		}
		return ast.NewRealNode(*token)
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
	} else if token.TokenType == lexer.ID {
		return r.variable()
	}

	return nil, fmt.Errorf("Could not read factor")
}

// factor((MUL | INTEGER_DIV | FLOAT_DIV) factor)*
func (r *BasicParser) term() (ast.Node, error) {
	node, err := r.factor()
	if err != nil {
		return nil, err
	}

	for r.isValidToken(*r.Lexer.GetCurrentToken(), lexer.MUL, lexer.FLOAT_DIV, lexer.INTEGER_DIV) {
		token := r.Lexer.GetCurrentToken()

		if token.TokenType == lexer.MUL {
			err = r.Lexer.Eat(lexer.MUL)
			if err != nil {
				return nil, err
			}
		} else if token.TokenType == lexer.INTEGER_DIV {
			err = r.Lexer.Eat(lexer.INTEGER_DIV)
			if err != nil {
				return nil, err
			}
		} else if token.TokenType == lexer.FLOAT_DIV {
			err = r.Lexer.Eat(lexer.FLOAT_DIV)
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

func (r *BasicParser) isValidToken(token lexer.BasicToken, types ...lexer.TokenType) bool {
	for _, t := range types {
		if token.TokenType == t {
			return true
		}
	}
	return false
}

// term ((PLUS|MINUS) term)*
func (r *BasicParser) Expr() (ast.Node, error) {
	node, err := r.term()
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

func (r *BasicParser) empty() ast.Node {
	return ast.NewNoOp()
}

// var: ID
func (r *BasicParser) variable() (ast.Node, error) {
	token := r.Lexer.GetCurrentToken()
	node, err := ast.NewVar(*token)
	return node, err
}

// assignment: variable ASSIGN expr
func (r *BasicParser) assignment() (ast.Node, error) {
	left, err := r.variable()
	if err != nil {
		return nil, err
	}

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
func (r *BasicParser) statementList() ([]ast.Node, error) {
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
func (r *BasicParser) compound() (ast.Node, error) {
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
func (r *BasicParser) statement() (ast.Node, error) {
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

// program: compound DOT
func (r *BasicParser) program() (ast.Node, error) {
	r.Lexer.Eat(lexer.PROGRAM)
	varNode, err := r.variable()
	programName := varNode.GetToken().TokenValue

	r.Lexer.Eat(lexer.SEMICOLON)

	block, err := r.block()
	if err != nil {
		return nil, err	
	}

	program := ast.NewProgram(programName, block.(ast.Block), *r.Lexer.GetCurrentToken())
	r.Lexer.Eat(lexer.DOT)
	return program, nil
}

// INTEGER | REAL
func (r *BasicParser) typeSpec() (ast.Node, error) {
	token := r.Lexer.GetCurrentToken()
	if token.TokenType == lexer.INTEGER_DECLARAION {
		err := r.Lexer.Eat(lexer.INTEGER_DECLARAION)
		if err != nil {
			return nil, err
		}
		return ast.NewTypeSpec(*token), nil
	} else if token.TokenType == lexer.REAL {
		err := r.Lexer.Eat(lexer.INTEGER_DECLARAION)
		if err != nil {
			return nil, err
		}
		return ast.NewRealNode(*token)
	}

	return nil, fmt.Errorf("Unknown type specification %v", token.TokenType)
}


func (r *BasicParser) block() (ast.Node, error) {
	declarationNodes, err := r.declarations()
	if err != nil {
		return nil, err
	}

	compoundNode, err := r.compound()
	if err != nil {
		return nil, err
	}


	var castedDeclarations []ast.VarDeclaration
	for _, v := range declarationNodes {
		casted, ok := v.(ast.VarDeclaration)
		if !ok {
			return nil, fmt.Errorf("Cannot cast %v to variable declaration", v)
		}
		castedDeclarations = append(castedDeclarations, casted)
	}

	node := ast.NewBlock(castedDeclarations, compoundNode.(ast.Compound), *r.Lexer.GetCurrentToken())
	return node, nil
}

//declarations: VAR (varDeclaration SEMICOLON)+ | empty
func (r *BasicParser) declarations() ([]ast.Node, error) {
	var declarations []ast.Node

	if r.Lexer.GetCurrentToken().TokenType == lexer.VAR {
		r.Lexer.Eat(lexer.VAR)
		for r.Lexer.GetCurrentToken().TokenType == lexer.ID {
			declaration, err := r.varDeclaration()
			if err != nil {
				return nil, err
			}
			declarations = append(declarations, declaration...)
			r.Lexer.Eat(lexer.SEMICOLON)
		}
	}

	return declarations, nil
}

// varDeclaration: ID (COMMA ID)* COLON typeSpec
func (r *BasicParser) varDeclaration() ([]ast.Node, error) {
	var varNodes []ast.Var
	firstVar, err := ast.NewVar(*r.Lexer.GetCurrentToken())
	if err != nil {
		return nil, err
	}

	varNodes = append(varNodes, firstVar)
	for r.Lexer.GetCurrentToken().TokenType == lexer.COMMA {
		r.Lexer.Eat(lexer.COMMA)
		varNode, err := ast.NewVar(*r.Lexer.GetCurrentToken())
		if err != nil {
			return nil, err
		}

		varNodes = append(varNodes, varNode)
		r.Lexer.Eat(lexer.ID)
	}
	r.Lexer.Eat(lexer.COLON)

	typeNode, err := r.typeSpec()
	
	var declarations []ast.Node
	for _, v := range varNodes {
		d := ast.NewVarDeclaration(v, typeNode.(ast.TypeSpec), *r.Lexer.GetCurrentToken())
		declarations = append(declarations, d)
	}

	return declarations, nil
}

func (r *BasicParser) Parse() (ast.Node, error) {
	node, err := r.program()
	if err != nil {
		return nil, err
	}
	if r.Lexer.GetCurrentToken().TokenType != lexer.EOF {
		return nil, fmt.Errorf("EOF expected, got %v instead", r.Lexer.GetCurrentToken())
	}
	return node, nil
}

func NewParser(lexer lexer.BasicLexer) (*BasicParser, error) {
	interpreter := BasicParser{
		Lexer: &lexer,
	}

	if err := interpreter.Lexer.Initialize(); err != nil {
		return nil, err
	}

	return &interpreter, nil
}
