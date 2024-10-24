package interpreter

import (
	"fmt"

	"github.com/anuarkaliyev23/simple-interpreter-go/public/ast"
	"github.com/anuarkaliyev23/simple-interpreter-go/public/lexer"
)

type EvaluatorVisitor struct {
	GloabalScope map[string]any
}

func (r *EvaluatorVisitor) visitOperationNode(node ast.BinaryOperation) (int, error) {
	operation := node.GetToken().TokenType

	left, err := r.Visit(node.Left)
	if err != nil {
		return 0, err
	}

	right, err := r.Visit(node.Right)
	if err != nil {
		return 0, err
	}

	if operation == lexer.MINUS {
		return left - right, nil
	} else if operation == lexer.PLUS {
		return left + right, nil
	} else if operation == lexer.MUL {
		return left * right, nil
	} else if operation == lexer.INTEGER_DIV {
		return left / right, nil
	} else if operation == lexer.FLOAT_DIV {
		return left / right, nil
	}

	return 0, fmt.Errorf("Cannot evaluate BinaryOperation node %v", node)
}


func (r *EvaluatorVisitor) visitUnaryNode(node ast.UnaryOperation) (int, error) {
	operation := node.GetToken().TokenType

	right, err := r.Visit(node.Right)
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

func (r *EvaluatorVisitor) visitIntNode(node ast.IntNode) (int, error) {
	return node.Value, nil
}

func (r *EvaluatorVisitor) visitRealNode(node ast.RealNode) (float64, error) {
	return node.Value, nil
}

func (r *EvaluatorVisitor) visitCompound(node ast.Compound) (int, error) {
	for _, v := range node.Children {
		r.Visit(v)
	}
	return 0, nil
}

func (r *EvaluatorVisitor) visitNoOp(node ast.NoOp) (int, error) {
	return 0, nil 
}

func (r *EvaluatorVisitor) visitAssign(node ast.AssignOperation) (int, error) {
	varName := node.Left.(ast.Var).Value
	rightValue, err := r.Visit(node.Right)
	if err != nil {
		return ErrorCode, nil
	}
	r.GloabalScope[varName] = rightValue
	return 0, nil
}

func (r *EvaluatorVisitor) visitVar(node ast.Var) (int, error) {
	varName := node.Value
	varValue, ok := r.GloabalScope[varName]
	if !ok {
		return ErrorCode, fmt.Errorf("var %v is not initialized", varName)
	}
	return varValue.(int), nil
}

func (r *EvaluatorVisitor) visitProgram(node ast.Program) (int, error) {
	return r.visitBlock(node.Block)
}

func (r *EvaluatorVisitor) visitBlock(node ast.Block) (int, error) {
	for _, v := range node.Declarations {
		r.visitVarDeclaration(v)
	}
	return r.visitCompound(node.Compound)
}

func (r *EvaluatorVisitor) visitVarDeclaration(node ast.VarDeclaration) (int, error) {
	return 0, nil
}

func (r *EvaluatorVisitor) visitTypeSpec(node ast.TypeSpec) (int, error) {
	return 0, nil
}

//TODO update visits for new ast nodes
func (r *EvaluatorVisitor) Visit(node ast.Node) (int, error) {
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

	// castedRealNode, ok := node.(ast.RealNode)
	// if ok {
	// 	return r.visitRealNode(castedRealNode)
	// }

	castedVarNode, ok := node.(ast.Var)
	if ok {
		return r.visitVar(castedVarNode)
	}

	castedVarDeclarationNode, ok := node.(ast.VarDeclaration)
	if ok {
		return r.visitVarDeclaration(castedVarDeclarationNode)
	}

	castedCompountNode, ok := node.(ast.Compound)
	if ok {
		return r.visitCompound(castedCompountNode)
	}

	castedBlockNode, ok := node.(ast.Block)
	if ok {
		return r.visitBlock(castedBlockNode)
	}

	castedProgramNode, ok := node.(ast.Program)
	if ok {
		return r.visitProgram(castedProgramNode)
	}

	castedTypeSpecNode, ok := node.(ast.TypeSpec)
	if ok {
		return r.visitTypeSpec(castedTypeSpecNode)
	}

	castedAssignNode, ok := node.(ast.AssignOperation)
	if ok {
		return r.visitAssign(castedAssignNode)
	}

	castedNoOpNode, ok := node.(ast.NoOp)
	if ok {
		return r.visitNoOp(castedNoOpNode)
	}

	return 0, fmt.Errorf("Cannot evaluate node of unknown type %v", node)
}

func NewEvaluatorVisitor() EvaluatorVisitor {
	return EvaluatorVisitor{
		GloabalScope: map[string]any{},
	}
}
