package lexer

import (
	"github.com/tysufa/qfa/token"
)

type Lexer struct {
	curChar  byte
	peekChar byte
	input    string
	pos      int
	line     int
}

func New(input string) *Lexer {
	l := &Lexer{input: input, pos: -1, line: 1}
	l.nextChar()
	l.nextChar()

	return l
}

func (l *Lexer) GetToken() token.Token {
	// TODO: check for strings and floating numbers
	var tok token.Token

	l.skipSpaces()
	switch l.curChar {
	case 0:
		tok.Type = token.EOF
		tok.Value = ""
		tok.Line = l.line
	case '\n':
		tok.Type = token.NL
		tok.Value = "\n"
		tok.Line = l.line
		l.line++
	case ';':
		tok.Type = token.SEMICOLON
		tok.Value = string(l.curChar)
		tok.Line = l.line
	case '!':
		tok.Type = token.BANG
		tok.Value = string(l.curChar)
		tok.Line = l.line
	case '/':
		tok.Type = token.SLASH
		tok.Value = string(l.curChar)
		tok.Line = l.line
	case '*':
		tok.Type = token.STAR
		tok.Value = string(l.curChar)
		tok.Line = l.line
	case '-':
		tok.Type = token.MINUS
		tok.Value = string(l.curChar)
		tok.Line = l.line
	case '+':
		tok.Type = token.PLUS
		tok.Value = string(l.curChar)
		tok.Line = l.line
	case '(':
		tok.Type = token.LPAR
		tok.Value = string(l.curChar)
		tok.Line = l.line
	case ')':
		tok.Type = token.RPAR
		tok.Value = string(l.curChar)
		tok.Line = l.line
	case '{':
		tok.Type = token.LBR
		tok.Value = string(l.curChar)
		tok.Line = l.line
	case '}':
		tok.Type = token.RBR
		tok.Value = string(l.curChar)
		tok.Line = l.line
	case '=':
		tok.Type = token.EQ
		tok.Value = string(l.curChar)
		tok.Line = l.line
	default:
		if isLetter(l.curChar) {
			literal := l.getWord()
			if token.Reserved[literal] != "" {
				tok.Type = token.Reserved[literal]
			} else {
				tok.Type = token.IDENT
			}
			tok.Value = literal
			tok.Line = l.line
		} else if isNumber(l.curChar) {
			nb := l.getInt()
			tok.Type = token.INT
			tok.Value = nb
			tok.Line = l.line
		}
	}

	l.nextChar()
	return tok
}

func (l *Lexer) getInt() string {
	res := ""
	for isNumber(l.peekChar) {
		res += string(l.curChar)
		l.nextChar()
	}
	res += string(l.curChar)

	return res
}

func (l *Lexer) getWord() string {
	res := ""
	for isLetter(l.peekChar) || isNumber(l.peekChar) {
		res += string(l.curChar)
		l.nextChar()
	}
	res += string(l.curChar)

	return res
}

func isNumber(char byte) bool {
	return ('0' <= char && char <= '9')
}

func isLetter(char byte) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || (char == '_')
}

func (l *Lexer) skipSpaces() {
	for l.curChar == ' ' || l.curChar == '\r' || l.curChar == '\t' {
		l.nextChar()
	}
}

func (l *Lexer) nextChar() {
	l.curChar = l.peekChar
	if l.pos+1 < len(l.input) {
		l.pos++
		l.peekChar = l.getCurChar()
	} else {
		l.peekChar = 0
	}
}

func (l *Lexer) getCurChar() byte {
	return l.input[l.pos]
}
