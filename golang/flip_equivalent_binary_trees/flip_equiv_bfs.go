package flip_equivalent_binary_trees

import (
	"container/list"
	SD "github.com/lee-hen/leecode/serialize_and_deserialize_binary_tree"
)

func flipEquiv(root1 *SD.TreeNode, root2 *SD.TreeNode) bool {
	return bfs(root1, root2)
}

func bfs(t1, t2 *SD.TreeNode) bool {
	queue1 := list.New()
	queue1.PushBack(t1)

	queue2 := list.New()
	queue2.PushBack(t2)

	for queue1.Len() != 0 && queue2.Len() != 0 {
		current1 := queue1.Front()
		queue1.Remove(current1)

		current2 := queue2.Front()
		queue2.Remove(current2)

		treeNode1 := current1.Value.(*SD.TreeNode)
		treeNode2 := current2.Value.(*SD.TreeNode)

		var flip bool
		if treeNode1 != nil && treeNode2 != nil {
			if treeNode1.Val != treeNode2.Val {
				return false
			}

			if treeNode1.Left != nil && treeNode2.Right != nil &&
				treeNode1.Left.Val == treeNode2.Right.Val &&
				treeNode1.Right != nil && treeNode2.Left != nil &&
				treeNode1.Right.Val == treeNode2.Left.Val ||
				treeNode1.Left == nil && treeNode1.Left == treeNode2.Right ||
				treeNode1.Right == nil && treeNode1.Right == treeNode2.Left {
				flip = true
			}
		} else if treeNode1 == nil && treeNode2 != nil ||
			treeNode1 != nil && treeNode2 == nil	{
			return false
		}

		if treeNode1 != nil {
			queue1.PushBack(treeNode1.Left)
			queue1.PushBack(treeNode1.Right)
		}

		if treeNode2 != nil {
			if flip {
				queue2.PushBack(treeNode2.Right)
				queue2.PushBack(treeNode2.Left)
			} else {
				queue2.PushBack(treeNode2.Left)
				queue2.PushBack(treeNode2.Right)
			}
		}
	}

	return true
}
