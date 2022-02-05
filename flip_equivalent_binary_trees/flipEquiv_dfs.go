package flip_equivalent_binary_trees

import SD "github.com/lee-hen/leecode/serialize_and_deserialize_binary_tree"

func FlipEquiv(root1 *SD.TreeNode, root2 *SD.TreeNode) bool {
	vals1 := make([]*int, 0)
	vals2 := make([]*int, 0)
	dfs(root1, &vals1)
	dfs(root2, &vals2)

	if len(vals1) != len(vals2) {
		return false
	}
	for i := range vals1 {
		if vals1[i] == nil && vals2[i] != nil ||
			vals2[i] == nil && vals1[i] != nil ||
			vals1[i] != nil && vals2[i] != nil &&
			*vals1[i] != *vals2[i] {
			return false
		}
	}
	return true
}

func dfs(node *SD.TreeNode, vals *[]*int) {
	if node != nil {
		*vals = append(*vals, &node.Val)

		var l, r int
		if node.Left != nil {
			l = node.Left.Val
		} else {
			l = -1
		}

		if node.Right != nil {
			r = node.Right.Val
		} else {
			r = -1
		}

		if l < r {
			dfs(node.Left, vals)
			dfs(node.Right, vals)
		} else {
			dfs(node.Right, vals)
			dfs(node.Left, vals)
		}
		*vals = append(*vals, nil)
	}
}
