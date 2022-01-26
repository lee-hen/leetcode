package copy_list_with_random_pointer

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func NewNode(val int, others ...int) *Node {
	l := &Node{Val: val}
	current := l
	for _, other := range others {
		current.Next = &Node{Val: other}
		current = current.Next
	}
	return l
}

func TestCopyRandomList(t *testing.T) {
	head := NewNode(7, 13, 11, 10, 1)
	lookup := make(map[int]*Node)
	for i := 0; head != nil; i++ {
		lookup[i] = head
		head = head.Next
	}

	curr := head
	var i int
	for curr != nil {
		var idx int
		switch i {
		case 1:
			idx = 0
		case 2:
			idx = 4
		case 3:
			idx = 2
		case 4:
			idx = 0
		}
		if _, ok := lookup[idx]; ok {
			curr.Random = lookup[idx]
		}
		curr = curr.Next
		i++
	}

	require.Equal(t, copyRandomList(head), CopyRandomList(head))
}
