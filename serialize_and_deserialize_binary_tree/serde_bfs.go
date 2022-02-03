package serialize_and_deserialize_binary_tree

import (
	"container/list"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {}

func Constructor() Codec {
	return Codec{}
}

func (c *Codec) serialize(root *TreeNode) string {
	return BFS(root)
}

func (c *Codec) deserialize(data string) *TreeNode {
	strs := strings.Split(data, ";")

	str := strs[0]
	val, err := strconv.Atoi(str)
	if err != nil {
		return nil
	}
	root := &TreeNode{Val: val}

	queue := list.New()
	queue.PushBack(root)

	for i := 1; i < len(strs)-1; i+=2 {
		current := queue.Front()
		queue.Remove(current)
		curr := current.Value.(*TreeNode)

		leftIdx := i
		rightIdx := i+1

		if strs[leftIdx] == "N" && strs[rightIdx] == "N"{
			continue
		}

		if strs[leftIdx] != "N" {
			val, err = strconv.Atoi(strs[leftIdx])
			if err != nil {
				return nil
			}
			curr.Left = &TreeNode{Val: val}
			queue.PushBack(curr.Left)
		}

		if strs[rightIdx] != "N" {
			val, err = strconv.Atoi(strs[rightIdx])
			if err != nil {
				return nil
			}
			curr.Right = &TreeNode{Val: val}
			queue.PushBack(curr.Right)
		}
	}

	return root
}

func BFS(node *TreeNode) string {
	queue := list.New()
	queue.PushBack(node)

	data := new(strings.Builder)

	for queue.Len() != 0 {
		current := queue.Front()
		queue.Remove(current)

		treeNode := current.Value.(*TreeNode)

		var children []*TreeNode
		if treeNode == nil {
			data.WriteString("N;")
		} else {
			data.WriteString(strconv.Itoa(treeNode.Val))
			data.WriteString(";")
			children = []*TreeNode{treeNode.Left, treeNode.Right}
		}

		for _, child := range children {
			queue.PushBack(child)
		}
	}

	return data.String()
}
