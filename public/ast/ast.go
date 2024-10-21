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

type ValueNode[T any] struct {
	BasicNode
	value T
}

func (r ValueNode[T]) GetValue() any {
	return r.value
}

func NewValueNode[T any](t lexer.BasicToken, value T) ValueNode[T] {
	return ValueNode[T]{
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
