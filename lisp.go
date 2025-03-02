package main

import (
	"github.com/jlucero805/golang-lisp/evaluator"
	"github.com/jlucero805/golang-lisp/lexer"
	"github.com/jlucero805/golang-lisp/parser"
)

var input string = `
(set x 1)

(print ((lambda (x)
	     (+ x
	        x
	        x))
	    x))

(set add (lambda (x y) (+ x y)))

(set x 33)
(set y 66)
(print (add (add x (add x y)) (add x y)))
`

func main() {
	tokens := lexer.Lex(input)
	program := parser.ParseProgram(tokens)
	evaluator.EvaluateProgram(program)
}
