// lexer/lexer_test.go

package lexer

import (
	"testing"
	"monkeyc/token"
)

func Test_NextToken_RecognizesOperators(t *testing.T) {
	input := `=+`
	
	tests := []token.Token {
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.EOF, ""},
	}
	
	lexer := New(input)
	
	for i, expectedToken := range tests {
		receivedToken := lexer.NextToken()
		
		if receivedToken.Type != expectedToken.Type {
			t.Fatalf("tests[%d] - wrong token type: expected %q, expected %q instead", i, expectedToken.Type, receivedToken.Type)
		}
		
		if receivedToken.Literal != expectedToken.Literal {
			t.Fatalf("tests[%d] - wrong token literal: expected %q, expected %q instead", i, expectedToken.Literal, receivedToken.Literal)
		}	
	}
}

func Test_NextToken_RecognizesDelimiters(t *testing.T) {
	input := `,;(){}`
	
	tests := []token.Token {
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}
	
	lexer := New(input)
	
	for i, expectedToken := range tests {
		receivedToken := lexer.NextToken()
		
		if receivedToken.Type != expectedToken.Type {
			t.Fatalf("tests[%d] - wrong token type: expected %q, expected %q instead", i, expectedToken.Type, receivedToken.Type)
		}
		
		if receivedToken.Literal != expectedToken.Literal {
			t.Fatalf("tests[%d] - wrong token literal: expected %q, expected %q instead", i, expectedToken.Literal, receivedToken.Literal)
		}	
	}
}

