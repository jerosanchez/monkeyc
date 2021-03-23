// ast/ast.go

package ast

import (
	"monkeyc/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (self *Program) TokenLiteral() string {
	if len(self.Statements) > 0 {
		return self.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// let

type LetStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

func (self *LetStatement) statementNode() { }

func (self *LetStatement) TokenLiteral() string { 
	return self.Token.Literal
}

// identifier

type Identifier struct {
	Token token.Token

	// Value is not really needed; included to keep things simple,
	// since most other expression have a value
	Value string
}

func (self *Identifier) expressionNode() { }

func (self *Identifier) TokenLiteral() string {
	return self.Token.Literal
}