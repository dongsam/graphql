package parser

import (
	"github.com/riverleo/graphql/language/ast"
	"github.com/riverleo/graphql/language/source"
)

func Parse(s *source.Source) *ast.Document {
	return NewDocument(nil)
}
