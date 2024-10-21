package ast

import (
	"testing"

	"github.com/anuarkaliyev23/simple-interpreter-go/public/lexer"
)

var _ BinaryNode = BinaryOperation{
	left: ValueNode[int]{
		BasicNode: BasicNode{
			token: lexer.BasicToken{TokenType: lexer.INTEGER, TokenValue: 5},
		},
		value: 5,
	},
	right: ValueNode[int]{
		BasicNode: BasicNode{
			token: lexer.BasicToken{TokenType: lexer.INTEGER, TokenValue: 5},
		},
		value: 5,
	},

	BasicNode: BasicNode{
		token: lexer.BasicToken{TokenType: lexer.PLUS},
	},
}

func TestNumberNode(t *testing.T) {
}
