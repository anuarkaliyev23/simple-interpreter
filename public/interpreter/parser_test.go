package interpreter

import (
	"strconv"
	"testing"

	"github.com/anuarkaliyev23/simple-interpreter-go/public/ast"
	"github.com/anuarkaliyev23/simple-interpreter-go/public/lexer"
	"github.com/stretchr/testify/require"
)



// func TestBasicParser_Expr(t *testing.T) {
// 	t.Run("Program", func(t *testing.T) {
// 		t.Run("Empty program block", func(t *testing.T) {
// 			text := `
// 				PROGRAM program;
// 				BEGIN
// 				END.
// 			`
// 			lxr := lexer.NewLexer(text)
// 			parser, err := NewParser(lxr)
// 			require.NoError(t, err)
// 			parsed, err := parser.Parse()
// 		})
// 	})
//
//
// 	t.Run("Unary Operations", func(t *testing.T) {
// 		t.Run("-5 + 3", func(t *testing.T) {
// 			lxr := lexer.NewLexer("-5 + 3")
// 			parser, err := NewParser(lxr)
// 			require.NoError(t, err)
//
// 			node, err := parser.Expr()
// 			require.NoError(t, err)
// 			require.IsType(t, ast.BinaryOperation{}, node)
// 			require.Equal(t, lexer.PLUS, node.GetToken().TokenType)
//
// 			unaryNode := node.(ast.BinaryOperation).Left
// 			require.IsType(t, ast.UnaryOperation{}, unaryNode)
// 			require.Equal(t, lexer.MINUS, unaryNode.GetToken().TokenType)
// 	
//
// 			require.Equal(t, intNode(5) , unaryNode.(ast.UnaryOperation).Right)
// 	
//
// 			require.Equal(t, intNode(3), node.(ast.BinaryOperation).Right)
// 		})
//
// 		t.Run("+5 + 3", func(t *testing.T) {
// 			lxr := lexer.NewLexer("+5 + 3")
// 			parser, err := NewParser(lxr)
// 			require.NoError(t, err)
//
// 			node, err := parser.Expr()
// 			require.NoError(t, err)
// 			require.IsType(t, ast.BinaryOperation{}, node)
// 			require.Equal(t, lexer.PLUS, node.GetToken().TokenType)
//
// 			unaryNode := node.(ast.BinaryOperation).Left
// 			require.IsType(t, ast.UnaryOperation{}, unaryNode)
// 			require.Equal(t, lexer.PLUS, unaryNode.GetToken().TokenType)
//
// 			require.Equal(t, intNode(5), unaryNode.(ast.UnaryOperation).Right)		
// 			require.Equal(t, intNode(3), node.(ast.BinaryOperation).Right)
// 		})
//
// 		t.Run("+5 + -3", func(t *testing.T) {
// 			lxr := lexer.NewLexer("+5 + -3")
// 			parser, err := NewParser(lxr)
// 			require.NoError(t, err)
//
// 			node, err := parser.Expr()
// 			require.NoError(t, err)
// 			require.IsType(t, ast.BinaryOperation{}, node)
// 			require.Equal(t, lexer.PLUS, node.GetToken().TokenType)
//
// 			leftUnaryNode := node.(ast.BinaryOperation).Left
// 			require.IsType(t, ast.UnaryOperation{}, leftUnaryNode)
// 			require.Equal(t, lexer.PLUS, leftUnaryNode.GetToken().TokenType)
//
// 			require.Equal(t, intNode(5), leftUnaryNode.(ast.UnaryOperation).Right)
//
// 			rightUnaryNode := node.(ast.BinaryOperation).Right
// 			require.IsType(t, ast.UnaryOperation{}, rightUnaryNode)
// 			require.Equal(t, lexer.MINUS, rightUnaryNode.GetToken().TokenType)
//
// 			require.Equal(t, intNode(3), rightUnaryNode.(ast.UnaryOperation).Right)
// 		})
// 	})
// 	t.Run("Binary Operations", func(t *testing.T) {
//
// 		t.Run("5 + 3", func(t *testing.T) {
// 			lxr := lexer.NewLexer("5 + 3")
// 			parser, err := NewParser(lxr)
//
// 			node, err := parser.Expr()
// 			require.NoError(t, err)
// 			require.IsType(t, ast.BinaryOperation{}, node)
// 			require.Equal(t, lexer.PLUS, node.GetToken().TokenType)
//
//
// 			require.Equal(t, intNode(5), node.(ast.BinaryOperation).Left)
// 			require.Equal(t, intNode(3), node.(ast.BinaryOperation).Right)
// 		})
//
// 		t.Run("(2 + 3) * 4", func(t *testing.T) {
// 			lxr := lexer.NewLexer("(2 + 3) * 4")
// 			parser, err := NewParser(lxr)
//
// 			node, err := parser.Expr()
// 			require.NoError(t, err)
// 			require.IsType(t, ast.BinaryOperation{}, node)
// 			require.Equal(t, lexer.MUL, node.GetToken().TokenType)
//
// 			plusNode := node.(ast.BinaryOperation).Left
// 			require.IsType(t, ast.BinaryOperation{}, plusNode)
// 			require.Equal(t, lexer.PLUS, plusNode.GetToken().TokenType)
//
//
// 			
// 			require.Equal(t, intNode(2), plusNode.(ast.BinaryOperation).Left)
//
// 			require.Equal(t, intNode(3), plusNode.(ast.BinaryOperation).Right)
//
// 			require.Equal(t, intNode(4), node.(ast.BinaryOperation).Right)
// 		})
//
// 	})
// }

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
// 				parser, err := NewInterpreter(lxr)
// 				require.NoError(t, err)
//
// 				result, err := parser.Interpret()
// 				require.NoError(t, err)
// 				require.Equal(t, 8, result)
// 			})
//
// 			t.Run("5 + 3 + 10 = 18", func(t *testing.T) {
// 				lxr := lexer.NewLexer("5 + 3 + 10")
// 				parser, err := NewInterpreter(lxr)
// 				require.NoError(t, err)
//
// 				result, err := parser.Interpret()
// 				require.NoError(t, err)
// 				require.Equal(t, 18, result)
// 			})
//
// 			t.Run("5 * 3 + 10 = 25", func(t *testing.T) {
// 				lxr := lexer.NewLexer("5 * 3 + 10")
// 				parser, err := NewInterpreter(lxr)
// 				require.NoError(t, err)
//
// 				result, err := parser.Interpret()
// 				require.NoError(t, err)
// 				require.Equal(t, 25, result)
// 			})
//
// 			t.Run("(5 + 3) * 2 = 16", func(t *testing.T) {
// 				lxr := lexer.NewLexer("(5 + 3) * 2")
// 				parser, err := NewInterpreter(lxr)
// 				require.NoError(t, err)
//
// 				result, err := parser.Interpret()
// 				require.NoError(t, err)
// 				require.Equal(t, 16, result)
// 			})
// 		})
//
// 		t.Run("Unary Operations", func(t *testing.T) {
// 			t.Run("5 + -3", func(t *testing.T) {
// 				lxr := lexer.NewLexer("5 + -3")
// 				parser, err := NewInterpreter(lxr)
// 				require.NoError(t, err)
// 				
// 				result, err := parser.Interpret()
// 				require.NoError(t, err)
// 				require.Equal(t, 2, result)
// 			})
//
//
// 			t.Run("-(-3)", func(t *testing.T) {
// 				lxr := lexer.NewLexer("-(-3)")
// 				parser, err := NewInterpreter(lxr)
// 				require.NoError(t, err)
// 				
// 				result, err := parser.Interpret()
// 				require.NoError(t, err)
// 				require.Equal(t, 3, result)
// 			})
// 		})
// 	})
// }

func TestBasicInterpreter_Parse(t *testing.T) {
	t.Run("variables", func(t *testing.T) {
		lxr := lexer.NewLexer(`
			PROGRAM Part10;
			VAR
			   number     : INTEGER;
			   a, b, c, x : INTEGER;
			   y          : REAL;

			BEGIN {Part10}
			   BEGIN
				  number := 2;
				  a := number;
				  b := 10 * a + 10 * number DIV 4;
				  c := a - - b
			   END;
			   x := 11;
			   y := 20 / 7 + 3.14;
			END.  {Part10}
		`)
		parser, err := NewParser(lxr)
		require.NoError(t, err)
		parsed, err := parser.Parse()
		require.NoError(t, err)
		require.NotNil(t, parsed)
	})
}
