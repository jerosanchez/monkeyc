// parser/parser-let.go

package parser

import (
	"monkeyc/ast"
	"monkeyc/token"
)

func (self *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{
		Token: self.currentToken,
	}

	if !self.consumeToken(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{
		Token: self.currentToken,
		Value: self.currentToken.Literal,
	}

	if !self.consumeToken(token.ASSIGN) {
		return nil
	}

	// TODO: parse expression;
	// meanwhile, skip everithing up to a SEMICOLON
	for !self.currentTokenIs(token.SEMICOLON) {
		self.readToken()
	}

	return statement
}