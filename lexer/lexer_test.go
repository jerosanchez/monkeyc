// lexer/lexer_test.go

package lexer

import (
	"testing"
	"monkeyc/token"
)

func Test_NextToken_RecognizesOperators(t *testing.T) {
	input := `=+-*/!<>`
	
	expectedTokens := []token.Token {
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.ASTERISK, "*"},
		{token.SLASH, "/"},
		{token.BANG, "!"},
		{token.LT, "<"},
		{token.GT, ">"},
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

func Test_NextToken_RecognizesKeywords(t *testing.T) {
	input := `else false fn if let return true`
	
	expectedTokens := []token.Token {
		{token.ELSE, "else"},
		{token.FALSE, "false"},
		{token.FUNCTION, "fn"},
		{token.IF, "if"},
		{token.LET, "let"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
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

func Test_NextToken_RecognizesASimpleProgram(t *testing.T) {
	input := `
		let five = 5;
		let ten = 10;
			
		let add = fn(x, y) {
			x + y;
		};
			
		let result = add(five, ten);
	`
	
	expectedTokens := []token.Token {
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
	
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		
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
