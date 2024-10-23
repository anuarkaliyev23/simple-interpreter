package interpreter

import "github.com/anuarkaliyev23/simple-interpreter-go/public/ast"

const ErrorCode int = 1

type Parser interface {
	Parse() (ast.Node, error)
	Expr() (ast.Node, error)
}

type NodeVisitor interface {
	Visit(node ast.Node) (int, error)
}

type BasicInterpreter struct {
	Parser Parser
	Evaluator NodeVisitor
}

func (r BasicInterpreter) Interpret() (int, error) {
	astTree, err := r.Parser.Parse()
	if err != nil {
		return ErrorCode, err
	}

	result, err := r.Evaluator.Visit(astTree)
	if err != nil {
		return ErrorCode, err
	}

	return result, nil
}
