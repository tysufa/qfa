package parser

import (
	"github.com/tysufa/qfa/ast"
	"github.com/tysufa/qfa/lexer"
	"github.com/tysufa/qfa/token"
)

type Parser struct {
	lex       lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l lexer.Lexer) *Parser {
	cur := l.GetToken()
	peek := l.GetToken()
	return &Parser{lex: l, curToken: cur, peekToken: peek}
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lex.GetToken()
}

func (p *Parser) GetStatements() ast.Program {
	res := ast.Program{}

	for p.curToken.Type != token.EOF {
		print(p.curToken.Type)
		switch p.curToken.Type {
		case token.LET:
			res.Statements = append(res.Statements, p.parseLet())
		}
		p.nextToken()
	}
	return res
}

func (p *Parser) parseLet() *ast.LetStatement {
	return &ast.LetStatement{Token: p.curToken, Name: &ast.Identifier{Value: p.peekToken.Value}}
}
