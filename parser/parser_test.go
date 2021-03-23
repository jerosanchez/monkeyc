// parser/parser_test.go

package parser

import (
	"testing"
	"monkeyc/ast"
	"monkeyc/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let foo = 10;
	`

	aLexer := lexer.New(input)
	sut := New(aLexer)

	program := sut.ParseProgram()

	expectedStatementsCount := 2
	assertProgram(program, expectedStatementsCount, t)

	expectedIdentifiers := []struct {
		name string
	}{
		{"x"},
		{"foo"},
	}

	for i, expectedIdentifier := range expectedIdentifiers {
		statement := program.Statements[i]
		if !assertLetStatement(statement, expectedIdentifier.name, t) {
			return
		}
	}
}

// Helpers

func assertProgram(program *ast.Program, expectedStatementsCount int, t *testing.T) {
	if program == nil {
		t.Fatalf("Expected statements, got nil instead")
	}

	statementsCount := len(program.Statements)
	if statementsCount != expectedStatementsCount {
		t.Fatalf("Expected %d statements, got %d instead", expectedStatementsCount, statementsCount)
	}
}

func assertLetStatement(statement ast.Statement, expectedName string, t *testing.T) bool {
	if statement.TokenLiteral() != "let" {
		t.Errorf("Expected 'let' literal, got %q instead", statement.TokenLiteral())
	}

	letStatement, ok := statement.(*ast.LetStatement)
	if !ok {
		t.Errorf("Expected a LetStatement node, got %T instead", statement)
		return false
	}

	if letStatement.Name.Value != expectedName {
		t.Errorf("Expected identifier with name %s, got %s instead", expectedName, letStatement.Name.Value)
		return false
	}

	if letStatement.Name.TokenLiteral() != expectedName {
		t.Errorf("Expected  name %s, got %s instead", expectedName, letStatement.Name.TokenLiteral())
		return false
	}

	return true
}