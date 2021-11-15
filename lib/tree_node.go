package lib

type TreeNode struct {
	Val, size int
	Left       *TreeNode
	Right      *TreeNode
	Parent     *TreeNode
}

func NewTreeNode(v int) *TreeNode {
	return &TreeNode{
		Val: v,
		size: 1,
	}
}

func (t *TreeNode) setLeftChild(left *TreeNode) {
	t.Left = left
	if left != nil {
		left.Parent = t
	}
}

func (t *TreeNode) setRightChild(right *TreeNode) {
	t.Right = right
	if right != nil {
		right.Parent = t
	}
}

func (t *TreeNode) InsertInOrder(d int) {
	if d <= t.Val {
		if t.Left == nil {
			t.setLeftChild(NewTreeNode(d))
		} else {
			t.Left.InsertInOrder(d)
		}
	} else {
		if t.Right == nil {
			t.setRightChild(NewTreeNode(d))
		} else {
			t.Right.InsertInOrder(d)
		}
	}
	t.size++
}

func (t *TreeNode) Size() int {
	if t == nil {
		return 0
	}
	return t.size
}

func (t *TreeNode) IsBST() bool {
	if t.Left != nil {
		if t.Val < t.Left.Val || !t.Left.IsBST() {
			return false
		}
	}

	if t.Right != nil {
		if t.Val >= t.Right.Val || !t.Right.IsBST() {
			return false
		}
	}

	return true
}

func (t *TreeNode) Height() int {
	if t == nil {
		return 0
	}

	return 1 + max(t.Left.Height(), t.Right.Height())
}

func (t *TreeNode) Find(d int) *TreeNode {
	if t != nil {
		if d == t.Val {
			return t
		} else if d <= t.Val {
			return t.Left.Find(d)
		} else if d > t.Val {
			return t.Right.Find(d)
		}
	}

	return nil
}

func (t *TreeNode) GetRandomNode() *TreeNode {
	var leftSize int
	if t.Left != nil {
		leftSize = t.Left.Size()
	}
	index := RandomInt(t.Size())

	if index < leftSize {
		return t.Left.GetRandomNode()
	} else if index == leftSize {
		return t
	} else {
		return t.Right.GetRandomNode()
	}
}

func (t *TreeNode) GetIthNode(i int) *TreeNode {
	var leftSize int
	if t.Left != nil {
		leftSize = t.Left.Size()
	}
	if i < leftSize {
		return t.Left.GetIthNode(i)
	} else if i == leftSize {
		return t
	} else {
		return t.Right.GetIthNode(i - (leftSize + 1))
	}
}

func (t *TreeNode) Print() {
	PrintNode(t)
}

func CreateMinimalBst(array []int) *TreeNode {
	return createMinimalBST(array, 0, len(array)-1)
}

func createMinimalBST(arr []int, start, end int) *TreeNode {
	if end < start {
		return nil
	}

	mid := (start + end) / 2
	n := NewTreeNode(arr[mid])
	n.setLeftChild(createMinimalBST(arr, start, mid-1))
	n.setRightChild(createMinimalBST(arr, mid+1, end))

	return n
}

func RandomBST(N, min, max int) *TreeNode {
	d := RandomIntInRange(min, max)
	root := NewTreeNode(d)
	for i := 1; i < N; i++ {
		root.InsertInOrder(RandomIntInRange(min, max))
	}
	return root
}

func CreateTreeFromArray(array []int) *TreeNode {
	if len(array) > 0 {
		root := NewTreeNode(array[0])

		queue := make([]*TreeNode, 0)
		queue = append(queue, root)
		done := false
		i := 1
		for !done {
			r := queue[0]
			if r.Left == nil {
				r.Left = NewTreeNode(array[i])
				i++
				queue = append(queue, r.Left)
			} else if r.Right == nil {
				r.Right = NewTreeNode(array[i])
				i++
				queue = append(queue, r.Left)
			} else {
				queue = queue[1:]
			}
			if i == len(array) {
				done = true
			}
		}

		return root
	} else {
		return nil
	}
}
