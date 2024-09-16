package engine

import (
	"math/big"
	"slices"
)

var infinityCharacters = []string{
	"*", "+",
}

func NewCharNode(regex string) (*CharNode, TreeError) {
	if len(regex) != 1 {
		return nil, ErrInvalidChar
	}
	if slices.Contains(infinityCharacters, regex) {
		return nil, ErrInfinity
	}
	return &CharNode{Node: Node{str: regex}}, nil
}

func (ch *CharNode) Calculate() *big.Int {
	return big.NewInt(1)
}
