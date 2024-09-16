package engine

import "math/big"

func (node *CompositeNode) Calculate() *big.Int {
	sum := big.NewInt(1)
	for _, ch := range node.children {

		sum.Mul(sum, ch.Calculate())
	}
	return sum
}

func (node *CompositeNode) Prepare() {
	for _, ch := range node.children {
		in, ok := ch.(Preparable)
		if !ok {
			continue
		}
		in.Prepare()
	}

}
