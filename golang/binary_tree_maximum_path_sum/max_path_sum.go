package binary_tree_maximum_path_sum

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxPathSum(root *TreeNode) int {
	maxPath := math.MinInt32
	helper(root, &maxPath)
	return maxPath
}

func helper(t *TreeNode, maxPath *int) int {
	if t == nil {
		return 0
	}

	leftPath := helper(t.Left, maxPath)
	rightPath := helper(t.Right, maxPath)

	maxWithLeftAndRight := leftPath + rightPath + t.Val
	maxWithLeftOrRight := Max(Max(leftPath, rightPath) + t.Val, t.Val)
	*maxPath = Max(*maxPath, maxWithLeftAndRight, maxWithLeftOrRight)
	return maxWithLeftOrRight
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
