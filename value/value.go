package value

import (
	"github.com/jlucero805/golang-lisp/ast"
)

const (
	NUMBER = iota
	CLOSURE
	IDENT
)

type Value interface {
	ValueType() int
}

type Ident struct {
	Literal string
}

func (v Ident) ValueType() int {
	return IDENT
}

type Number struct {
	Literal int
}

func (v Number) ValueType() int {
	return NUMBER
}

type Closure struct {
	Parameters []Ident
	Body       ast.Expression
	Env        *Env
}

func (v Closure) ValueType() int {
	return CLOSURE
}
