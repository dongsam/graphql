package ast

import (
	"github.com/riverleo/graphql/language/source"
)

type Location struct {
	Start  int
	End    int
	Source *source.Source
}
