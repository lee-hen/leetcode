package merge_two_sorted_lists

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func NewListNode(val int, others ...int) *ListNode {
	l := &ListNode{Val: val}
	current := l
	for _, other := range others {
		current.Next = &ListNode{Val: other}
		current = current.Next
	}
	return l
}

func (l *ListNode) ToArray() []int {
	values := make([]int, 0)
	current := l
	for current != nil {
		values = append(values, current.Val)
		current = current.Next
	}
	return values
}

func TestMergeTwoLists(t *testing.T) {
	list1 := NewListNode(-9, -7, -3, -3, -1, 2, 3)
	list2 := NewListNode(-7, -7, -6, -6, -5, -3, 2, 4)
	output := MergeTwoLists(list1, list2)
	expectedNodes := []int{-9, -7, -7, -7, -6, -6, -5, -3, -3, -3, -1, 2, 2, 3, 4}
	require.Equal(t, expectedNodes, output.ToArray())
}
