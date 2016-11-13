package ast

import (
	"github.com/riverleo/graphql/language/kinds"
)

// DocumentNode ...
type DocumentNode struct {
	Node
	Operations []OperationNode
}

// NewDocument ...
func NewDocument(n *DocumentNode) *DocumentNode {
	if n == nil {
		n = &DocumentNode{}
	}

	n.Kind = kinds.Document

	return n
}
