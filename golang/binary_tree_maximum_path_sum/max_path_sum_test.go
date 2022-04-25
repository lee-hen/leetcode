package binary_tree_maximum_path_sum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	root := NewBinaryTree(-10)
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	require.Equal(t, 42, MaxPathSum(root))

	root = NewBinaryTree(1)
	root.Left = &TreeNode{Val: 2}
	require.Equal(t,3, MaxPathSum(root))

	root = NewBinaryTree(-2)
	root.Left = &TreeNode{Val: -1}
	root.Right = &TreeNode{Val: -3}
	require.Equal(t,-1, MaxPathSum(root))

	root = NewBinaryTree(2)
	root.Left = &TreeNode{Val: -1}
	root.Right = &TreeNode{Val: -2}
	require.Equal(t,2, MaxPathSum(root))

	root = NewBinaryTree(9)
	root.Left = &TreeNode{Val: 6}
	root.Right = &TreeNode{Val: -3}
	root.Right.Left = &TreeNode{Val: -6}
	root.Right.Right = &TreeNode{Val: 2}
	root.Right.Right.Left = &TreeNode{Val: 2}
	root.Right.Right.Left.Left = &TreeNode{Val: -6}
	root.Right.Right.Left.Right = &TreeNode{Val: -6}
	root.Right.Right.Left.Left.Left = &TreeNode{Val: -6}
	require.Equal(t,16, MaxPathSum(root))
}

func NewBinaryTree(value int) *TreeNode {
	return &TreeNode{Val: value}
}
