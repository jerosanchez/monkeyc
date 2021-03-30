// parser/parser.go (return)

package parser

import (
	"monkeyc/ast"
	"monkeyc/token"
)

func (self *Parser) parseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{
		Token: self.currentToken,
	}

	self.readToken()

	// TODO: parse expression;
	// meanwhile, skip everithing up to a SEMICOLON
	for !self.currentTokenIs(token.SEMICOLON) {
		self.readToken()
	}

	return statement
}