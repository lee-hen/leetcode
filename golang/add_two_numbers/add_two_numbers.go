package add_two_numbers

type ListNode struct {
	 Val int
	 Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addLists(l1, l2, 0)
}

func addLists(l1, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	result := new(ListNode)
	value := carry

	if l1 != nil {
		value += l1.Val
	}

	if l2 != nil {
		value += l2.Val
	}

	result.Val = value % 10
	if l1 != nil || l2 != nil {
		var more, next1, next2 *ListNode

		if l1 != nil {
			next1 = l1.Next
		}

		if l2 != nil {
			next2 = l2.Next
		}

		if value >= 10 {
			more = addLists(next1, next2, 1)
		} else {
			more = addLists(next1, next2, 0)
		}

		result.Next = more
	}
	return result
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	var head = new(ListNode)
	l := head

	ptr1, ptr2 := l1, l2

	for ptr1 != nil || ptr2 != nil {
		actualValue := carry

		if ptr1 != nil {
			actualValue += ptr1.Val
		}

		if ptr2 != nil {
			actualValue += ptr2.Val
		}

		value := actualValue % 10

		if actualValue >= 10 {
			carry = 1
		} else {
			carry = 0
		}

		l.Next = &ListNode{Val: value}
		l = l.Next

		if ptr1 != nil {
			ptr1 = ptr1.Next
		}

		if ptr2 != nil {
			ptr2 = ptr2.Next
		}
	}

	if carry > 0 {
		l.Next =  &ListNode{Val: carry}
	}

	return head.Next
}
