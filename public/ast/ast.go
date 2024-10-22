package ast

import "github.com/anuarkaliyev23/simple-interpreter-go/public/lexer"

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

func NewIntNode(t lexer.BasicToken, value int) IntNode {
	return IntNode{
		value: value,
		BasicNode: BasicNode{
			token: t,
		},
	}
}


type BinaryOperation struct {
	BasicNode
	left Node
	right Node
}

func (r BinaryOperation) GetLeft() Node {
	return r.left
}

func (r BinaryOperation) GetRight() Node {
	return r.right
}

func NewBinaryOperation(left Node, right Node, operation lexer.BasicToken,) BinaryOperation {
	return BinaryOperation{
		BasicNode: BasicNode{
			token: operation,
		},
		left: left,
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
