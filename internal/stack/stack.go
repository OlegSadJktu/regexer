package stack

type Stack[T any] []T

func New[T any]() Stack[T] {
	return make(Stack[T], 0)
}

func (st Stack[T]) Push(v T) Stack[T] {
	return append(st, v)
}

func (st Stack[T]) Pop() (Stack[T], T) {
	l := len(st)
	return st[:l-1], st[l-1]
}
