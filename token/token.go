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
	MINUS = "-"
	ASTERISK = "*"
	SLASH = "/"
	BANG = "!"

	LT = "<"
	GT = ">"
	
	// Delimiters
	COMMA = ","
	SEMICOLON = ";"
	
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	
	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"

	// Control
	EOF = "EOF"
	ILLEGAL = "ILLEGAL"
)

func TypeFor(ident string) TokenType {
	if keywordType, ok := keywords[ident]; ok {
		return keywordType
	}
	
	return IDENT
}

var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
}

