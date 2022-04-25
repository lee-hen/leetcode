package serialize_and_deserialize_binary_tree

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSerializeAndDeserialize(t *testing.T) {
	dataStr := "4;-7;-3;N;N;-9;-3;9;-7;-4;N;6;N;-6;-6;N;N;0;6;5;N;9;N;N;-1;-4;N;N;N;-2;N;N;N;N;N;N;N;"
	c := Constructor()
	tree := c.deserialize(dataStr)
	require.Equal(t, dataStr, c.serialize(tree))

	arr := make([]int, 0)
	tree.InOrder(&arr)


	c1 := Constructor()
	data := c1.Serialize(tree)
	t1 := c1.Deserialize(data)

	arr1 := make([]int, 0)
	t1.InOrder(&arr1)

	require.Equal(t, arr, arr1)

}

func (t *TreeNode) InOrder(arr *[]int) {
	if t == nil {
		return
	}

	t.Left.InOrder(arr)
	*arr = append(*arr, t.Val)
	t.Right.InOrder(arr)
}

