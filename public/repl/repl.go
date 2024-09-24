package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/anuarkaliyev23/simple-interpreter-go/public/interpreter"
	"github.com/anuarkaliyev23/simple-interpreter-go/public/lexer"
)

type Interpreter interface {
	Expr() (any, error)
}

type Repl struct {
	Prefix string
	Reader bufio.Reader 
}

func (r Repl) Iter() error {
	fmt.Printf("%v", r.Prefix)
	text, err := r.Reader.ReadString('\n')
	if err != nil {
		return err
	}
	lexer := lexer.NewLexer(text)
	interpreter := interpreter.BasicInterpreter{
		Lexer: lexer,
	}
	result, err := interpreter.Expr()
	if err != nil {
		return err
	}

	fmt.Println(result)
	return nil
}

func NewRepl() Repl {
	return Repl {
		Prefix: "calc>",
		Reader: *bufio.NewReader(os.Stdin),
	}
}
