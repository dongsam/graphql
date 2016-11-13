package ast

import (
	"github.com/riverleo/graphql/language/kinds"
	"strconv"
)

// Valuable ...
type Valuable interface {
	GetType() string
	GetValue() *interface{}
	ToInt() int
	ToFloat() float32
	ToBool() bool
	ToString() string
	ToList() []ValueNode
	ToMap() []ObjectFieldNode
}

// ValueNode ...
type ValueNode struct {
	Node
	Type  string "null|int|float|bool|string|list|map"
	Value interface{}
}

func (n *ValueNode) GetValue() *interface{} {
	return &n.Value
}

func (n *ValueNode) ToInt() int {
	v, err := strconv.ParseInt(n.Value.(string), 10, 32)

	if err != nil {
		panic(err)
	}

	return int(v)
}

func (n *ValueNode) ToFloat() float32 {
	v, err := strconv.ParseFloat(n.Value.(string), 32)

	if err != nil {
		panic(err)
	}

	return float32(v)
}

func (n *ValueNode) ToString() string {
	return n.Value.(string)
}

func (n *ValueNode) ToList() []ValueNode {
	return n.Value.([]ValueNode)
}

func (n *ValueNode) ToMap() []ObjectFieldNode {
	return n.Value.([]ObjectFieldNode)
}

// NewValue ...
func NewValue(n *ValueNode) *ValueNode {
	if n == nil {
		n = &ValueNode{}
	}

	n.Kind = kinds.Value

	return n
}

// ObjectFieldNode ...
type ObjectFieldNode struct {
	Node
	Name  NameNode
	Value ValueNode
}

// NewObjectFieldNode ...
func NewObjectFieldNode(n *ObjectFieldNode) *ObjectFieldNode {
	if n == nil {
		n = &ObjectFieldNode{}
	}

	n.Kind = kinds.ObjectFieldNode

	return n
}
