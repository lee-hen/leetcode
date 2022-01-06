package merge_k_sorted_lists

import (
	"container/heap"
)

// ListNode
// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode {
	merge2Lists := func(l1, l2 *ListNode) *ListNode {
		head := new(ListNode)
		point := head
		for l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				point.Next = l1
				l1 = l1.Next
			} else {
				point.Next = l2
				l2 = l1
				l1 = point.Next.Next

			}
			point = point.Next
		}

		if l1 == nil {
			point.Next = l2
		} else if l2 == nil {
			point.Next = l1
		}

		return head.Next
	}

	interval := 1
	for interval < len(lists) {
		for i := 0; i < len(lists)-interval; i += interval * 2 {
			lists[i] = merge2Lists(lists[i], lists[i + interval])
		}
		interval *= 2
	}

	if len(lists) > 0 {
		return lists[0]
	}

	return nil
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &MinHeap{}
	for _, list := range lists {
		if list != nil { // []*ListNode{nil}
			*h = append(*h, list)
		}
	}
	heap.Init(h)

	l := new(ListNode)
	ptr := l
	for len(*h) > 0 {
		top := heap.Pop(h).(*ListNode)

		ptr.Next = top
		ptr = ptr.Next

		if top.Next != nil {
			heap.Push(h, top.Next)
		}
	}

	return l.Next
}

// An MinHeap is a min-heap of ints.
type MinHeap []*ListNode

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
