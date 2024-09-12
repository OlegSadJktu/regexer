package engine

import "math/big"

func NewCharNode(regex string) *CharNode {
	if len(regex) != 1 {
		panic("not a char")
	}
	return &CharNode{Node: Node{str: regex}}
}

func (ch *CharNode) Calculate() *big.Int {
	return big.NewInt(1)
}
