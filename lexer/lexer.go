package lexer

import (
	"github.com/tysufa/qfa/token"
)

type Lexer struct {
	input   string
	pos     int
	nextPos int
	ch      byte
}

func New(input string) Lexer {
	return Lexer{input, 0, 1, input[0]}
}

func createToken(t token.TokenType, l string) token.Token {
	return token.Token{Type: t, Literal: l}
}

func (l *Lexer) NextToken() token.Token {
	if !endOfFile(l.input, l.pos) {
		tok := token.Token{}
		switch l.input[l.pos] {
		case '=':
			tok = createToken(token.EQUAL, "=")
		case '+':
			tok = createToken(token.ADD, "+")
		case '(':
			tok = createToken(token.L_PAR, "(")
		case ')':
			tok = createToken(token.R_PAR, ")")
		case '{':
			tok = createToken(token.L_BR, "{")
		case '}':
			tok = createToken(token.R_BR, "}")
		case ',':
			tok = createToken(token.COMMA, ",")
		case ';':
			tok = createToken(token.SEMICOLON, ";")
		}

		l.pos = l.nextPos
		l.nextPos++
		return tok

	} else {
		return createToken(token.EOF, "")
	}
}

func endOfFile(s string, pos int) bool {
	if pos >= len(s) {
		return true
	} else {
		return false
	}
}
