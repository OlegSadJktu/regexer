package engine

import (
	"math/big"
)

func NewCharNode(regex string) (*CharNode, TreeError) {
	if len(regex) != 1 {
		return nil, ErrInvalidChar
	}
	return &CharNode{Node: Node{str: regex}}, nil
}

func (ch *CharNode) Calculate() *big.Int {
	return big.NewInt(1)
}
