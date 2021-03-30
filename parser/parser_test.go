// parser/parser_test.go

package parser

import (
	"testing"
	"monkeyc/ast"
	"monkeyc/lexer"
)

func TestParserReportsErrors(t *testing.T) {
	input := `
	let 123 = 5;
	let = ;
	let foo 5;
	`

	aLexer := lexer.New(input)
	sut := New(aLexer)

	_ = sut.ParseProgram()

	errorsCount := len(sut.Errors())
	if errorsCount != 3 {
		t.Errorf("Expected %d parsing errors, got %d instead:", 3, errorsCount)
		for _, message := range sut.Errors() {
			t.Errorf("   > %q", message)
		}
		t.Fatalf("")
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

func checkParserErrors(theParser *Parser, t *testing.T) {
	errors := theParser.Errors()

	if len(errors) == 0 {
		return 
	}

	t.Errorf("Parser found %d errors:", len(errors))
	for _, message := range errors {
		t.Errorf("   > %q", message)
	}

	t.FailNow()
}
