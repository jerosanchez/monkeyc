// parser/parser.go

package parser

import (
	"monkeyc/ast"
	"monkeyc/lexer"
	"monkeyc/token"
)

type Parser struct {
	theLexer *lexer.Lexer

	currentToken token.Token
	nextToken token.Token
}

func New(aLexer *lexer.Lexer) *Parser {
	instance := &Parser{theLexer: aLexer}

	instance.readToken()
	instance.readToken()

	return instance
}

func (self *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	for self.currentToken.Type != token.EOF {
		statement := self.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}

		self.readToken()
	}

	return program
}

// Helpers

func (self *Parser) readToken() {
	self.currentToken = self.nextToken
	self.nextToken = self.theLexer.NextToken()
}

func (self *Parser) nextTokenIs(expectedType token.TokenType) bool {
	return self.nextToken.Type == expectedType
}

func (self *Parser) currentTokenIs(expectedType token.TokenType) bool {
	return self.currentToken.Type == expectedType
}

func (self *Parser) consumeToken(expectedType token.TokenType) bool {
	if self.nextTokenIs(expectedType) {
		self.readToken()
		return true
	} else {
		return false
	}
}

func (self *Parser) parseStatement() ast.Statement {
	switch self.currentToken.Type {
	case token.LET:
		return self.parseLetStatement()
	default:
		return nil
	}
}
