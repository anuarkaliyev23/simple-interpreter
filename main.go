package main

import (
	"github.com/anuarkaliyev23/simple-interpreter-go/public/repl"
)

func main() {
	repl := repl.NewRepl()
	for true {
		repl.Iter()
	}
}
