package merge_two_sorted_lists

type ListNode struct {
	Val int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val > list2.Val {
		return MergeTwoLists(list2, list1)
	}

	ptr1, ptr2 := list1, list2
	next := ptr1

	for ptr1 != nil && ptr2 != nil {
		if ptr1.Next != nil && ptr2.Val > ptr1.Next.Val {
			ptr1 = ptr1.Next
			continue
		}

		if ptr1.Val <= ptr2.Val {
			next = ptr1.Next
			ptr1.Next = ptr2
			ptr1 = next
		}
		ptr1, ptr2 = ptr2, ptr1

	}
	return list1
}
