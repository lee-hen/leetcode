package add_two_numbers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := addMany([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := addMany([]int{9, 9, 9, 9})
	require.Equal(t, AddTwoNumbers(l1, l2), addTwoNumbers(l1, l2))
}

func addMany(values []int) *ListNode {
	listNode := new(ListNode)
	current := listNode
	for current.Next != nil {
		current = current.Next
	}
	for _, value := range values {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}
	return listNode
}
