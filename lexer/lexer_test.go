// lexer/lexer_test.go

package lexer

import (
	"testing"
	"monkeyc/token"
)

func Test_NextToken_RecognizesOperators(t *testing.T) {
	input := `=+`
	
	expectedTokens := []token.Token {
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.EOF, ""},
	}
	
	assertProducing(input, expectedTokens, t)	
}

func Test_NextToken_RecognizesDelimiters(t *testing.T) {
	input := `,;(){}`
	
	expectedTokens := []token.Token {
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}
			
	assertProducing(input, expectedTokens, t)	
}

func Test_NextToken_RecognizesIdentifiers(t *testing.T) {
	input := `an_Identifier`
	
	expectedTokens := []token.Token {
		{token.IDENT, "an_Identifier"},
		{token.EOF, ""},
	}
			
	assertProducing(input, expectedTokens, t)	
}

func Test_NextToken_RecognizesIntegers(t *testing.T) {
	input := `0123456789`
	
	expectedTokens := []token.Token {
		{token.INT, "0123456789"},
		{token.EOF, ""},
	}
			
	assertProducing(input, expectedTokens, t)	
}

// Helpers

func assertProducing(input string, expectedTokens []token.Token, t *testing.T) {
	lexer := New(input)

	for i, expectedToken := range expectedTokens {
		receivedToken := lexer.NextToken()
		
		if receivedToken.Type != expectedToken.Type {
			t.Fatalf("tests[%d] - wrong token type: expected %q, got %q instead", i, expectedToken.Type, receivedToken.Type)
		}
		
		if receivedToken.Literal != expectedToken.Literal {
			t.Fatalf("tests[%d] - wrong token literal: expected %q, got %q instead", i, expectedToken.Literal, receivedToken.Literal)
		}	
	}
}
