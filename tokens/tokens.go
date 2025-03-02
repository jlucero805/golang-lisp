package tokens

const (
	L_PAREN = iota
	R_PAREN
	IDENT
	NUMBER
)

type Token struct {
	Line   int
	Column int
	Type   int
	Lexeme string
}

type TokenOptions func(*Token)

func WithLine(line int) TokenOptions {
	return func(t *Token) {
		t.Line = line
	}
}

func WithColumn(col int) TokenOptions {
	return func(t *Token) {
		t.Column = col
	}
}

func WithType(tokenType int) TokenOptions {
	return func(t *Token) {
		t.Type = tokenType
	}
}

func WithLexeme(lexeme string) TokenOptions {
	return func(t *Token) {
		t.Lexeme = lexeme
	}
}

func NewToken(options ...TokenOptions) *Token {
	token := &Token{}

	for _, option := range options {
		option(token)
	}

	return token
}
