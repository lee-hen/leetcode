package binary_tree_vertical_order_traversal

import (
	"container/list"
	"fmt"
	"github.com/lee-hen/leecode/lib"
	"math"
	"sort"
)

var minColumn, maxColumn int
var columnTable map[int][]map[int]int

func VerticalOrder(root *lib.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	verticalOrders := make(map[int][]int)
	BFS(root, &verticalOrders)

	minColumn, maxColumn = 0, 0
	columnTable = make(map[int][]map[int]int)

	DFS(root, 0, 0)
	// Retrieve the resuts, by ordering by column and sorting by row
	fmt.Println(getVerticalOrder())
	fmt.Println(convertVerticalOrders(verticalOrders))
	return ConvertVerticalOrders(verticalOrders)
}

func BFS(root *lib.TreeNode, verticalOrders *map[int][]int) {
	queue, currIndices := list.New(), list.New()
	if root != nil {
		queue.PushBack(root)
		currIndices.PushBack(0)
	}

	for queue.Len() > 0 {
		curr := queue.Front()
		parent := curr.Value.(*lib.TreeNode)

		parentIdx := currIndices.Front().Value.(int)
		(*verticalOrders)[parentIdx] = append((*verticalOrders)[parentIdx], parent.Val)

		currIndices.Remove(currIndices.Front())
		queue.Remove(curr)

		if parent.Left != nil {
			queue.PushBack(parent.Left)
			currIndices.PushBack(parentIdx - 1)
		}

		if parent.Right != nil {
			queue.PushBack(parent.Right)
			currIndices.PushBack(parentIdx + 1)
		}
	}
}

func ConvertVerticalOrders(verticalOrders map[int][]int) [][]int {
	var minIdx, maxIdx = math.MaxInt32, math.MinInt32
	for key := range verticalOrders {
		minIdx = min(minIdx, key)
		maxIdx = max(maxIdx, key)
	}

	result := make([][]int, 0)
	for key := minIdx; key <= maxIdx; key++ {
		result = append(result, verticalOrders[key])
	}
	return result
}

func convertVerticalOrders(verticalOrders map[int][]int) [][]int {
	sorted := make([]int, len(verticalOrders))
	var i int
	for key := range verticalOrders {
		sorted[i] = key
		i++
	}

	sort.Ints(sorted)
	result := make([][]int, 0)
	for _, key := range sorted {
		result = append(result, verticalOrders[key])
	}
	return result
}

func DFS(node *lib.TreeNode, row, column int) {
	if node == nil {
		return
	}

	if _, ok := columnTable[column]; !ok {
		columnTable[column] = make([]map[int]int, 0)
	}

	columnTable[column] = append(columnTable[column], map[int]int{row: node.Val})

	minColumn = min(minColumn, column)
	maxColumn = max(maxColumn, column)

	DFS(node.Left, row + 1, column - 1)
	DFS(node.Right, row + 1, column + 1)
}

func getVerticalOrder() [][]int {
	output := make([][]int, 0)
	for i := minColumn; i <= maxColumn; i++ {
		sort.Slice(columnTable[i], func(j, k int) bool {
			var key1 int
			for key := range columnTable[i][j] {
				key1 = key
			}

			var key2 int
			for key := range columnTable[i][k] {
				key2 = key
			}

			return key1 < key2
		})

		sortedColumn := make([]int, 0)
		for _, pair := range columnTable[i] {
			for _, val := range pair {
				sortedColumn = append(sortedColumn, val)
			}
		}
		output = append(output, sortedColumn)
	}
	return output
}

func verticalOrder(root *lib.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	h := Height(root)
	size := math.Pow(2., float64(h))

	verticalOrders := make([][]int, int(size)+1)
	bfs(root, int(size/2), verticalOrders)

	result := make([][]int, 0)
	for _, order := range verticalOrders {
		if len(order) > 0 {
			result = append(result, order)
		}
	}

	return result
}

func bfs(root *lib.TreeNode, rootIdx int, verticalOrders [][]int) {
	queue, currIndices := list.New(), list.New()
	if root != nil {
		queue.PushBack(root)
		currIndices.PushBack(rootIdx)
	}

	for queue.Len() > 0 {
		curr := queue.Front()
		parent := curr.Value.(*lib.TreeNode)

		parentIdx := currIndices.Front().Value.(int)
		verticalOrders[parentIdx] = append(verticalOrders[parentIdx], parent.Val)

		currIndices.Remove(currIndices.Front())
		queue.Remove(curr)

		if parent.Left != nil {
			queue.PushBack(parent.Left)
			currIndices.PushBack(parentIdx - 1)
		}

		if parent.Right != nil {
			queue.PushBack(parent.Right)
			currIndices.PushBack(parentIdx + 1)
		}
	}
}

func Height(node *lib.TreeNode) int {
	if node == nil {
		return -1
	}

	return max(Height(node.Left), Height(node.Right)) + 1
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

func min(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if curr > num {
			curr = num
		}
	}
	return curr
}
