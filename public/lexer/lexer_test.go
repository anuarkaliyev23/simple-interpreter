package lexer

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewLexer(t *testing.T) {
	t.Run("On passing empty string EOF is set to true", func(t *testing.T) {
		lexer := NewLexer("")
		require.Equal(t, 0, lexer.Position)
		require.True(t, lexer.IsReachedEOF)
		require.Nil(t, lexer.CurrentToken)
	})

	t.Run("On passing valid non-empty string EOF is set to false", func(t *testing.T) {
		lexer := NewLexer("5 + 3")
		require.Equal(t, 0, lexer.Position)
		require.False(t, lexer.IsReachedEOF)
		require.Nil(t, lexer.CurrentToken)
	})
}

func TestBasicLexer_advance(t *testing.T) {
	t.Run("Advance doesn't update position on empty string", func(t *testing.T) {
		lexer := NewLexer("")
		lexer.advance()

		require.Equal(t, 0, lexer.Position)
		require.True(t, lexer.IsReachedEOF)
	})

	t.Run("Advance updates position on valid string", func(t *testing.T) {
		lexer := NewLexer("5 + 3")
		lexer.advance()

		require.Equal(t, 1, lexer.Position)
		require.False(t, lexer.IsReachedEOF)
	})

	t.Run("Repeated call doesn't update position after eof", func(t *testing.T) {
		lexer := NewLexer("")
		lexer.advance()
		lexer.advance()

		require.Equal(t, 0, lexer.Position)
		require.True(t, lexer.IsReachedEOF)
	})
}

func TestBasicLexer_NextToken(t *testing.T) {
	t.Run("EOF expected", func(t *testing.T) {
		lexer := NewLexer("")
		token, err := lexer.NextToken()
		require.NoError(t, err)
		require.Equal(t, EOF, token.TokenType)
	})

	t.Run("'5 + 3' INTEGER Expected", func(t *testing.T) {
		lexer := NewLexer("5 + 3")
		token, err := lexer.NextToken()
		require.NoError(t, err)
		require.Equal(t, INTEGER, token.TokenType)
		require.Equal(t, 1, lexer.Position) 
		require.Equal(t, 5, token.TokenValue.(int))
	})

	t.Run("Consecutive calls", func(t *testing.T) {
		lexer := NewLexer("5 + 3")
		token, err := lexer.NextToken()
		require.NoError(t, err)
		
		token, err = lexer.NextToken()
		require.NoError(t, err)
		require.Equal(t, PLUS, token.TokenType)
	})

	t.Run("'5 + 3' Whitespaces being ignored", func(t *testing.T) {
		lexer := NewLexer("5  \t+ 3")
		token, err := lexer.NextToken()
		require.NoError(t, err)
		
		token, err = lexer.NextToken()
		require.NoError(t, err)
		require.Equal(t, PLUS, token.TokenType)
	})

	t.Run("'(5 + 3)' parenthesis are handled", func(t *testing.T) {
		lexer := NewLexer("(5 + 3)")
		token, err := lexer.NextToken()
		require.NoError(t, err)
		require.Equal(t, token.TokenType, LPAREN)
	})


	t.Run("'5 + 3' RPAREN is handled", func(t *testing.T) {
		lexer := NewLexer("()")
		token, err := lexer.NextToken()
		require.NoError(t, err)
		require.Equal(t, token.TokenType, LPAREN)

		token2, err := lexer.NextToken()
		require.NoError(t, err)
		require.Equal(t, token2.TokenType, RPAREN)
	})

}

func TestBasicLexer_Eat(t *testing.T) {
	t.Run("Empty string produces no error on EOF", func(t *testing.T) {
		lexer := NewLexer("")
		lexer.Initialize()
		err := lexer.Eat(EOF)
		require.NoError(t, err)
		require.Equal(t, EOF, lexer.CurrentToken.TokenType)
	})

	t.Run("'5 + 3' returns no error on eating INTEGER", func(t *testing.T) {
		lexer := NewLexer("5 + 3")
		err := lexer.Initialize()
		require.NoError(t, err)

		err = lexer.Eat(INTEGER)
		require.NoError(t, err)
	})
}

func TestBasicLexer_Initialize(t *testing.T) {
	t.Run("'5 + 3' initialized on Number", func(t *testing.T) {
		lexer := NewLexer("5 + 3")
		require.Nil(t, lexer.CurrentToken)
		err := lexer.Initialize()
		require.NoError(t, err)
		require.Equal(t, INTEGER, lexer.CurrentToken.TokenType)
		require.Equal(t, 5, lexer.CurrentToken.TokenValue)
	})

	t.Run("Empty string produces EOF", func(t *testing.T) {
		lexer := NewLexer("")
		err := lexer.Initialize()
		require.NoError(t, err)
		require.Equal(t, EOF, lexer.CurrentToken.TokenType)
	})
}
