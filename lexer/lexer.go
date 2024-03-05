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

func createToken(t token.TokenType, l byte) token.Token {
	return token.Token{Type: t, Literal: string(l)}
}

func isLetter(ch byte) bool {
	if ('A' <= ch && ch <= 'Z') || ('a' <= ch && ch <= 'z') || ch == '_' {
		return true
	} else {
		return false
	}
}

func (l *Lexer) nextChar() {
	if !endOfFile(l.input, l.pos) {
		l.ch = l.input[l.nextPos]
		l.pos = l.nextPos
		l.nextPos++
	} else {
		l.ch = 0
	}
}

func (l *Lexer) readIdentifier() token.Token {
	tok := token.Token{}
	for isLetter(l.ch) {
		tok.Literal += string(l.ch)
		l.nextChar()
	}

	switch tok.Literal {
	case "let":
		tok.Type = token.LET
	case "fn":
		tok.Type = token.FUNCTION
	default:
		tok.Type = token.IDENT
	}

	return tok
}

func (l *Lexer) readInt() token.Token {
	tok := token.Token{Type: token.INT}
	for isDigit(l.ch) {
		tok.Literal += string(l.ch)
		l.nextChar()
	}
	return tok
}

func isDigit(ch byte) bool {
	if '0' <= ch && ch <= '9' {
		return true
	}
	return false
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.nextChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	tok := token.Token{}

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = createToken(token.EQUAL, l.ch)
	case '+':
		tok = createToken(token.ADD, l.ch)
	case '(':
		tok = createToken(token.L_PAR, l.ch)
	case ')':
		tok = createToken(token.R_PAR, l.ch)
	case '{':
		tok = createToken(token.L_BR, l.ch)
	case '}':
		tok = createToken(token.R_BR, l.ch)
	case ',':
		tok = createToken(token.COMMA, l.ch)
	case ';':
		tok = createToken(token.SEMICOLON, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok = l.readIdentifier()
		} else if isDigit(l.ch) {
			tok = l.readInt()
		} else {
			tok = createToken(token.ILLEGAL, l.ch)
		}
		return tok
	}

	l.nextChar()

	return tok
}

func endOfFile(s string, pos int) bool {
	if pos >= len(s)-1 {
		return true
	} else {
		return false
	}
}
