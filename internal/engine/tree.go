package engine

import (
	"math/big"
)

type Calculatable interface {
	Calculate() *big.Int
}

var _ Calculatable = (*CharNode)(nil)
var _ Calculatable = (*Node)(nil)

func (node *Node) Calculate() *big.Int {
	sum := big.NewInt(0)
	for _, ch := range node.children {
    sum.Add(sum, ch.Calculate())
	}
	return sum
}

type Node struct {
	children []*Node
	str      string
}

type (
	CharNode       struct{ Node }
	GroupNode      struct{ Node }
	QuantifierNode struct{ Node }
	CharsetNode    struct{ Node }
)

func FromRegex(regex string) *Node {
	node := &Node{str: regex}
  return node
}

func (ch *CharNode) Calculate() *big.Int {
	return big.NewInt(1)
}
