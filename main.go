package main

import (
	"os"

	"github.com/anuarkaliyev23/simple-interpreter-go/public/repl"
)

func main() {
	repl := repl.NewRepl()
	for {
		err := repl.Iter()
		if err != nil {
			os.Exit(1)
		}
	}
}
