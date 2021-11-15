package lib

import (
	"math/rand"
	"time"
)

type Tree struct {
	root *TreeNode
}

func (t *Tree) InsertInOrder(value int) {
	if t.root == nil {
		t.root = NewTreeNode(value)
	} else {
		t.root.InsertInOrder(value)
	}
}

func (t *Tree) Size() int {
	if t.root == nil {
		return 0
	}
	return t.root.Size()
}

func (t *Tree) GetRandomNode() *TreeNode {
	if t.root == nil {
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(t.Size())
	return t.root.GetIthNode(i)
}

func (t *Tree) GetRoot() *TreeNode {
	return t.root
}

func RandomInt(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func RandomIntInRange(min, max int) int {
	return RandomInt(max+1-min) + min
}
