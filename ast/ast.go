package ast

import (
	"fmt"
	"strings"

	"github.com/jlucero805/golang-lisp/tokens"
)

type Expression interface {
	Print(spaces int)
	ExpressionNode()
}

type Statement interface {
	StatementNode()
}

type Program struct {
	Statements []Statement
}

type SetStatement struct {
	Ident IdentExpression
	Value Expression
}

func (e *SetStatement) StatementNode() {}

func (e *SetStatement) Print(spaces int) {}

type StatementExpression struct {
	Expression Expression
}

func (e *StatementExpression) StatementNode() {}

type ListExpression struct {
	Elements []Expression
}

func (e ListExpression) ExpressionNode() {}

func (e ListExpression) Print(spaces int) {
	fmt.Printf("%s%s", genSpaces(spaces), "(")
	fmt.Printf("\n")
	for _, v := range e.Elements {
		v.Print(spaces + 2)
		fmt.Printf("\n")
	}
	fmt.Printf("%s%s", genSpaces(spaces), ")")
}

type FunctionExpression struct {
	Parameters []IdentExpression
	Body       Expression
}

func (e FunctionExpression) ExpressionNode() {}

func (e FunctionExpression) Print(spaces int) {
	fmt.Printf("%s(lambda (", genSpaces(spaces))

	params := []string{}
	for _, v := range e.Parameters {
		params = append(params, v.Value)
	}
	paramString := strings.Join(params, ", ")
	fmt.Printf("%s) ...)", paramString)
}

type IdentExpression struct {
	Token tokens.Token
	Value string
}

func (e IdentExpression) ExpressionNode() {}

func (e IdentExpression) Print(spaces int) {
	fmt.Printf("%s%v", genSpaces(spaces), e.Value)
}

type NumberExpression struct {
	Token tokens.Token
	Value int
}

func (e NumberExpression) ExpressionNode() {}

func (e NumberExpression) Print(spaces int) {
	fmt.Printf("%s%v [type number]", genSpaces(spaces), e.Value)
}

func genSpaces(num int) string {
	spaces := []string{}
	i := 0
	for i < num {
		spaces = append(spaces, " ")
		i += 1
	}
	return strings.Join(spaces, "")
}
