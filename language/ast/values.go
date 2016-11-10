package ast

type ValueNode interface {
	GetLoc() *Location
	GetKind() string
	GetValue() interface{}
}

type IntValueNode struct {
	Loc   Location
	Kind  string
	Value string
}

type FloatValueNode struct {
	Loc   Location
	Kind  string
	Value string
}

type StringValueNode struct {
	Loc   Location
	Kind  string
	Value string
}

type BooleanValueNode struct {
	Loc   Location
	Kind  string
	Value string
}

type NullValueNode struct {
	Loc  Location
	Kind string
}

type EnumValueNode struct {
	Loc   Location
	Kind  string
	Value string
}

type ListValueNode struct {
	Loc   Location
	Kind  string
	Value []ValueNode
}

type ObjectValueNode struct {
	Loc   Location
	Kind  string
	Value []ObjectFieldNode
}

type ObjectFieldNode struct {
	Loc   Location
	Kind  string
	Name  NameNode
	Value ValueNode
}
