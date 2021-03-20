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
	ELSE = "ELSE"
	FALSE = "FALSE"
	FUNCTION = "FUNCTION"
	IF = "IF"
	LET = "LET"
	RETURN = "RETURN"
	TRUE = "TRUE"

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
	"else": ELSE,
	"false": FALSE,
	"fn": FUNCTION,
	"if": IF,
	"let": LET,
	"return": RETURN,
	"true": TRUE,
}

