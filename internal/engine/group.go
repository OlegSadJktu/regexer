package engine

import (
	"math/big"
	"strings"
)

func NewGroupNode(regex string) *GroupNode {
	node := &GroupNode{Node{str: regex}}
	if regex[0] != '(' || regex[len(regex)-1] != ')' {
		panic("not a group")
	}
	withoutBrackets := regex[1 : len(regex)-1]
	nodes := strings.Split(withoutBrackets, "|")
	for _, strNode := range nodes {
		node.children = append(node.children, NewNode(strNode))
	}
	return node

}

func (gr *GroupNode) Calculate() *big.Int {
	sum := big.NewInt(0)
	for _, ch := range gr.children {
		sum.Add(sum, ch.Calculate())
	}
	return sum
}
