package validation

import "github.com/OlegSadJktu/regexer/internal/stack"

var brackets = []rune{'[', '{', '('}
var closingBrackets = []rune{']', '}', ')'}

func contains(arr []rune, elem rune) bool {
	for _, r := range arr {
		if r == elem {
			return true
		}
	}
	return false
}

func isClosing(open, close rune) bool {
	var index int
	for i, v := range brackets {
		if v == open {
			index = i
			break
		}
	}
	return closingBrackets[index] == close
}

func CheckBrackets(str string) bool {
	stack := stack.New[rune]()
	for _, r := range str {
		if contains(brackets, r) {
			stack = stack.Push(r)
		}
		if contains(closingBrackets, r) {
			var val rune
			stack, val = stack.Pop()
			if !isClosing(val, r) {
				return false
			}
		}

	}
	return len(stack) == 0
}
