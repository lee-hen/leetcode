package longest_increasing_path

import (
	"fmt"
)

type Node struct {
	i, j int
}

func LongestIncreasingPath(grid [][]int) int {
	m := len(grid)

	if m == 0 {
		return 0
	}

	n := len(grid[0])

	// padding the matrix with zero as boundaries
	// assuming all positive integer, otherwise use INT_MIN as boundaries
	matrix := make([][]int, m+2)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, n+2)
	}

	for i := 0; i < m; i++ {
		copy(matrix[i+1][1:n+1], grid[i])
	}

	// calculate outDegrees
	outDegree := make([][]int, m+2)
	for i := 0; i < len(outDegree); i++ {
		outDegree[i] = make([]int, n+2)
	}

	rowOffset := []int{-1, 0, 1, 0}
	colOffset := []int{0, 1, 0, -1}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for idx := 0; idx < 4; idx++ {
				row := i + rowOffset[idx]
				col := j + colOffset[idx]

				if matrix[i][j] < matrix[row][col] {
					outDegree[i][j]++
				}
			}
		}
	}

	// find leaves who have zero out degree as the initial level
	leaves := make([]Node, 0)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if outDegree[i][j] == 0 {
				leaves = append(leaves, Node{i,j})
			}
		}
	}

	// remove leaves level by level in topological order
	height := 0
	for len(leaves) > 0 {
		height++
		newLeaves := make([]Node, 0)
		for _, node := range leaves {
			for idx := 0; idx < 4; idx++ {
				row := node.i + rowOffset[idx]
				col := node.j + colOffset[idx]

				if matrix[node.i][node.j] > matrix[row][col] {
					outDegree[row][col]--

					if outDegree[row][col] == 0 {
						newLeaves = append(newLeaves, Node{row, col})
					}
				}
			}
		}
		leaves = newLeaves
	}

	return height
}

func longestIncreasingPath(matrix [][]int) int {
	longestPath := 0
	visited := make(map[string]int)
	for row := range matrix {
		for col := range matrix[row] {
			longestPath = Max(longestPath, dfs(row, col,  matrix, visited))
		}
	}

	return longestPath
}

func dfs(row, col int, matrix [][]int, visited map[string]int) int {
	rowOffset := []int{-1, 0, 1, 0}
	colOffset := []int{0, 1, 0, -1}

	key := fmt.Sprintf("%d, %d", row, col)

	if _, ok := visited[key]; ok {
		return visited[key]
	}

	longestPath := 0
	for idx := 0; idx < 4; idx++ {
		newRow := row + rowOffset[idx]
		newCol := col + colOffset[idx]

		if newRow < 0 || newRow >= len(matrix) || newCol < 0 || newCol >= len(matrix[row]) {
			continue
		}

		if matrix[row][col] < matrix[newRow][newCol] {
			longestPath = Max(longestPath, dfs(newRow, newCol, matrix, visited))
		}
	}

	longestPath += 1

	visited[key] = longestPath
	return longestPath
}

func Max(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if curr < num {
			curr = num
		}
	}
	return curr
}
