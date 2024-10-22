package interpreter

import (
	"testing"

	"github.com/anuarkaliyev23/simple-interpreter-go/public/ast"
	"github.com/anuarkaliyev23/simple-interpreter-go/public/lexer"
	"github.com/stretchr/testify/require"
)

func TestBasicInterpreter_Expr(t *testing.T) {

	t.Run("Unary Operations", func(t *testing.T) {
		t.Run("-5 + 3", func(t *testing.T) {
			parser := lexer.NewLexer("-5 + 3")
			interpreter, err := NewInterpreter(parser)
			require.NoError(t, err)

			node, err := interpreter.Expr()
			require.NoError(t, err)
			require.IsType(t, ast.BinaryOperation{}, node)
			require.Equal(t, lexer.PLUS, node.GetToken().TokenType)

			unaryNode := node.(ast.BinaryOperation).GetLeft()
			require.IsType(t, ast.UnaryOperation{}, unaryNode)
			require.Equal(t, lexer.MINUS, unaryNode.GetToken().TokenType)

			require.Equal(t, ast.NewIntNode(
				lexer.BasicToken{
					TokenType:  lexer.INTEGER,
					TokenValue: 5,
				},
				5,
			), unaryNode.(ast.UnaryOperation).GetRight())

			require.Equal(t, ast.NewIntNode(
				lexer.BasicToken{
					TokenType:  lexer.INTEGER,
					TokenValue: 3,
				},
				3,
			), node.(ast.BinaryOperation).GetRight())
		})

		t.Run("+5 + 3", func(t *testing.T) {
			parser := lexer.NewLexer("+5 + 3")
			interpreter, err := NewInterpreter(parser)
			require.NoError(t, err)

			node, err := interpreter.Expr()
			require.NoError(t, err)
			require.IsType(t, ast.BinaryOperation{}, node)
			require.Equal(t, lexer.PLUS, node.GetToken().TokenType)

			unaryNode := node.(ast.BinaryOperation).GetLeft()
			require.IsType(t, ast.UnaryOperation{}, unaryNode)
			require.Equal(t, lexer.PLUS, unaryNode.GetToken().TokenType)

			require.Equal(t, ast.NewIntNode(
				lexer.BasicToken{
					TokenType:  lexer.INTEGER,
					TokenValue: 5,
				},
				5,
			), unaryNode.(ast.UnaryOperation).GetRight())

			require.Equal(t, ast.NewIntNode(
				lexer.BasicToken{
					TokenType:  lexer.INTEGER,
					TokenValue: 3,
				},
				3,
			), node.(ast.BinaryOperation).GetRight())
		})

		t.Run("+5 + -3", func(t *testing.T) {
			parser := lexer.NewLexer("+5 + -3")
			interpreter, err := NewInterpreter(parser)
			require.NoError(t, err)

			node, err := interpreter.Expr()
			require.NoError(t, err)
			require.IsType(t, ast.BinaryOperation{}, node)
			require.Equal(t, lexer.PLUS, node.GetToken().TokenType)

			leftUnaryNode := node.(ast.BinaryOperation).GetLeft()
			require.IsType(t, ast.UnaryOperation{}, leftUnaryNode)
			require.Equal(t, lexer.PLUS, leftUnaryNode.GetToken().TokenType)

			require.Equal(t, ast.NewIntNode(
				lexer.BasicToken{
					TokenType:  lexer.INTEGER,
					TokenValue: 5,
				},
				5,
			), leftUnaryNode.(ast.UnaryOperation).GetRight())

			rightUnaryNode := node.(ast.BinaryOperation).GetRight()
			require.IsType(t, ast.UnaryOperation{}, rightUnaryNode)
			require.Equal(t, lexer.MINUS, rightUnaryNode.GetToken().TokenType)

			require.Equal(t, ast.NewIntNode(
				lexer.BasicToken{
					TokenType:  lexer.INTEGER,
					TokenValue: 3,
				},
				3,
			), rightUnaryNode.(ast.UnaryOperation).GetRight())
		})
	})
	t.Run("Binary Operations", func(t *testing.T) {

		t.Run("5 + 3", func(t *testing.T) {
			parser := lexer.NewLexer("5 + 3")
			interpreter, err := NewInterpreter(parser)

			node, err := interpreter.Expr()
			require.NoError(t, err)
			require.IsType(t, ast.BinaryOperation{}, node)
			require.Equal(t, lexer.PLUS, node.GetToken().TokenType)

			require.Equal(t, ast.NewIntNode(
				lexer.BasicToken{
					TokenType:  lexer.INTEGER,
					TokenValue: 5,
				},
				5,
			), node.(ast.BinaryOperation).GetLeft())

			require.Equal(t, ast.NewIntNode(
				lexer.BasicToken{
					TokenType:  lexer.INTEGER,
					TokenValue: 3,
				},
				3,
			), node.(ast.BinaryOperation).GetRight())
		})

		t.Run("(2 + 3) * 4", func(t *testing.T) {
			parser := lexer.NewLexer("(2 + 3) * 4")
			interpreter, err := NewInterpreter(parser)

			node, err := interpreter.Expr()
			require.NoError(t, err)
			require.IsType(t, ast.BinaryOperation{}, node)
			require.Equal(t, lexer.MUL, node.GetToken().TokenType)

			plusNode := node.(ast.BinaryOperation).GetLeft()
			require.IsType(t, ast.BinaryOperation{}, plusNode)
			require.Equal(t, lexer.PLUS, plusNode.GetToken().TokenType)

			require.Equal(t, ast.NewIntNode(
				lexer.BasicToken{
					TokenType:  lexer.INTEGER,
					TokenValue: 2,
				},
				2,
			), plusNode.(ast.BinaryOperation).GetLeft())

			require.Equal(t, ast.NewIntNode(
				lexer.BasicToken{
					TokenType:  lexer.INTEGER,
					TokenValue: 3,
				},
				3,
			), plusNode.(ast.BinaryOperation).GetRight())

			require.Equal(t, ast.NewIntNode(
				lexer.BasicToken{
					TokenType:  lexer.INTEGER,
					TokenValue: 4,
				},
				4,
			), node.(ast.BinaryOperation).GetRight())
		})

	})
}

