package lexer

import (
	"testing"

	"github.com/tysufa/qfa/token"
)

func TestGetToken(t *testing.T) {
	input := `let foo1 = bar; 
	let toto;
	3 <= 2 == 4 > 1
	42;
	fn test2;`

	l := New(input)

	tests := []struct {
		expectedValue string
		expectedType  token.TokenType
		expectedLine  int
	}{
		{"let", token.LET, 1},
		{"foo1", token.IDENT, 1},
		{"=", token.EQ, 1},
		{"bar", token.IDENT, 1},
		{";", token.SEMICOLON, 1},
		{"\n", token.NL, 1},
		{"let", token.LET, 2},
		{"toto", token.IDENT, 2},
		{";", token.SEMICOLON, 2},
		{"\n", token.NL, 2},
		{"3", token.INT, 3},
		{"<=", token.LEQT, 3},
		{"2", token.INT, 3},
		{"==", token.EQEQ, 3},
		{"4", token.INT, 3},
		{">", token.GT, 3},
		{"1", token.INT, 3},
		{"\n", token.NL, 3},
		{"42", token.INT, 4},
		{";", token.SEMICOLON, 4},
		{"\n", token.NL, 4},
		{"fn", token.FN, 5},
		{"test2", token.IDENT, 5},
		{";", token.SEMICOLON, 5},
		{"", token.EOF, 5},
	}

	for _, tt := range tests {
		tok := l.GetToken()
		if tt.expectedType != tok.Type {
			t.Fatalf("wrong token type, expected '%s', got '%s' instead", tt.expectedType, tok.Type)
		}
		if tt.expectedValue != tok.Value {
			t.Fatalf("wrong token value, expected %s, got %s instead", tt.expectedValue, tok.Value)
		}
		if tt.expectedLine != tok.Line {
			t.Fatalf("wrong token line, expected %v, got %v instead", tt.expectedLine, tok.Line)
		}
	}

}

func TestNextChar(t *testing.T) {
	input := `foo bar`

	test := [7]byte{'f', 'o', 'o', ' ', 'b', 'a', 'r'}

	l := New(input)

	for _, char := range test {
		if char != l.curChar {
			t.Fatalf("wrong character, expected '%c', got '%c' instead.", char, l.curChar)
		}
		l.nextChar()
	}
}
