package copy_list_with_random_pointer

type Node struct {
	Val int
	Next *Node
	Random *Node
}

func CopyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// Creating a new weaved list of original and copied nodes.
	ptr := head
	for ptr != nil {

		// Cloned node
		newNode := &Node{Val: ptr.Val}

		// Inserting the cloned node just next to the original node.
		// If A->B->C is the original linked list,
		// Linked list after weaving cloned nodes would be A->A'->B->B'->C->C'
		newNode.Next = ptr.Next
		ptr.Next = newNode
		ptr = newNode.Next
	}

	ptr = head

	// Now link the random pointers of the new nodes created.
	// Iterate the newly created list and use the original nodes' random pointers,
	// to assign references to random pointers for cloned nodes.
	for ptr != nil {
		if ptr.Random != nil {
			ptr.Next.Random = ptr.Random.Next
		} else {
			ptr.Next.Random = nil
		}

		ptr = ptr.Next.Next
	}

	// Unweave the linked list to get back the original linked list and the cloned list.
	// i.e. A->A'->B->B'->C->C' would be broken to A->B->C and A'->B'->C'
	ptrOldList := head // A->B->C
	ptrNewList := head.Next // A'->B'->C'
	headOld := head.Next
	for ptrOldList != nil {
		ptrOldList.Next = ptrOldList.Next.Next
		if ptrNewList.Next != nil {
			ptrNewList.Next = ptrNewList.Next.Next
		} else {
			ptrNewList.Next = nil
		}

		ptrOldList = ptrOldList.Next
		ptrNewList = ptrNewList.Next
	}
	return headOld
}

func copyRandomList(head *Node) *Node {
	return deepCopy(head)
}

func buildNodeLookup(head *Node) map[*Node]int {
	lookup := make(map[*Node]int)

	for i := 0; head != nil; i++ {
		lookup[head] = i
		head = head.Next
	}

	return lookup
}

func deepCopy(head *Node) *Node {
	lookup := buildNodeLookup(head)
	curr := head

	node := new(Node)
	cpyNode := node

	reverseIndices := make(map[int]*Node)
	for i := 0; curr != nil; i++ {
		var randomNode, cpy *Node

		if curr.Random != nil {
			idx := lookup[curr.Random]
			if _, ok := reverseIndices[idx]; ok {
				randomNode = reverseIndices[idx]
			} else {
				randomNode = new(Node)
				randomNode.Val = curr.Random.Val
				reverseIndices[idx] = randomNode
			}
		} else {
			randomNode = nil
		}

		if _, ok := reverseIndices[i]; ok {
			cpy = reverseIndices[i]
		} else {
			cpy = new(Node)
			cpy.Val = curr.Val
		}

		cpy.Random = randomNode
		node.Next = cpy
		node = node.Next
		curr = curr.Next
		reverseIndices[i] = cpy
	}


	//curr = head
	//node = cpyNode.Next
	//for i := 0; curr != nil; i++ {
	//	if curr.Random != nil {
	//		idx := lookup[curr.Random]
	//		if randomNode, ok := reverseIndices[idx]; ok {
	//			node.Random = randomNode
	//		}
	//	}
	//	node = node.Next
	//	curr = curr.Next
	//}

	return cpyNode.Next
}

