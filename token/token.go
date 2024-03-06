package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT     = "IDENT"
	INT       = "INT"
	LET       = "LET"
	ADD       = "+"
	STAR      = "*"
	L_PAR     = "("
	R_PAR     = ")"
	L_BR      = "{"
	R_BR      = "}"
	ASSIGN    = "="
	SEMICOLON = ";"
	COMMA     = ","
	FUNCTION  = "FUNCTION"

	BANG  = "!"
	MINUS = "-"
	SLASH = "/"

	LT = "<"
	GT = ">"

	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
	TRUE   = "TRUE"
	FALSE  = "FALSE"

	EQUAL     = "=="
	NOT_EQUAL = "!="
)

var Keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
