package interpreter

import (
	"strconv"
	"testing"

	"github.com/anuarkaliyev23/simple-interpreter-go/public/ast"
	"github.com/anuarkaliyev23/simple-interpreter-go/public/lexer"
	"github.com/stretchr/testify/require"
)

func TestBasicInterpreter_Expr(t *testing.T) {

	t.Run("Unary Operations", func(t *testing.T) {
		t.Run("-5 + 3", func(t *testing.T) {
			lxr := lexer.NewLexer("-5 + 3")
			interpreter, err := NewInterpreter(lxr)
			require.NoError(t, err)

			node, err := interpreter.Expr()
			require.NoError(t, err)
			require.IsType(t, ast.BinaryOperation{}, node)
			require.Equal(t, lexer.PLUS, node.GetToken().TokenType)

			unaryNode := node.(ast.BinaryOperation).Left
			require.IsType(t, ast.UnaryOperation{}, unaryNode)
			require.Equal(t, lexer.MINUS, unaryNode.GetToken().TokenType)
	

			require.Equal(t, intNode(5) , unaryNode.(ast.UnaryOperation).Right)
	

			require.Equal(t, intNode(3), node.(ast.BinaryOperation).Right)
		})

		t.Run("+5 + 3", func(t *testing.T) {
			lxr := lexer.NewLexer("+5 + 3")
			interpreter, err := NewInterpreter(lxr)
			require.NoError(t, err)

			node, err := interpreter.Expr()
			require.NoError(t, err)
			require.IsType(t, ast.BinaryOperation{}, node)
			require.Equal(t, lexer.PLUS, node.GetToken().TokenType)

			unaryNode := node.(ast.BinaryOperation).Left
			require.IsType(t, ast.UnaryOperation{}, unaryNode)
			require.Equal(t, lexer.PLUS, unaryNode.GetToken().TokenType)

			require.Equal(t, intNode(5), unaryNode.(ast.UnaryOperation).Right)		
			require.Equal(t, intNode(3), node.(ast.BinaryOperation).Right)
		})

		t.Run("+5 + -3", func(t *testing.T) {
			lxr := lexer.NewLexer("+5 + -3")
			interpreter, err := NewInterpreter(lxr)
			require.NoError(t, err)

			node, err := interpreter.Expr()
			require.NoError(t, err)
			require.IsType(t, ast.BinaryOperation{}, node)
			require.Equal(t, lexer.PLUS, node.GetToken().TokenType)

			leftUnaryNode := node.(ast.BinaryOperation).Left
			require.IsType(t, ast.UnaryOperation{}, leftUnaryNode)
			require.Equal(t, lexer.PLUS, leftUnaryNode.GetToken().TokenType)

			require.Equal(t, intNode(5), leftUnaryNode.(ast.UnaryOperation).Right)

			rightUnaryNode := node.(ast.BinaryOperation).Right
			require.IsType(t, ast.UnaryOperation{}, rightUnaryNode)
			require.Equal(t, lexer.MINUS, rightUnaryNode.GetToken().TokenType)

			require.Equal(t, intNode(3), rightUnaryNode.(ast.UnaryOperation).Right)
		})
	})
	t.Run("Binary Operations", func(t *testing.T) {

		t.Run("5 + 3", func(t *testing.T) {
			lxr := lexer.NewLexer("5 + 3")
			interpreter, err := NewInterpreter(lxr)

			node, err := interpreter.Expr()
			require.NoError(t, err)
			require.IsType(t, ast.BinaryOperation{}, node)
			require.Equal(t, lexer.PLUS, node.GetToken().TokenType)


			require.Equal(t, intNode(5), node.(ast.BinaryOperation).Left)
			require.Equal(t, intNode(3), node.(ast.BinaryOperation).Right)
		})

		t.Run("(2 + 3) * 4", func(t *testing.T) {
			lxr := lexer.NewLexer("(2 + 3) * 4")
			interpreter, err := NewInterpreter(lxr)

			node, err := interpreter.Expr()
			require.NoError(t, err)
			require.IsType(t, ast.BinaryOperation{}, node)
			require.Equal(t, lexer.MUL, node.GetToken().TokenType)

			plusNode := node.(ast.BinaryOperation).Left
			require.IsType(t, ast.BinaryOperation{}, plusNode)
			require.Equal(t, lexer.PLUS, plusNode.GetToken().TokenType)


			
			require.Equal(t, intNode(2), plusNode.(ast.BinaryOperation).Left)

			require.Equal(t, intNode(3), plusNode.(ast.BinaryOperation).Right)

			require.Equal(t, intNode(4), node.(ast.BinaryOperation).Right)
		})

	})
}

