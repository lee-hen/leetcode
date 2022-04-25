package diameter_of_binary_tree

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	root := NewBinaryTree(1)
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	require.Equal(t, 3, DiameterOfBinaryTree(root))

	root = NewBinaryTree(1)
	root.Left = &TreeNode{Val: 2}
	require.Equal(t,1, DiameterOfBinaryTree(root))
}

func NewBinaryTree(value int) *TreeNode {
	return &TreeNode{Val: value}
}
