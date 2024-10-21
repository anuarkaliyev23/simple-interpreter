package interpreter

import (
	"github.com/anuarkaliyev23/simple-interpreter-go/public/ast"
	"github.com/anuarkaliyev23/simple-interpreter-go/public/lexer"
)

type AstNodeExprVisitor struct {

}

func (r AstNodeExprVisitor) visitOperationNode(node ast.BinaryOperation) int {
	if node.GetToken().TokenType == lexer.MINUS {
		return r.Visit(node.GetLeft()) - r.Visit(node.GetRight())
	} else if node.GetToken().TokenType == lexer.PLUS {
		return r.Visit(node.GetLeft()) + r.Visit(node.GetRight())
	} else if node.GetToken().TokenType == lexer.MUL {
		return r.Visit(node.GetLeft()) * r.Visit(node.GetRight())
	} else if node.GetToken().TokenType == lexer.DIV {
		return r.Visit(node.GetLeft()) * r.Visit(node.GetRight())
	}

	panic("Unknown operation")
}


func (r AstNodeExprVisitor) visitIntNode(node ast.IntNode) int {
	return node.GetValue()
}

func (r AstNodeExprVisitor) Visit(node ast.Node) int {
	castedOpNode, ok := node.(ast.BinaryOperation)
	if ok {
		return r.visitOperationNode(castedOpNode)
	}

	castedIntNode, ok := node.(ast.IntNode)
	if ok {
		return r.visitIntNode(castedIntNode)
	}
	
	panic("Cannot cast node to any known type")
}
