// token/token.go

package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (	
	// Operators
	ASSIGN = "="
	PLUS = "+"
	
	// Control
	EOF = "EOF"
)
