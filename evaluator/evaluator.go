package evaluator

import (
	"fmt"

	"github.com/jlucero805/golang-lisp/ast"
	"github.com/jlucero805/golang-lisp/value"
)

type Debug struct {
	Num int
}

var debug Debug

func EvaluateProgram(program ast.Program) *value.Env {
	env := value.NewEnv()
	for _, statement := range program.Statements {
		EvaluateStatement(env, statement)
	}
	return env
}

func EvaluateStatement(e *value.Env, statement ast.Statement) {
	switch s := statement.(type) {
	case *ast.StatementExpression:
		P("Eval statement")
		EvaluateExpression(e, s.Expression)
		return
	}
	panic("eval statement failed")
}

func P(s string) {
	fmt.Println(fmt.Sprintf("\n[%d]::%s\n", debug.Num, s))
	debug.Num += 1
}

func EvaluateExpression(e *value.Env, expr ast.Expression) value.Value {
	P("Eval Expression Parent")
	fmt.Println(expr)
	switch v := expr.(type) {
	case ast.IdentExpression:
		P("Eval Ident")
		return e.Get(value.Ident{Literal: v.Value})
	case ast.NumberExpression:
		P("Eval Number")
		return value.Number{Literal: v.Value}
	// Temporary "Call" Expression
	case ast.ListExpression:
		P("Eval List")
		return EvaluateCall(e, v)
	case ast.FunctionExpression:
		P("Function Expression")
		return EvaluateFunction(e, v)
	default:
		fmt.Println("start")
		fmt.Printf("%+v", v)
		fmt.Printf("%+v\n", expr)
		fmt.Println("end")
	}
	panic("eval Expression failed")
}

func EvaluateFunction(e *value.Env, expr ast.FunctionExpression) value.Closure {
	return value.Closure{
		Parameters: EvaluateFunctionParams(expr.Parameters),
		Body:       expr.Body,
		Env:        e,
	}
}

func EvaluateFunctionParams(params []ast.IdentExpression) []value.Ident {
	res := []value.Ident{}
	for _, v := range params {
		res = append(res, value.Ident{Literal: v.Value})
	}
	return res
}

func EvaluateCall(e *value.Env, expr ast.ListExpression) value.Value {
	calleeExpr := expr.Elements[0]
	switch v := calleeExpr.(type) {
	case ast.IdentExpression:
		if v.Value == "+" {
			P("Eval +")
			sum := value.Number{Literal: 0}
			for _, val := range expr.Elements[1:] {
				num := EvaluateExpression(e, val)
				n := ValidateNumber(num)
				sum.Literal += n.Literal
			}
			return sum
		} else if v.Value == "print" {
			P("Eval Print")
			return EvaluatePrint(e, expr.Elements[1])
		} else if v.Value == "set" {
			switch exprIdent := expr.Elements[1].(type) {
			case ast.IdentExpression:
				e.Set(value.Ident{Literal: exprIdent.Value}, EvaluateExpression(e, expr.Elements[2]))
				return value.Number{}
			}
			panic("Failed to set")
		} else {
			callee := e.Get(value.Ident{Literal: v.Value})
			switch calleeV := callee.(type) {
			case value.Closure:
				newEnv := value.WrapEnv(e)
				params := calleeV.Parameters
				arguments := []value.Value{}
				for _, argExpr := range expr.Elements[1:] {
					arguments = append(arguments, EvaluateExpression(e, argExpr))
				}
				if len(params) != len(arguments) {
					panic("function has invalid arity")
				}
				for i, param := range params {
					newEnv.Set(param, arguments[i])
				}
				return EvaluateExpression(newEnv, calleeV.Body)
			}
		}
	case ast.FunctionExpression:
		callee := EvaluateFunction(e, v)
		newEnv := value.WrapEnv(e)
		params := callee.Parameters
		arguments := []value.Value{}
		for _, argExpr := range expr.Elements[1:] {
			arguments = append(arguments, EvaluateExpression(e, argExpr))
		}
		if len(params) != len(arguments) {
			panic("function has invalid arity")
		}
		for i, param := range params {
			newEnv.Set(param, arguments[i])
		}
		return EvaluateExpression(newEnv, callee.Body)
	}
	panic("Eval call failed")
}

func EvaluatePrint(e *value.Env, expr ast.Expression) value.Value {
	fmt.Printf("%v\n", EvaluateExpression(e, expr))
	return value.Number{}
}

func ValidateNumber(v value.Value) value.Number {
	switch val := v.(type) {
	case value.Number:
		return val
	}
	panic("Number Expected")
}
