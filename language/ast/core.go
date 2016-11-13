package ast

import (
	"github.com/riverleo/graphql/language"
	"github.com/riverleo/graphql/language/kinds"
)

// Nodeable ...
type Nodeable interface {
	GetLoc() *Location
	GetKind() string
}

// Node ...
type Node struct {
	Loc  *Location
	Kind string
}

// GetLoc ...
func (n *Node) GetLoc() *Location {
	return n.Loc
}

// GetKind ...
func (n *Node) GetKind() string {
	return n.Kind
}

// Location ...
type Location struct {
	Start  int
	End    int
	Source *language.Source
}

// NameNode ...
type NameNode struct {
	Node
	Value string
}

// NewName ...
func NewName(n *NameNode) *NameNode {
	if n == nil {
		n = &NameNode{}
	}

	n.Kind = kinds.Name

	return n
}
