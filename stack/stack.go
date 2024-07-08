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

func (stack *Stack[T]) Top() *T {
	if stack.top != nil {
		return &stack.top.element
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
	var oldTop = stack.top
	var oldSecond = stack.top.previous
	stack.top = oldSecond
	oldTop.previous = oldSecond.previous
	oldSecond.previous = oldTop
}
