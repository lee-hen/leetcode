package merge_k_sorted_lists

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	lists := []*ListNode{CreateList([]int{1, 4, 5}), CreateList([]int{1, 3, 4}), CreateList([]int{2, 6})}

	ptr1 := CreateList([]int{1, 1, 2, 3, 4, 4, 5, 6})
	ptr2 := MergeKLists(lists)

	for ptr1 != nil {
		require.Equal(t, ptr1.Val, ptr2.Val)

		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	var expected *ListNode = nil
	require.Equal(t, expected, MergeKLists([]*ListNode{nil}))
}


func Test_mergeKLists(t *testing.T) {
	lists := []*ListNode{CreateList([]int{1, 4, 5}), CreateList([]int{1, 3, 4}), CreateList([]int{2, 6})}

	ptr1 := CreateList([]int{1, 1, 2, 3, 4, 4, 5, 6})
	ptr2 := mergeKLists(lists)

	for ptr1 != nil {
		require.Equal(t, ptr1.Val, ptr2.Val)

		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	var expected *ListNode = nil
	require.Equal(t, expected, mergeKLists([]*ListNode{nil}))
}

func CreateList(array []int) *ListNode {
	if len(array) == 0 {
		return nil
	}

	front := new(ListNode)
	front.Val = array[0]

	ptr := front
	for i := 1; i < len(array); i++ {
		node := new(ListNode)
		node.Val = array[i]
		ptr.Next = node
		ptr = ptr.Next
	}
	return front
}
