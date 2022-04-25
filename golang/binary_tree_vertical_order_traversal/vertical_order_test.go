package binary_tree_vertical_order_traversal

import (
	"github.com/lee-hen/leecode/lib"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	root := lib.NewTreeNode(3)
	root.Left = lib.NewTreeNode(9)
	root.Right = lib.NewTreeNode(20)
	root.Right.Left = lib.NewTreeNode(15)
	root.Right.Right = lib.NewTreeNode(7)

	require.Equal(t, [][]int{{9},{3,15},{20},{7}}, VerticalOrder(root))
}

func TestCase2(t *testing.T) {
	root := lib.NewTreeNode(3)
	root.Left = lib.NewTreeNode(9)
	root.Right = lib.NewTreeNode(8)

	root.Left.Left = lib.NewTreeNode(4)
	root.Left.Right = lib.NewTreeNode(0)

	root.Right.Left = lib.NewTreeNode(1)
	root.Right.Right = lib.NewTreeNode(7)

	require.Equal(t, [][]int{{4},{9},{3,0,1},{8},{7}}, VerticalOrder(root))
}

func TestCase3(t *testing.T) {
	root := lib.NewTreeNode(3)
	root.Left = lib.NewTreeNode(9)
	root.Right = lib.NewTreeNode(8)

	root.Left.Left = lib.NewTreeNode(4)
	root.Left.Right = lib.NewTreeNode(0)
	root.Left.Right.Left = lib.NewTreeNode(5)
	root.Left.Right.Right = lib.NewTreeNode(2)

	root.Right.Left = lib.NewTreeNode(1)
	root.Right.Right = lib.NewTreeNode(7)

	require.Equal(t, [][]int{{4},{9,5},{3,0,1},{8,2},{7}}, VerticalOrder(root))
}

func TestCase4(t *testing.T) {
	require.Equal(t, [][]int{}, VerticalOrder(nil))
}

func TestVerticalOrder(t *testing.T) {
	array := make([]int, 0)
	for i := 1; i <= 15; i++ {
		array = append(array, i)
	}

	root := lib.CreateMinimalBst(array)
	root.Print()
	t.Log(verticalOrder(root))

	rootRand := lib.RandomBST(10, 0, 100)
	rootRand.Print()
	t.Log(verticalOrder(rootRand))
}
