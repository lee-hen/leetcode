package count_complete_tree_nodes

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCountNodes(t *testing.T) {
	root := NewBinaryTree(1)
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	require.Equal(t, countNodes(root), CountNodes(root))
}

func NewBinaryTree(value int) *TreeNode {
	return &TreeNode{Val: value}
}
