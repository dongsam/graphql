package lexer

import (
	"github.com/riverleo/graphql/language/source"
	"testing"
)

func TestLex(t *testing.T) {
	body := `
		{
			graphql: go
		}
	`

	lexer := Lex(&source.Source{
		Body: body,
	})

	lexer(1)
}
