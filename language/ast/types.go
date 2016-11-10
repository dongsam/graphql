package ast

type TypeNode interface {
	getLoc() *Location
	getKind() string
}
