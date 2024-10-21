package interpreter

import (
	"testing"

	"github.com/anuarkaliyev23/simple-interpreter-go/public/lexer"
	"github.com/stretchr/testify/require"
)


func TestBasicInterpreter_Visit(t *testing.T) {
	t.Run("5 + 3 = 8", func(t *testing.T) {
		lexer := lexer.NewLexer("5 + 3")
		interpreter, err := NewInterpreter(lexer)
		require.NoError(t, err)

		result, err := interpreter.Expr()
		require.NoError(t, err)
		require.Equal(t, 8, interpreter.Visit(result))
	})

	t.Run("5 + 3 + 10 = 18", func(t *testing.T) {
		lexer := lexer.NewLexer("5 + 3 + 10")
		interpreter, err := NewInterpreter(lexer)
		require.NoError(t, err)

		result, err := interpreter.Expr()
		require.NoError(t, err)
		require.Equal(t, 18, interpreter.Visit(result))
	})


	t.Run("5 * 3 + 10 = 25", func(t *testing.T) {
		lexer := lexer.NewLexer("5 * 3 + 10")
		interpreter, err := NewInterpreter(lexer)
		require.NoError(t, err)

		result, err := interpreter.Expr()
		require.NoError(t, err)
		require.Equal(t, 25, interpreter.Visit(result))
	})


	t.Run("(5 + 3) * 2 = 16", func(t *testing.T) {
		lexer := lexer.NewLexer("(5 + 3) * 2")
		interpreter, err := NewInterpreter(lexer)
		require.NoError(t, err)

		result, err := interpreter.Expr()
		require.NoError(t, err)
		require.Equal(t, 16, interpreter.Visit(result))
	})
}
