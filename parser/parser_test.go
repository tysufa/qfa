package parser

import (
	"github.com/tysufa/qfa/lexer"
	"testing"
)

func TestGetStatement(t *testing.T) {
	input := `let x = 3;`

	l := lexer.New(input)
	p := New(l)
}
