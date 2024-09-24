package interpreter

import (
	"testing"

	"github.com/anuarkaliyev23/simple-interpreter-go/public/lexer"
	"github.com/stretchr/testify/require"
)


func TestBasicInterpreter_Expr(t *testing.T) {
	t.Run("Supported operation gives sum of integers", func(t *testing.T) {
		lexer := lexer.NewLexer("5 + 3")
		interpreter := BasicInterpreter{
			Lexer: lexer,
		}

		result, err := interpreter.Expr()
		require.NoError(t, err)
		require.Equal(t, 8, result)
	})
}
