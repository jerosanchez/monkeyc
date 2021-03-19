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
		
	// Delimiters
	case ',':
		tokenFound = newToken(token.COMMA, self.currentChar)
	case ';':
		tokenFound = newToken(token.SEMICOLON, self.currentChar)
	case '(':
		tokenFound = newToken(token.LPAREN, self.currentChar)
	case ')':
		tokenFound = newToken(token.RPAREN, self.currentChar)
	case '{':
		tokenFound = newToken(token.LBRACE, self.currentChar)
	case '}':
		tokenFound = newToken(token.RBRACE, self.currentChar)
	
	// Control
	case 0:
		tokenFound.Type = token.EOF
		tokenFound.Literal = ""
		
	// Identifiers
	default:
		if isLetter(self.currentChar) {
			tokenFound.Type = token.IDENT
			tokenFound.Literal = self.readIdentifier()
			
		} else {
			tokenFound = newToken(token.ILLEGAL, self.currentChar)
		}
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

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 
		'A' <= char && char <= 'Z' || 
		char == '_'
}

func (self *Lexer) readIdentifier() string {
	initialPos := self.currentPos
	for isLetter(self.currentChar) {
		self.readChar()
	}
	return self.input[initialPos:self.currentPos]
}
