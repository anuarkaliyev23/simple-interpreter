package ast

import (
	"fmt"
	"strconv"

	"github.com/anuarkaliyev23/simple-interpreter-go/public/lexer"
)

type Node interface {
	GetToken() lexer.BasicToken
}

type BasicNode struct {
	token lexer.BasicToken
}

func (r BasicNode) GetToken() lexer.BasicToken {
	return r.token
}

type IntNode struct {
	BasicNode
	Value int
}


func NewIntNode(t lexer.BasicToken) (IntNode, error) {
	if (t.TokenType != lexer.INTEGER) {
		return IntNode{}, fmt.Errorf("Cannot parse token %v to Integer AST node", t)
	}

	parsedValue, err := strconv.Atoi(t.TokenValue)
	if err != nil {
		return IntNode{}, err
	}

	return IntNode{
		Value: parsedValue,
		BasicNode: BasicNode{
			token: t,
		},
	}, nil
}


type RealNode struct {
	BasicNode
	Value float64
}

func NewRealNode(t lexer.BasicToken) (RealNode, error) {
	if t.TokenType != lexer.REAL {
		return RealNode{}, fmt.Errorf("Cannot parse token %v to Real AST node", t)
	}

	parsedValue, err := strconv.ParseFloat(t.TokenValue, 64)
	if err != nil {
		return RealNode{}, nil
	}

	return RealNode{
		Value: parsedValue,
		BasicNode: BasicNode{
			token: t,
		},
	}, nil
}

type BinaryOperation struct {
	BasicNode
	Left  Node
	Right Node
}

func NewBinaryOperation(left Node, right Node, operation lexer.BasicToken) BinaryOperation {
	return BinaryOperation{
		BasicNode: BasicNode{
			token: operation,
		},
		Left:  left,
		Right: right,
	}
}

type UnaryOperation struct {
	BasicNode
	Right Node
}

func NewUnaryOperation(right Node, operation lexer.BasicToken) UnaryOperation {
	return UnaryOperation{
		BasicNode: BasicNode{
			token: operation,
		},
		Right: right,
	}
}

type AssignOperation struct {
	BinaryOperation
}

func NewAssignt(left Node, right Node, operation lexer.BasicToken) AssignOperation {
	return AssignOperation{
		BinaryOperation: NewBinaryOperation(left, right, operation),
	}
}

type Var struct {
	BasicNode
	Value string
}

func NewVar(token lexer.BasicToken) (Var, error) {
	if token.TokenType != lexer.ID {
		return Var{}, fmt.Errorf("Cannot parse token %v to AST var node", token)
	}

	return Var{
		BasicNode: BasicNode{
			token: token,
		},
		Value: token.TokenValue,
	}, nil
}

type Compound struct {
	BasicNode
	Children []Node
}

func NewCompound(children []Node, token lexer.BasicToken) Compound {
	return Compound{
		BasicNode: BasicNode{
			token: token,
		},
		Children: children,
	}
}

type NoOp struct {
	BasicNode
}

func NewNoOp() NoOp {
	return NoOp{
		BasicNode: BasicNode{
			token: lexer.BasicToken{
				TokenType: lexer.SEMICOLON,
			},
		},
	}
}

type TypeSpec struct {
	BasicNode
	Value string
}

func NewTypeSpec(token lexer.BasicToken) TypeSpec {
	return TypeSpec{
		BasicNode: BasicNode{
			token: token,
		},
		Value: token.TokenValue,
	}
}

type VarDeclaration struct {
	BasicNode
	Variable Var
	TypeSpec TypeSpec
}

func NewVariableDeclaration(variable Var, typeSpec TypeSpec, token lexer.BasicToken) VarDeclaration {
	return VarDeclaration {
		BasicNode: BasicNode{
			token: token,
		},
		Variable: variable,
		TypeSpec: typeSpec,
	}
}

type Block struct {
	BasicNode
	Declarations []VarDeclaration
	Compound Compound
}

func NewBlock(declarations []VarDeclaration, compound Compound, token lexer.BasicToken) Block {
	return Block{
		BasicNode: BasicNode{
			token: token,
		},
		Declarations: declarations,
		Compound: compound,
	}
}

type Program struct {
	BasicNode
	Name string
	Block Block
}

func NewProgram(name string, block Block, token lexer.BasicToken) Program {
	return Program{
		BasicNode: BasicNode{
			token: token,
		},
		Block: block,
		Name: name,
	}
}
