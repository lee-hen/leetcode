package max_stack

import (
	"container/heap"
	"container/list"
)

type Node struct {
	id, value int
}

type MaxStack struct {
	stack *list.List
	maxHeap *IntHeap

	id int
	stackPosition map[int]*list.Element
}

func Constructor() MaxStack {
	maxHeap := &IntHeap{}
	heap.Init(maxHeap)

	return MaxStack {
		maxHeap: maxHeap,
		stack: list.New(),
		id: 0,
		stackPosition: make(map[int]*list.Element),
	}
}

func (s *MaxStack) Top() int {
	return s.stack.Back().Value.(*Node).value
}

func (s *MaxStack) Push(x int)  {
	nodeWithId := &Node{s.id, x}
	heap.Push(s.maxHeap,nodeWithId)
	s.stack.PushBack(nodeWithId)
	s.stackPosition[s.id] = s.stack.Back()
	s.id++
}

func (s *MaxStack) Pop() int {
	top := s.stack.Back()
	s.stack.Remove(top)
	node := top.Value.(*Node)

	maxNode := (*s.maxHeap)[0]
	if maxNode.id == node.id {
		heap.Pop(s.maxHeap)
	}

	delete(s.stackPosition, node.id)
	return node.value
}

func (s *MaxStack) PopMax() int {
	maxNode := heap.Pop(s.maxHeap).(*Node)
	max, found := s.stackPosition[maxNode.id]

	for !found {
		maxNode = heap.Pop(s.maxHeap).(*Node)
		max, found = s.stackPosition[maxNode.id]
	}

	s.stack.Remove(max)
	delete(s.stackPosition, maxNode.id)
	return maxNode.value
}

func (s *MaxStack) PeekMax() int {
	maxNode := (*s.maxHeap)[0]
	_, found := s.stackPosition[maxNode.id]

	for !found {
		heap.Pop(s.maxHeap)
		maxNode = (*s.maxHeap)[0]
		_, found = s.stackPosition[maxNode.id]
	}

	return maxNode.value
}

type IntHeap []*Node

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return  h[j].value <= h[i].value }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
