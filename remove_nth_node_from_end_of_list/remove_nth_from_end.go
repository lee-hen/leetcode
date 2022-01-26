package remove_nth_node_from_end_of_list

type ListNode struct {
	Val int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	if length == n {
		return head.Next
	}

	var i int
	curr := head
	for i < length-n-1 {
		curr = curr.Next
		i++
	}

	if curr.Next != nil {
		curr.Next = curr.Next.Next
	} else {
		curr.Next = nil
	}

	return head
}

func getLength(head *ListNode) int {
	length := 0
	for head != nil {
		head = head.Next
		length++
	}
	return length
}
