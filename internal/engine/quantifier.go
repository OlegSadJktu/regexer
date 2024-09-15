package engine

import (
	"math/big"
	"strconv"
	"strings"
)

type quantifierWrapper struct {
	result *big.Int
	child  Calculatable
}

var _ Calculatable = (*quantifierWrapper)(nil)

func (obj *quantifierWrapper) Calculate() *big.Int {
	return obj.result
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func NewQuantifierNode(regex string, parent *Node, index int) *QuantifierNode {
	if regex[0] != '{' || regex[len(regex)-1] != '}' {
		panic("not a quantifier")
	}
	node := &QuantifierNode{
		Node:   Node{str: regex},
		parent: parent,
		index:  index,
	}
	withoutBrackets := regex[1 : len(regex)-1]
	withComma := strings.Contains(withoutBrackets, ",")
	if withComma {
		numbers := strings.Split(withoutBrackets, ",")
		if len(numbers) != 2 {
			panic("invalid quantifier args")
		}
		var err error
		node.from, err = strconv.Atoi(numbers[0])
		check(err)
		node.to, err = strconv.Atoi(numbers[1])
		check(err)
	} else {
		fromTo, _ := strconv.Atoi(withoutBrackets)
		node.from = fromTo
		node.to = fromTo
	}
	return node
}

func (q *QuantifierNode) Calculate() *big.Int {
	return big.NewInt(1)
}

func (q *QuantifierNode) Prepare() {
	prevNode := &(q.parent.children[q.index])
	prevValue := (*prevNode).Calculate()
	var sum big.Int
	for i := q.from; i <= q.to; i++ {
		var pow big.Int
		pow.Exp(prevValue, big.NewInt(int64(i)), nil)
		sum.Add(&sum, &pow)
	}
	*prevNode = &quantifierWrapper{
		result: &sum,
		child:  *prevNode,
	}
}
