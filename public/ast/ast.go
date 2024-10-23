package ast

import (
	"fmt"
	"strconv"

	"github.com/anuarkaliyev23/simple-interpreter-go/public/lexer"
)

type Node interface {
	GetToken() lexer.BasicToken
}

type BinaryNode interface {
	Node
	GetLeft() Node
	GetRight() Node
}

type BasicNode struct {
	token lexer.BasicToken
}

func (r BasicNode) GetToken() lexer.BasicToken {
	return r.token
}

type IntNode struct {
	BasicNode
	value int
}

func (r IntNode) GetValue() int {
	return r.value
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
		value: parsedValue,
		BasicNode: BasicNode{
			token: t,
		},
	}, nil
}

type BinaryOperation struct {
	BasicNode
	left  Node
	right Node
}

func (r BinaryOperation) GetLeft() Node {
	return r.left
}

func (r BinaryOperation) GetRight() Node {
	return r.right
}

func NewBinaryOperation(left Node, right Node, operation lexer.BasicToken) BinaryOperation {
	return BinaryOperation{
		BasicNode: BasicNode{
			token: operation,
		},
		left:  left,
		right: right,
	}
}

type UnaryOperation struct {
	BasicNode
	right Node
}

func (r UnaryOperation) GetRight() Node {
	return r.right
}

func NewUnaryOperation(right Node, operation lexer.BasicToken) UnaryOperation {
	return UnaryOperation{
		BasicNode: BasicNode{
			token: operation,
		},
		right: right,
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
	value string
}

func (r Var) GetValue() string {
	return r.value
}

func NewVar(token lexer.BasicToken) (Var, error) {
	if token.TokenType != lexer.ID {
		return Var{}, fmt.Errorf("Cannot parse token %v to AST var node", token)
	}

	return Var{
		BasicNode: BasicNode{
			token: token,
		},
		value: token.TokenValue,
	}, nil
}

type Compound struct {
	BasicNode
	children []Node
}

func (r Compound) GetChildren() []Node {
	return r.children
}

func NewCompound(children []Node, token lexer.BasicToken) Compound {
	return Compound{
		BasicNode: BasicNode{
			token: token,
		},
		children: children,
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
