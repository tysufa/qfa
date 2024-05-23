package parser

import (
	"testing"

	"github.com/tysufa/qfa/ast"
	"github.com/tysufa/qfa/lexer"
)

func TestLetStatements(t *testing.T) {

	tests := []struct {
		input         string
		expectedValue string
	}{
		{"let x = 3;", "x"},
		{"let y = 3;", "y"},
		// {"let foo_bar = bar;", "foo_bar"},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		stmts := p.GetStatements()

		if len(p.errors) > 0 {
			for _, err := range p.errors {
				t.Fatalf(err)
			}
		}

		if len(stmts.Statements) != 1 {
			t.Fatalf("wrong number of stmts.Statements, expected 1, got %v instead", len(stmts.Statements))
		}
		letStmt, ok := stmts.Statements[0].(*ast.LetStatement)
		if !ok {
			t.Fatalf("stmts.Statements[0] is not *ast.LetStatement, got '%T' instead", stmts.Statements[0])
		}
		if letStmt.Name.Value != test.expectedValue {
			t.Fatalf("letStmt.Name is not '%s', got '%s' instead", test.expectedValue, letStmt.Name.Value)
		}
	}
}
