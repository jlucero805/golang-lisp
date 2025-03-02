package evaluator_test

import (
	"testing"

	"github.com/jlucero805/golang-lisp/evaluator"
	"github.com/jlucero805/golang-lisp/lexer"
	"github.com/jlucero805/golang-lisp/parser"
	"github.com/jlucero805/golang-lisp/value"
)

func Test_EvaluateProgram(t *testing.T) {
	tests := []struct {
		input       string
		resultIdent value.Ident
		result      value.Value
	}{
		{
			`
			(set a 1)
			(set b 2)
			(set c 3)
			(set add (lambda (a b c) (+ a b c)))
			(set result (add a b c))
			`,
			value.Ident{Literal: "result"},
			value.Number{Literal: 6},
		},
		{
			`
			(set abc 123)
			(set lol-bruh 321)
			(set add (lambda (x y)
			          (+ x
			             y)))
			(set result ((lambda (x)
			              (add (add lol-bruh
			                        abc)
			                   x))
			             abc))
			`,
			value.Ident{Literal: "result"},
			value.Number{Literal: 567},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			tokens := lexer.Lex(tt.input)
			program := parser.ParseProgram(tokens)
			e := evaluator.EvaluateProgram(program)
			if e.Get(tt.resultIdent) != tt.result {
				t.Errorf("Expected %+v but got %+v", tt.result, e.Get(tt.resultIdent))
			}
		})
	}
}
