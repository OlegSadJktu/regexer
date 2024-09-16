package engine

import (
	"math/big"
)

func CalculateRange(regex string) *big.Int {
	str := []rune(regex)
	var sum int64 = 0
	var i int
	for i = 0; i < len(str); i++ {
		r := str[i]
		if r == '-' {
			_range := str[i+1] - str[i-1]
			if _range < 0 {
				panic("invalid range")
			}
			sum += int64(_range) - 1
			i += 2
		} else {
			sum += 1
		}
	}
	return big.NewInt(sum)
}

func NewCharsetNode(regex string) (Calculatable, error) {
	if len(regex) == 0 || regex[0] != '[' || regex[len(regex)-1] != ']' {
		return nil, ErrInvalidCharset
	}
	node := &CharsetNode{
		Node: Node{
			str: regex,
		},
	}
	return node, nil
}

func (n *CharsetNode) Calculate() *big.Int {
	return CalculateRange(n.str)
}
