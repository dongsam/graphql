package ast

import (
	"github.com/riverleo/graphql/language/kinds"
)

type Node interface {
	getLoc() *Location
	getKind() string
}

type DocumentNode struct {
	Loc        Location
	Kind       string
	Operations []OperationNode
}

func NewDocument(n *DocumentNode) *DocumentNode {
	if n == nil {
		n = &DocumentNode{}
	}

	n.Kind = kinds.Document

	return n
}

type NameNode struct {
	Loc   Location
	Kind  string
	Value string
}

func NewName(n *NameNode) *NameNode {
	if n == nil {
		n = &NameNode{}
	}

	n.Kind = kinds.Name

	return n
}