func intNode(value int) ast.IntNode{
	node, _ := ast.NewIntNode(lexer.BasicToken{
		TokenType: lexer.INTEGER,
		TokenValue: strconv.Itoa(value),
	})
	return node
}

// func TestBasicInterpreter_Interpret(t *testing.T) {
// 	t.Run("Arithmetics", func(t *testing.T) {
// 		t.Run("BinaryOperations", func(t *testing.T) {
// 			t.Run("5 + 3 = 8", func(t *testing.T) {
// 				lxr := lexer.NewLexer("5 + 3")
// 				interpreter, err := NewInterpreter(lxr)
// 				require.NoError(t, err)
//
// 				result, err := interpreter.Interpret()
// 				require.NoError(t, err)
// 				require.Equal(t, 8, result)
// 			})
//
// 			t.Run("5 + 3 + 10 = 18", func(t *testing.T) {
// 				lxr := lexer.NewLexer("5 + 3 + 10")
// 				interpreter, err := NewInterpreter(lxr)
// 				require.NoError(t, err)
//
// 				result, err := interpreter.Interpret()
// 				require.NoError(t, err)
// 				require.Equal(t, 18, result)
// 			})
//
// 			t.Run("5 * 3 + 10 = 25", func(t *testing.T) {
// 				lxr := lexer.NewLexer("5 * 3 + 10")
// 				interpreter, err := NewInterpreter(lxr)
// 				require.NoError(t, err)
//
// 				result, err := interpreter.Interpret()
// 				require.NoError(t, err)
// 				require.Equal(t, 25, result)
// 			})
//
// 			t.Run("(5 + 3) * 2 = 16", func(t *testing.T) {
// 				lxr := lexer.NewLexer("(5 + 3) * 2")
// 				interpreter, err := NewInterpreter(lxr)
// 				require.NoError(t, err)
//
// 				result, err := interpreter.Interpret()
// 				require.NoError(t, err)
// 				require.Equal(t, 16, result)
// 			})
// 		})
//
// 		t.Run("Unary Operations", func(t *testing.T) {
// 			t.Run("5 + -3", func(t *testing.T) {
// 				lxr := lexer.NewLexer("5 + -3")
// 				interpreter, err := NewInterpreter(lxr)
// 				require.NoError(t, err)
// 				
// 				result, err := interpreter.Interpret()
// 				require.NoError(t, err)
// 				require.Equal(t, 2, result)
// 			})
//
//
// 			t.Run("-(-3)", func(t *testing.T) {
// 				lxr := lexer.NewLexer("-(-3)")
// 				interpreter, err := NewInterpreter(lxr)
// 				require.NoError(t, err)
// 				
// 				result, err := interpreter.Interpret()
// 				require.NoError(t, err)
// 				require.Equal(t, 3, result)
// 			})
// 		})
// 	})
// }

// func TestBasicInterpreter_Parse(t *testing.T) {
// 	t.Run("variables", func(t *testing.T) {
// 		lxr := lexer.NewLexer(`
// 			BEGIN
// 				BEGIN
//         			number := 2;
//         			a := number * 2;
// 				END;
// 				x := 11;
// 			END.
// 		`)
// 		interpreter, err := NewInterpreter(lxr)
// 		require.NoError(t, err)
// 		parsed, err := interpreter.Parse()
// 		require.NoError(t, err)
// 		require.NotNil(t, parsed)
// 	})
// }