func TestBasicInterpreter_Interpret(t *testing.T) {

	t.Run("Arithmetics", func(t *testing.T) {
		t.Run("BinaryOperations", func(t *testing.T) {
			t.Run("5 + 3 = 8", func(t *testing.T) {
				parser := lexer.NewLexer("5 + 3")
				interpreter, err := NewInterpreter(parser)
				require.NoError(t, err)

				result, err := interpreter.Interpret()
				require.NoError(t, err)
				require.Equal(t, 8, result)
			})

			t.Run("5 + 3 + 10 = 18", func(t *testing.T) {
				parser := lexer.NewLexer("5 + 3 + 10")
				interpreter, err := NewInterpreter(parser)
				require.NoError(t, err)

				result, err := interpreter.Interpret()
				require.NoError(t, err)
				require.Equal(t, 18, result)
			})

			t.Run("5 * 3 + 10 = 25", func(t *testing.T) {
				parser := lexer.NewLexer("5 * 3 + 10")
				interpreter, err := NewInterpreter(parser)
				require.NoError(t, err)

				result, err := interpreter.Interpret()
				require.NoError(t, err)
				require.Equal(t, 25, result)
			})

			t.Run("(5 + 3) * 2 = 16", func(t *testing.T) {
				parser := lexer.NewLexer("(5 + 3) * 2")
				interpreter, err := NewInterpreter(parser)
				require.NoError(t, err)

				result, err := interpreter.Interpret()
				require.NoError(t, err)
				require.Equal(t, 16, result)
			})
		})

		t.Run("Unary Operations", func(t *testing.T) {
			t.Run("5 + -3", func(t *testing.T) {
				parser := lexer.NewLexer("5 + -3")
				interpreter, err := NewInterpreter(parser)
				require.NoError(t, err)
				
				result, err := interpreter.Interpret()
				require.NoError(t, err)
				require.Equal(t, 2, result)
			})


			t.Run("-(-3)", func(t *testing.T) {
				parser := lexer.NewLexer("-(-3)")
				interpreter, err := NewInterpreter(parser)
				require.NoError(t, err)
				
				result, err := interpreter.Interpret()
				require.NoError(t, err)
				require.Equal(t, 3, result)
			})
		})
	})
}
