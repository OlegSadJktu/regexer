package engine

import "math/big"

func (node *StringNode) Calculate() *big.Int {
	sum := big.NewInt(1)
	for _, ch := range node.children {
    
		sum.Mul(sum, ch.Calculate())
	}
	return sum
}
