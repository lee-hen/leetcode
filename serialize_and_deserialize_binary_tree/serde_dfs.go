package serialize_and_deserialize_binary_tree

import (
	"strconv"
	"strings"
)

func (c *Codec) Serialize(root *TreeNode) string {
	return rSerialize(root, "")
}

func rSerialize(root *TreeNode, str string) string {
	if root == nil {
		str += "N,"
	} else {
		str += strconv.Itoa(root.Val) + ","
		str = rSerialize(root.Left, str)
		str = rSerialize(root.Right, str)
	}
	return str
}

func (c *Codec) Deserialize(data string) *TreeNode {
	dataArr := strings.Split(data, ",")
	idx := 0
	return rDeserialize(dataArr, &idx)
}

func rDeserialize(dataArr []string, idx *int) *TreeNode {
	if dataArr[*idx] == "N" {
		return nil
	}

	val, _ := strconv.Atoi(dataArr[*idx])
	root := &TreeNode{Val : val}

	*idx++
	root.Left = rDeserialize(dataArr, idx)
	*idx++
	root.Right = rDeserialize(dataArr, idx)

	return root
}
