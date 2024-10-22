package interpreter

import (
	"fmt"

	"github.com/anuarkaliyev23/simple-interpreter-go/public/ast"
	"github.com/anuarkaliyev23/simple-interpreter-go/public/lexer"
)

type EvaluatorVisitor struct {

}

func (r EvaluatorVisitor) visitOperationNode(node ast.BinaryOperation) (int, error) {
	operation := node.GetToken().TokenType

	left, err := r.Visit(node.GetLeft())
	if err != nil {
		return 0, err
	}

	right, err := r.Visit(node.GetRight())
	if err != nil {
		return 0, err
	}

	if operation == lexer.MINUS {

		return left - right, nil
	} else if operation == lexer.PLUS {
		return left + right, nil
	} else if operation == lexer.MUL {
		return left * right, nil
	} else if operation == lexer.DIV {
		return left / right, nil
	}
	
	return 0, fmt.Errorf("Cannot evaluate BinaryOperation node %v", node)
}


func (r EvaluatorVisitor) visitUnaryNode(node ast.UnaryOperation) (int, error) {
	operation := node.GetToken().TokenType

	right, err := r.Visit(node.GetRight())
	if err != nil {
		return 0, err
	}

	if operation == lexer.PLUS {
		return +right, nil
	} else if operation == lexer.MINUS {
		return -right, nil
	}

	return 0, fmt.Errorf("Cannot evaluate UnaryOperation node %v", node)
}

func (r EvaluatorVisitor) visitIntNode(node ast.IntNode) (int, error) {
	return node.GetValue(), nil
}

func (r EvaluatorVisitor) Visit(node ast.Node) (int, error) {
	castedOpNode, ok := node.(ast.BinaryOperation)
	if ok {
		return r.visitOperationNode(castedOpNode)
	}

	castedIntNode, ok := node.(ast.IntNode)
	if ok {
		return r.visitIntNode(castedIntNode)
	}

	castedUnaryNode, ok := node.(ast.UnaryOperation)
	if ok {
		return r.visitUnaryNode(castedUnaryNode)
	}

	return 0, fmt.Errorf("Cannot evaluate node of unknown type %v", node)
}
