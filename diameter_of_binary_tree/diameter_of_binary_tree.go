package diameter_of_binary_tree

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DiameterOfBinaryTree(root *TreeNode) int {
	longestPath := math.MinInt32
	helper(root, &longestPath)
	return longestPath
}


func helper(t *TreeNode, longestPath *int) int {
	if t == nil {
		return 0
	}

	leftHeight := helper(t.Left, longestPath)
	rightHeight := helper(t.Right, longestPath)

	*longestPath = Max(*longestPath, leftHeight+rightHeight)
	return Max(leftHeight, rightHeight) + 1
}

func Max(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if curr < num {
			curr = num
		}
	}
	return curr
}
