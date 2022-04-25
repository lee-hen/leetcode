package remove_nth_node_from_end_of_list

type ListNode struct {
	Val int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	first := dummy
	second := dummy

	// Advances first pointer so that the gap between first and second is n nodes apart
	for i := 1; i <= n + 1; i++ {
		first = first.Next
	}

	// Move first to the end, maintaining the gap
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
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
