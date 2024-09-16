package engine

import (
	"math/big"
	"strings"
)

func NewGroupNode(regex string) (*GroupNode, TreeError) {
	if len(regex) == 0 {
		return nil, ErrInvalidQuantifier
	}
	node := &GroupNode{Node{str: regex}}
	if regex[0] != '(' || regex[len(regex)-1] != ')' {
		return nil, ErrInvalidGroup
	}
	withoutBrackets := regex[1 : len(regex)-1]
	nodes := strings.Split(withoutBrackets, "|")
	for _, strNode := range nodes {
		new, err := NewNode(strNode)
		if err != nil {
			return nil, err
		}
		node.children = append(node.children, new)
	}
	return node, nil

}

func (gr *GroupNode) Calculate() *big.Int {
	sum := big.NewInt(0)
	for _, ch := range gr.children {
		sum.Add(sum, ch.Calculate())
	}
	return sum
}
