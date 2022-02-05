package number_of_islands

import (
	"container/list"
)

func NumIslands(grid [][]byte) int {
	var sizes int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				continue
			}

			traverseNode(i, j, grid, &sizes)
		}
	}
	return sizes
}

type node struct {
	row, col int
}

func traverseNode(i, j int, grid [][]byte, sizes *int)  {
	currentRiverSize := 0

	nodesToExplore := list.New()
	nodesToExplore.PushBack(node{i, j})

	for nodesToExplore.Len() > 0 {
		front := nodesToExplore.Front()
		nodesToExplore.Remove(front)
		currentNode := front.Value.(node)

		if grid[currentNode.row][currentNode.col] == '0' {
			continue
		}

		grid[currentNode.row][currentNode.col] = '0'

		currentRiverSize += 1
		rowOffset := []int{-1, 0, 1, 0}
		colOffset := []int{0, 1, 0, -1}

		for idx := 0; idx < 4; idx++ {
			newRow := currentNode.row + rowOffset[idx]
			newCol := currentNode.col + colOffset[idx]

			if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[currentNode.row]) {
				continue
			}

			if grid[newRow][newCol] == '1' {
				nodesToExplore.PushBack(node{newRow, newCol})
			}
		}
	}

	if currentRiverSize > 0 {
		*sizes++
	}
}
