package parser

import (
	"github.com/tysufa/qfa/lexer"
	"github.com/tysufa/qfa/token"
)

type Parser struct {
	l         lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l lexer.Lexer) *Parser {
	tok := l.GetToken()
	return &Parser{l: l, curToken: tok}
}
