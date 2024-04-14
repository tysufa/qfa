package parser

import (
	"fmt"

	"github.com/tysufa/qfa/ast"
	"github.com/tysufa/qfa/lexer"
	"github.com/tysufa/qfa/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	// read 2 times because the first time curToken is not initialized, only peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	for !p.checkToken(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.checkPeekAdvance(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.checkPeekAdvance(token.ASSIGN) {
		return nil
	}

	for !p.checkToken(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) checkToken(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) checkPeekToken(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) checkPeekAdvance(t token.TokenType) bool {
	if p.checkPeekToken(t) {
		p.nextToken()
		return true
	}
	err := fmt.Sprintf("wrong token expected : '%s' got '%s' instead", t, p.curToken.Type)
	p.errors = append(p.errors, err)
	return false
}
