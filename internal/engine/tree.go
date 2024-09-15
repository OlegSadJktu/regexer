package engine

import (
	"math/big"
	"strings"
)

type (
	Calculatable interface{
		Calculate() *big.Int
	}
	Preparable interface{
		Prepare()
	}
	Node struct {
		children []Calculatable
		str      string
	}
)

var (
	_ Calculatable = (*CompositeNode)(nil)
	_ Calculatable = (*GroupNode)(nil)
	_ Calculatable = (*CharNode)(nil)

	// _ Calculatable = (*CharsetNode)(nil)
	_ Calculatable = (*QuantifierNode)(nil)
	_ Preparable = (*QuantifierNode)(nil)
)


type (
	CompositeNode  struct{ Node } // combined type: [1-9]{1,3}
	CharNode       struct{ Node } // 'a'
	GroupNode      struct{ Node } // (asd|dsa)
	QuantifierNode struct {
		Node
		parent *Node
		index int
		from, to int
	} // {1,3}
	CharsetNode struct{ Node } // [1-9]
)

func NewNode(regex string) Calculatable {
	node := &CompositeNode{
		Node: Node{
			str:      regex,
			children: make([]Calculatable, 0),
		},
	}
	var newNode Calculatable
	runes := []rune(regex)
	for i := 0; i < len(runes); i++ {
		ch := runes[i]

		switch ch {
		case '(':
			groupIndex := strings.Index(regex[i:], ")") + 1
			newNode = NewGroupNode(regex[i : groupIndex+i])
			i = i + groupIndex - 1
		case '{':
			quantifierIndex := strings.Index(regex[i:], "}") + 1
			newNode = NewQuantifierNode(regex[i:quantifierIndex+i], &node.Node, len(node.children) -1)
			i = i + quantifierIndex - 1
		default:
			newNode = NewCharNode(string(ch))
		}
		node.children = append(node.children, newNode)
	}

	return node
}
