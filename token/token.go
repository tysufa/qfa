package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT     = "IDENT"
	INT       = "INT"
	LET       = "LET"
	ADD       = "+"
	MULT      = "*"
	L_PAR     = "("
	R_PAR     = ")"
	L_BR      = "{"
	R_BR      = "}"
	EQUAL     = "="
	SEMICOLON = ";"
	COMMA     = ","
	FUNCTION  = "FUNCTION"
)

type tokenType string

type Token struct {
	Type    tokenType
	Literal string
}
