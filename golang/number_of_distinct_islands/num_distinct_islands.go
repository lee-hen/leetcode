package number_of_distinct_islands

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func NumDistinctIslands(grid [][]int) int {
	islands := make(map[string]struct{})
	visited := make(map[string]bool)

	for row := range grid {
		for col := range grid[row] {
			currentIsland := new(strings.Builder)
			var dfs func(row, col int, dir byte)
			dfs = func(row, col int, dir byte) {
				if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[row]) {
					return
				}

				key := generateKey(row, col)
				if visited[key] || grid[row][col] == 0 {
					return
				}

				visited[key] = true

				currentIsland.WriteByte(dir)
				dfs(row + 1, col, 'D')
				dfs(row - 1, col, 'U')
				dfs(row, col + 1, 'R')
				dfs(row, col - 1, 'L')

				// https://github.com/lee-hen/leetcode-issues/issues/31#issuecomment-980643176
				currentIsland.WriteByte('0')
			}

			dfs(row, col, '0')
			if currentIsland.Len() == 0 {
				continue
			}
			islands[currentIsland.String()] = struct{}{}
		}
	}

	return len(islands)
}

func numDistinctIslands(grid [][]int) int {
	visited := make(map[string]bool)
	islandShape := make(map[string]struct{})

	for i := range grid {
		for j := range grid[i] {
			key := generateKey(i, j)
			if grid[i][j] == 1 && !visited[key]{
				positions := make(map[int][]int)
				dfs(i, j, grid, visited, positions)

				rows := make([]int, 0)
				for row := range positions {
					rows = append(rows, row)
				}

				sort.Ints(rows)

				min := math.MaxInt32
				for _, row := range rows {
					sort.Ints(positions[row])

					if min > positions[row][0] {
						min = positions[row][0]
					}
				}

				shape := new(strings.Builder)
				for _, row := range rows {
					for _, col := range positions[row] {
						shape.WriteString(strconv.Itoa(col-min))
					}
				}
				islandShape[shape.String()] = struct{}{}
			}
		}
	}

	return len(islandShape)
}

func dfs(row, col int, grid [][]int, visited map[string]bool, positions map[int][]int) {
	key := generateKey(row, col)
	if grid[row][col] == 0 || visited[key] {
		return
	}

	visited[key] = true
	positions[row] = append(positions[row], col)
	neighbors := getNeighbors(grid, row, col)
	for _, neighbor := range neighbors {
		row, col := neighbor[0], neighbor[1]
		dfs(row, col, grid, visited, positions)
	}
}

func getNeighbors(matrix [][]int, row, col int) [][]int {
	neighbors := make([][]int, 0)
	numRows := len(matrix)
	numCols := len(matrix[row])

	if row-1 >= 0 {
		neighbors = append(neighbors, []int{row - 1, col}) // UP
	}
	if row+1 < numRows {
		neighbors = append(neighbors, []int{row + 1, col}) // DOWN
	}
	if col-1 >= 0 {
		neighbors = append(neighbors, []int{row, col - 1}) // LEFT
	}
	if col+1 < numCols {
		neighbors = append(neighbors, []int{row, col + 1}) // RIGHT
	}
	return neighbors
}

func generateKey(i, j int) string {
	return fmt.Sprintf("%d-%d", i, j)
}
