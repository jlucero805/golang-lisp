package lexer

import (
	"github.com/jlucero805/golang-lisp/tokens"
)

type Lexer struct {
	Line   int
	Column int
	Tokens []*tokens.Token
	Start  int
	End    int
	Source string
}

func Lex(source string) []*tokens.Token {
	lexer := &Lexer{
		Source: source,
	}

	lexer.lex()

	return lexer.Tokens
}

func (l *Lexer) lex() {
	for l.End < len(l.Source) {
		switch l.cur() {
		case '\n':
			l.incCur()
			l.incStart()
			l.Column = 0
		case '\t':
			l.incCur()
			l.incStart()
		case ' ':
			l.incCur()
			l.incStart()
		case '(':
			l.addToken(tokens.NewToken(
				tokens.WithLexeme("("),
				tokens.WithType(tokens.L_PAREN),
				tokens.WithColumn(l.Column),
				tokens.WithLine(l.Line),
			))
			l.incCur()
			l.incStart()
		case ')':
			l.addToken(tokens.NewToken(
				tokens.WithLexeme(")"),
				tokens.WithType(tokens.R_PAREN),
				tokens.WithColumn(l.Column),
				tokens.WithLine(l.Line),
			))
			l.incCur()
			l.incStart()
		default:
			if isIdentFirstCh(l.cur()) {
				l.incCur()
				l.lexIdent()
			} else if isNumberCh(l.cur()) {
				l.lexNumber()
			}
		}
	}
}

func (l *Lexer) lexNumber() {
	for isNumberCh(l.cur()) {
		l.incCur()
	}
	l.addToken(tokens.NewToken(
		tokens.WithLexeme(l.slice()),
		tokens.WithType(tokens.NUMBER),
		tokens.WithColumn(l.Column),
		tokens.WithLine(l.Line),
	))
	l.incStart()
}

func (l *Lexer) lexIdent() {
	for isIdentCh(l.cur()) {
		l.incCur()
	}
	l.addToken(tokens.NewToken(
		tokens.WithLexeme(l.slice()),
		tokens.WithType(tokens.IDENT),
		tokens.WithColumn(l.Column),
		tokens.WithLine(l.Line),
	))
	l.incStart()
}

func (l *Lexer) incStart() {
	l.Start = l.End
	l.Column += 1
}

func (l *Lexer) incCur() {
	l.End += 1
}

func isNumberCh(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isIdentCh(ch byte) bool {
	return isIdentFirstCh(ch) || (ch >= '0' && ch <= '9') || ch == '-'
}

func isIdentFirstCh(ch byte) bool {
	switch ch {
	case '+', '-', '*', '/', '$', '%', ':', '.':
		return true
	}
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func (l *Lexer) slice() string {
	return l.Source[l.Start:l.End]
}

func (l *Lexer) addToken(token *tokens.Token) {
	l.Tokens = append(l.Tokens, token)
}

func (l *Lexer) cur() byte {
	return l.Source[l.End]
}
