// token/token.go

package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (	
	// Identifiers
	IDENT = "IDENT"
	
	// Literals
	INT = "INT"

	// Operators
	ASSIGN = "="
	PLUS = "+"
	
	// Delimiters
	COMMA = ","
	SEMICOLON = ";"
	
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	
	// Control
	EOF = "EOF"
	ILLEGAL = "ILLEGAL"
)
