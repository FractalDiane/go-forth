package stack

type node[T any] struct {
	element  T
	previous *node[T]
}

type Stack[T any] struct {
	top  *node[T]
	size uint
}

func (stack *Stack[T]) Push(element T) {
	var newNode = new(node[T])
	newNode.element = element
	newNode.previous = stack.top
	stack.top = newNode

	stack.size++
}

func (stack *Stack[T]) Size() uint {
	return stack.size
}

func (stack *Stack[T]) IsEmpty() bool {
	return stack.size == 0
}

func (stack *Stack[T]) Top() *T {
	if stack.top != nil {
		return &stack.top.element
	} else {
		return nil
	}
}

func (stack *Stack[T]) Second() *T {
	if stack.top != nil && stack.top.previous != nil {
		return &stack.top.previous.element
	} else {
		return nil
	}
}

func (stack *Stack[T]) Pop() {
	if stack.top != nil {
		stack.top = stack.top.previous
		stack.size--
	}
}

func (stack *Stack[T]) Array() []T {
	var result = make([]T, stack.size)
	var index = 0
	for stackNode := stack.top; stackNode != nil; stackNode = stackNode.previous {
		result[index] = stackNode.element
		index++
	}

	return result
}

func (stack *Stack[T]) SwapTopElements() {
	if stack.size >= 2 {
		var oldTop = stack.top
		var oldSecond = stack.top.previous
		var oldThird = stack.top.previous.previous
		stack.top = oldSecond
		stack.top.previous = oldTop
		stack.top.previous.previous = oldThird
	}
}

func (stack *Stack[T]) RotateTopElements() {
	if stack.size >= 3 {
		var oldTop = stack.top
		var oldSecond = stack.top.previous
		var oldThird = stack.top.previous.previous
		var oldFourth = stack.top.previous.previous.previous

		stack.top = oldThird
		stack.top.previous = oldSecond
		stack.top.previous.previous = oldTop
		stack.top.previous.previous.previous = oldFourth
	}
}
