package ast

import (
	"github.com/riverleo/graphql/language/kinds"
)

type OperationNode struct {
	Loc                 Location
	Kind                string
	Name                NameNode
	Operation           string
	VariableDefinitions []VariableDefinitionNode
	SelectionSet        SelectionSetNode
}

func NewOperation(n *OperationNode) *OperationNode {
	if n == nil {
		n = &OperationNode{}
	}

	n.Kind = kinds.Operation

	return n
}

type SelectionSetNode struct {
	Loc        Location
	Kind       string
	Selections []SelectionNode
}

type SelectionNode struct {
	Loc          Location
	Kind         string
	name         NameNode
	alias        NameNode
	Arguments    []ArgumentNode
	SelectionSet SelectionSetNode
}

type ArgumentNode struct {
	Loc   Location
	Kind  string
	Name  NameNode
	Value ValueNode
}

func NewArgument(n *ArgumentNode) *ArgumentNode {
	if n == nil {
		n = &ArgumentNode{}
	}

	n.Kind = kinds.Argument

	return n
}

type VariableDefinitionNode struct {
	Loc          Location
	Kind         string
	Type         TypeNode
	Variable     VariableNode
	DefaultValue ValueNode
}

type VariableNode struct {
	Loc  Location
	Kind string
	Name NameNode
}

func NewVariable(d *VariableNode) *VariableNode {
	if n == nil {
		n = &VariableNode{}
	}

	n.Kind = kinds.Variable

	return n
}
