package engine

import (
	"math/big"
	"strings"
)

type Calculatable interface {
	Calculate() *big.Int
}

type Node struct {
	children []Calculatable
	str      string
}

var _ Calculatable = (*StringNode)(nil)
var _ Calculatable = (*GroupNode)(nil)
var _ Calculatable = (*CharNode)(nil)
// var _ Calculatable = (*CharsetNode)(nil)
// var _ Calculatable = (*QuantifierNode)(nil)

type (
	StringNode     struct{ Node }
	CharNode       struct{ Node }
	GroupNode      struct{ Node }
	QuantifierNode struct{ Node }
	CharsetNode    struct{ Node }
)

func NewNode(regex string) Calculatable {
	node := &StringNode{
		Node: Node{
			str:      regex,
			children: make([]Calculatable, 0),
		},
	}
	for i, ch := range regex {
		switch ch {
		case '(':
			groupIndex := strings.Index(regex[i:], "(") + 1
			NewGroupNode(regex[i:groupIndex])

		}
	}

	return node
}
