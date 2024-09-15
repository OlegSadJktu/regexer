package engine

import "math/big"

func Count(regex string) *big.Int {
  node := NewNode(regex)
  in := node.(Preparable)
  in.Prepare()
	return node.Calculate()
}