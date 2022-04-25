package min_stack

import (
	"container/list"
	"math"
)

type NodeWithMin struct {
	value, min int
}

type MinStack struct {
	stack *list.List
}

func Constructor() MinStack {
	return MinStack {
		stack: list.New(),
	}
}

func (s *MinStack) Push(x int)  {
	newMin := Min(x, s.GetMin())
	s.stack.PushBack(&NodeWithMin{x, newMin})
}

func (s *MinStack) Pop()  {
	s.stack.Remove(s.stack.Back())
}

func (s *MinStack) Top() int {
	return s.stack.Back().Value.(*NodeWithMin).value
}

func (s *MinStack) GetMin() int {
	if s.stack.Len() == 0 {
		return math.MaxInt32
	} else {
		return s.stack.Back().Value.(*NodeWithMin).min
	}
}

func Min(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if curr > num {
			curr = num
		}
	}
	return curr
}
