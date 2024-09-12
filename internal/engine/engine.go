package engine

import "math/big"

func Count(regex string) *big.Int {
  node := NewNode(regex)
	return node.Calculate()
}