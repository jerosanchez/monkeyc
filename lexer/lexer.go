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
	
	self.skipWhitespace()
	
	switch self.currentChar {
	
	// Operators
	case '=':
		if self.peekChar() == '=' {
			self.readChar()
			tokenFound = token.Token{token.EQ, "=="}
		} else {
			tokenFound = newToken(token.ASSIGN, self.currentChar)
		}
	
	case '+':
		tokenFound = newToken(token.PLUS, self.currentChar)
	
	case '-':
		tokenFound = newToken(token.MINUS, self.currentChar)
	
	case '*':
		tokenFound = newToken(token.ASTERISK, self.currentChar)
	
	case '/':
		tokenFound = newToken(token.SLASH, self.currentChar)
	
	case '!':
		if self.peekChar() == '=' {
			self.readChar()
			tokenFound = token.Token{token.NEQ, "!="}
		} else {
			tokenFound = newToken(token.BANG, self.currentChar)
		}
	
	case '<':
		tokenFound = newToken(token.LT, self.currentChar)
	
	case '>':
		tokenFound = newToken(token.GT, self.currentChar)	
		
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
			tokenFound.Literal = self.readIdentifier()
			tokenFound.Type = token.TypeFor(tokenFound.Literal)
			return tokenFound
			
		} else if isDigit(self.currentChar) {
			tokenFound.Type = token.INT
			tokenFound.Literal = self.readNumber()
			return tokenFound
			
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

func (self *Lexer) peekChar() byte {
	if self.readPos >= len(self.input) {
		return 0
	} else {
		return self.input[self.readPos]
	}
}

func (self *Lexer) skipWhitespace() {
	for self.currentChar == ' ' || self.currentChar == '\t' || self.currentChar == '\n' || self.currentChar == '\r' {
		self.readChar()
	}
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

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (self *Lexer) readNumber() string {
	initialPos := self.currentPos
	for isDigit(self.currentChar) {
		self.readChar()
	}
	return self.input[initialPos:self.currentPos]
}
