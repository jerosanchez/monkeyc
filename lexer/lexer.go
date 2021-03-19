// lexer/lexer.go

package lexer

import "monkeyc/token"

type Lexer struct {
	input string
	currentPos int
	readPos int
	currentChar byte
}

func New(input string) *Lexer {
	instance := &Lexer{ 
		input: input,
	}
	
	instance.readChar()
	
	return instance
}

func (self *Lexer) NextToken() token.Token {
	var tokenFound token.Token
	
	switch self.currentChar {
		
	// Operators
	case '=':
		tokenFound = newToken(token.ASSIGN, self.currentChar)
	case '+':
		tokenFound = newToken(token.PLUS, self.currentChar)
		
	// Control
	case 0:
		tokenFound.Type = token.EOF
		tokenFound.Literal = ""
	}
	
	self.readChar()
	
	return tokenFound
}

// Helpers

func (self *Lexer) readChar() {
	if self.readPos >= len(self.input) {
		self.currentChar = 0
	} else {
		self.currentChar = self.input[self.readPos]
	}
	
	self.currentPos = self.readPos
	self.readPos += 1
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type:tokenType, Literal: string(char)}
}
