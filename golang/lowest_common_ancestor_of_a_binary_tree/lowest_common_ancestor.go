package lowest_common_ancestor_of_a_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type result struct {
	node       *TreeNode
	isAncestor bool
}

func Result(n *TreeNode, isAnc bool) *result {
	r := new(result)
	r.node = n
	r.isAncestor = isAnc
	return r
}

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	r := commonAncestorHelper(root, p, q)
	if r.isAncestor {
		return r.node
	}
	return nil
}

func commonAncestorHelper(root, p, q *TreeNode) *result {
	if root == nil {
		return Result(nil, false)
	}

	if root == p && root == q {
		return Result(root, true)
	}

	rx := commonAncestorHelper(root.Left, p, q)
	if rx.isAncestor { // Found common ancestor
		return rx
	}

	ry := commonAncestorHelper(root.Right, p, q)
	if ry.isAncestor { // Found common ancestor
		return ry
	}

	if rx.node != nil && ry.node != nil {
		return Result(root, true) // This is the common ancestor
	} else if root == p || root == q {
		isAncestor := rx.node != nil || ry.node != nil
		return Result(root, isAncestor)
	} else {
		if rx.node != nil {
			return Result(rx.node, false)
		}
		return Result(ry.node, false)
	}
}
