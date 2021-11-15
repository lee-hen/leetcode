package lib

import "math"

func PrintNode(root *TreeNode) {
	maxLevel := maxLevel(root)
	printNodeInternal([]*TreeNode{root}, 1, maxLevel)
}

func printNodeInternal(nodes []*TreeNode, level, maxLevel int) {
	if len(nodes) == 0 || isAllElementsNull(nodes) {
		return
	}

	floor := maxLevel - level
	edgeLines := int(math.Pow(2, float64(max(floor-1, 0))))
	firstSpaces := int(math.Pow(2, float64(floor))) - 1
	betweenSpaces := int(math.Pow(2, float64(floor+1))) - 1

	printWhitespaces(firstSpaces)

	newNodes := make([]*TreeNode, 0)
	for _, node := range nodes {
		if node != nil {
			print(node.Val)
			newNodes = append(newNodes, node.Left)
			newNodes = append(newNodes, node.Right)
		} else {
			newNodes = append(newNodes, nil)
			newNodes = append(newNodes, nil)
			print(" ")
		}
		printWhitespaces(betweenSpaces)
	}

	println("")

	for i := 1; i <= edgeLines; i++ {
		for j := 0; j < len(nodes); j++ {
			printWhitespaces(firstSpaces - i)
			if nodes[j] == nil {
				printWhitespaces(edgeLines + edgeLines + i + 1)
				continue
			}

			if nodes[j].Left != nil {
				print("/")
			} else {
				printWhitespaces(1)
			}

			printWhitespaces(i + i - 1)

			if nodes[j].Right != nil {
				print("\\")
			} else {
				printWhitespaces(1)
			}
			printWhitespaces(edgeLines + edgeLines - i)
		}
		println("")
	}

	printNodeInternal(newNodes, level+1, maxLevel)
}

func printWhitespaces(count int) {
	for i := 0; i < count; i++ {
		print(" ")
	}
}

func isAllElementsNull(list []*TreeNode) bool {
	for _, object := range list {
		if object != nil {
			return false
		}
	}
	return true
}

func maxLevel(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return max(maxLevel(node.Left), maxLevel(node.Right)) + 1
}

func max(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if curr < num {
			curr = num
		}
	}
	return curr
}
