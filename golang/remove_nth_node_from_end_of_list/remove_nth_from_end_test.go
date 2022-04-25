package remove_nth_node_from_end_of_list

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	require.Equal(t, removeNthFromEnd(addMany([]int{1, 2, 3, 4, 5}), 2), RemoveNthFromEnd(addMany([]int{1, 2, 3, 4, 5}), 2))
	require.Equal(t, removeNthFromEnd(addMany([]int{1}), 1), RemoveNthFromEnd(addMany([]int{1}), 1))
	require.Equal(t, removeNthFromEnd(addMany([]int{1, 2}), 1), RemoveNthFromEnd(addMany([]int{1, 2}), 1))
	require.Equal(t, removeNthFromEnd(addMany([]int{1, 2}), 2), RemoveNthFromEnd(addMany([]int{1, 2}), 2))
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
	return listNode.Next
}



