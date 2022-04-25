package convert_binary_search_tree_to_sorted_doubly_linked_list

import (
	"github.com/lee-hen/leecode/lib"
)

type List struct {
	first *lib.TreeNode
	last *lib.TreeNode
}

func TreeToDoublyList(root *lib.TreeNode) *lib.TreeNode {
	if root == nil {
		return nil
	}

	list := new(List)
	helper(root, list)

	list.last.Right = list.first
	list.first.Left = list.last

	return list.first
}

func helper(node *lib.TreeNode, l *List) {
	if node != nil {
		helper(node.Left, l)
		if l.last != nil {
			l.last.Right = node
			node.Left =  l.last
		} else {
			l.first = node
		}
		l.last = node
		helper(node.Right, l)
	}
}

func treeToDoublyList(root *lib.TreeNode) *lib.TreeNode {
	if root == nil {
		return nil
	}

	left := leftMostChild(root)
	right := rightMostChild(root)
	InOrderTraversal(root)
	right.Right = left
	left.Left = right
	return left
}

func InOrderTraversal(node *lib.TreeNode) {
	if node == nil {
		return
	}

	InOrderTraversal(node.Left)
	right := node.Right

	if node.Left != nil  {
		rightMost := rightMostChild(node.Left)
		if node.Val > rightMost.Val {
			node.Left = rightMost
			rightMost.Right = node
		}
	}

	InOrderTraversal(right)
	if node.Right != nil {
		leftMost := leftMostChild(node.Right)
		if node.Val < leftMost.Val {
			node.Right = leftMost
			leftMost.Left = node
		}
	}
}

func leftMostChild(n *lib.TreeNode) *lib.TreeNode {
	if n == nil {
		return nil
	}

	for n.Left != nil {
		n = n.Left
	}
	return n
}

func rightMostChild(n *lib.TreeNode) *lib.TreeNode {
	if n == nil {
		return nil
	}

	for n.Right != nil {
		n = n.Right
	}
	return n
}
