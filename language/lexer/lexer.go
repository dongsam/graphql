package lexer

import (
	"github.com/riverleo/graphql/language/source"
	"log"
)

const (
	SOF       = "<SOF>"
	EOF       = "<EOF>"
	BANG      = "!"
	DOLLAR    = "$"
	PAREN_L   = "("
	PAREN_R   = ")"
	SPREAD    = "..."
	COLON     = ":"
	EQUALS    = "="
	AT        = "@"
	BRACKET_L = "["
	BRACKET_R = "]"
	BRACE_L   = "{"
	PIPE      = "|"
	BRACE_R   = "}"
	NAME      = "Name"
	INT       = "Int"
	FLOAT     = "Float"
	STRING    = "String"
	COMMENT   = "Comment"
)

type Token struct {
	Kind   int
	Start  int
	End    int
	Line   int
	Column int
	Prev   *Token
	Value  string
}

type Lexer func(resetPosition int) (Token, error)

func Lex(s *source.Source) Lexer {
	log.Println(s.Body)
	return func(resetPosition int) (Token, error) {
		return Token{}, nil
	}
}
