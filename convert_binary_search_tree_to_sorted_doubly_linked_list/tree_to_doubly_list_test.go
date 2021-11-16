package convert_binary_search_tree_to_sorted_doubly_linked_list

import (
	"github.com/lee-hen/leecode/lib"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	root := lib.NewTreeNode(30)
	root.InsertInOrder(13)
	root.InsertInOrder(-28)
	root.InsertInOrder(-44)
	root.InsertInOrder(-35)
	left := treeToDoublyList(root)
	right, leftToRight := getLeftToRight(left)
	_, rightToLeft := getRightToLeft(right)

	require.Equal(t, 5, len(leftToRight))
	require.Equal(t, 5, len(rightToLeft))

	for i, num := range leftToRight {
		require.Equal(t, num, rightToLeft[len(rightToLeft)-i-1])
	}
}

func TestTreeToDoublyList(t *testing.T) {
	for max := 1; max < 100; max++ {
		array := make([]int, 0)
		for i := -100; i < max; i++ {
			array = append(array, i)
		}
		bst := lib.CreateMinimalBst(array)
		left := TreeToDoublyList(bst)
		right, leftToRight := getLeftToRight(left)
		_, rightToLeft := getRightToLeft(right)

		require.Equal(t, len(array), len(leftToRight))
		require.Equal(t, len(array), len(rightToLeft))
		for i, num := range leftToRight {
			require.Equal(t, num, rightToLeft[len(rightToLeft)-i-1])
		}
	}
}

func getLeftToRight(left *lib.TreeNode) (*lib.TreeNode, []int) {
	leftToRight := make([]int, 0)
	curr := left
	leftToRight = append(leftToRight, curr.Val)
	for curr.Right != left {
		curr = curr.Right
		leftToRight = append(leftToRight, curr.Val)
	}
	return curr, leftToRight
}


func getRightToLeft(right *lib.TreeNode) (*lib.TreeNode, []int) {
	curr := right
	rightToLeft := make([]int, 0)
	rightToLeft = append(rightToLeft, curr.Val)
	for curr.Left != right {
		curr = curr.Left
		rightToLeft = append(rightToLeft, curr.Val)
	}
	return curr, rightToLeft
}