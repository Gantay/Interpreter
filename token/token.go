package token

type TokenType string

type Token struct {
	Token   TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Idntifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	//Operators
	ASSING = "="
	PLUS   = "+"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
