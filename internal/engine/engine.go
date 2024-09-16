package engine

import (
	"errors"
	"math/big"

	"github.com/OlegSadJktu/regexer/internal/validation"
)

func Count(regex string) (*big.Int, TreeError) {
	isValid := validation.CheckBrackets(regex)
	if !isValid {
		return nil, errors.New("invalid brackets")
	}
	node, err := NewNode(regex)
	if err != nil {
		return nil, err
	}
	in := node.(Preparable)
	in.Prepare()
	return node.Calculate(), nil
}
