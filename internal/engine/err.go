package engine

import "errors"

type TreeError error

var (
	ErrInfinity TreeError = errors.New("infinity possible variations error")

	ErrInvalidGroup TreeError = errors.New("invalid group")

	ErrInvalidCharset TreeError = errors.New("invalid charset")

	ErrInvalidQuantifier      TreeError = errors.New("invalid quantifier")
	ErrInvalidQuantifierRange TreeError = errors.New("invalid quantifier range")

	ErrInvalidChar TreeError = errors.New("invalid char")
)
