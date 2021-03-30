// parser/parser-return_test.go 

package parser

import (
	"testing"
	"monkeyc/ast"
	"monkeyc/lexer"
)

func TestReturnStamements(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 123456;
	`

	aLexer := lexer.New(input)
	sut := New(aLexer)

	program := sut.ParseProgram()
	checkParserErrors(sut, t)

	expectedStatementsCount := 3
	assertProgram(program, expectedStatementsCount, t)

	for _, statement := range program.Statements {
		assertReturnStatement(statement, t)
	}
}

func assertReturnStatement(statement ast.Statement, t *testing.T) {
	receivedStatement, ok := statement.(*ast.ReturnStatement)

	if !ok {
		t.Errorf("Expected a return statement, got %T instead", receivedStatement)
		return
	}

	if receivedStatement.TokenLiteral() != "return" {
		t.Errorf("Expected token literal 'return', got '%q' instead", receivedStatement.TokenLiteral())
	}
}
