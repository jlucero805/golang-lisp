package parser

import (
	"strconv"

	"github.com/jlucero805/golang-lisp/ast"
	"github.com/jlucero805/golang-lisp/tokens"
)

type Parser struct {
	Tokens  []*tokens.Token
	Cur     int
	Program ast.Program
}

func ParseProgram(tokens []*tokens.Token) ast.Program {
	p := &Parser{Tokens: tokens}
	for !p.eof() {
		expr := p.ParseExpression()
		p.Program.Statements = append(p.Program.Statements, &ast.StatementExpression{Expression: expr})
	}
	return p.Program
}

func (p *Parser) ParseExpression() ast.Expression {
	switch p.cur().Type {
	case tokens.IDENT:
		expr := ast.IdentExpression{
			Token: p.cur(),
			Value: p.cur().Lexeme,
		}
		p.inc()
		return expr
	case tokens.NUMBER:
		val, err := strconv.Atoi(p.cur().Lexeme)
		if err != nil {
			panic("Failure parsing number")
		}
		expr := ast.NumberExpression{
			Token: p.cur(),
			Value: val,
		}
		p.inc()
		return expr
	case tokens.L_PAREN:
		p.inc()
		return p.ParseList()
	}
	return ast.IdentExpression{
		Value: "nil",
	}
}

func (p *Parser) ParseList() ast.Expression {
	list := []ast.Expression{}
	for p.cur().Type != tokens.R_PAREN {
		list = append(list, p.ParseExpression())
	}
	// Eat the closing parenthesis
	p.inc()
	return TransformList(ast.ListExpression{Elements: list})
}

func TransformList(list ast.ListExpression) ast.Expression {
	switch v := list.Elements[0].(type) {
	case ast.IdentExpression:
		if v.Value == "lambda" {
			switch paramsList := list.Elements[1].(type) {
			case ast.ListExpression:
				return ast.FunctionExpression{
					Parameters: ParseFunctionIdents(paramsList),
					Body:       list.Elements[2],
				}
			}
		}
	default:
		return list
	}
	return list
}

func ParseFunctionIdents(list ast.ListExpression) []ast.IdentExpression {
	params := []ast.IdentExpression{}
	for _, value := range list.Elements {
		switch v := value.(type) {
		case ast.IdentExpression:
			params = append(params, v)
		}
	}
	return params
}

func (p *Parser) inc() {
	p.Cur += 1
}

func (p *Parser) cur() tokens.Token {
	return *p.Tokens[p.Cur]
}

func (p *Parser) eof() bool {
	return p.Cur >= len(p.Tokens)
}
