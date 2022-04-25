package count_complete_tree_nodes

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	d := depthWithoutLastLevel(root)
	if d == 0 {
		return 1
	}

	left, right := 1, int(math.Pow(2, float64(d)) -1)

	for left <= right {
		pivot := left + (right-left)/2
		if exists(pivot, d, root) {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}

	return int(math.Pow(2, float64(d)) -1) + left
}

func depthWithoutLastLevel(node *TreeNode) int {
	var d int
	for node.Left != nil {
		node = node.Left
		d++
	}
	return d
}

func exists(targetIdx, d int, node *TreeNode) bool {
	left, right := 0, int(math.Pow(2, float64(d))-1)
	var pivot int
	for i := 0; i < d; i++ {
		pivot = left + (right-left)/2
		if targetIdx <= pivot {
			node = node.Left
			right = pivot
		} else {
			node = node.Right
			left = pivot + 1
		}
	}
	return node != nil
}

func countNodes(root *TreeNode) int {
	return size(root)
}

func size(t *TreeNode) int {
	if t == nil {
		return 0
	}
	return size(t.Left) + size(t.Right) + 1
}
