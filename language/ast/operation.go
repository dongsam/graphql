package ast

import (
	"github.com/riverleo/graphql/language/kinds"
)

// OperationNode ...
type OperationNode struct {
	Node
	Name         NameNode
	Operation    string "mutation|query"
	Variables    []VariableNode
	SelectionSet SelectionSetNode
}

// NewOperation ...
func NewOperation(n *OperationNode) *OperationNode {
	if n == nil {
		n = &OperationNode{}
	}

	n.Kind = kinds.Operation

	return n
}

// SelectionSetNode ...
type SelectionSetNode struct {
	Node
	Selections []SelectionNode
}

func NewSelectionSet(n *SelectionSetNode) *SelectionSetNode {
	if n == nil {
		n = &SelectionSetNode{}
	}

	n.Kind = kinds.SelectionSet

	return n
}

// SelectionNode ...
type SelectionNode struct {
	Node
	name         NameNode
	alias        NameNode
	Arguments    []ArgumentNode
	SelectionSet SelectionSetNode
}

func NewSelection(n *SelectionNode) *SelectionNode {
	if n == nil {
		n = &SelectionNode{}
	}

	n.Kind = kinds.Selection

	return n
}

// ArgumentNode ...
type ArgumentNode struct {
	Node
	Name  NameNode
	Value ValueNode
}

// NewArgument ...
func NewArgument(n *ArgumentNode) *ArgumentNode {
	if n == nil {
		n = &ArgumentNode{}
	}

	n.Kind = kinds.Argument

	return n
}

// VariableNode ...
type VariableNode struct {
	Node
	Name NameNode
}

// NewVariable ...
func NewVariable(n *VariableNode) *VariableNode {
	if n == nil {
		n = &VariableNode{}
	}

	n.Kind = kinds.Variable

	return n
}
