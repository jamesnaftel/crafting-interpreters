package token

import "fmt"

type Token struct {
	type_   TokenType
	lexeme  string
	literal string // What should this be?
	line    int
}

func NewToken(type_ TokenType, lexeme string, literal string, line int) *Token {
	return &Token{type_, lexeme, literal, line}
}

func (t Token) String() string {
	return fmt.Sprintf("%s %s %s", t.type_, t.lexeme, t.literal)
}
