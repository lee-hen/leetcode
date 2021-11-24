package candy_crush

import (
	"fmt"
)

func CandyCrush(board [][]int) [][]int {
	rowLen, colLen := len(board), len(board[0])
	todo := false

	for r := 0; r < rowLen; r++ {
		for c := 0; c < colLen-2; c++ {
			v := Abs(board[r][c])
			if v != 0 && v == Abs(board[r][c+1]) && v == Abs(board[r][c+2]) {
				board[r][c], board[r][c+1], board[r][c+2] = -v, -v, -v
				todo = true
			}
		}
	}

	for r := 0; r < rowLen-2; r++ {
		for c := 0; c < colLen; c++ {
			v := Abs(board[r][c])
			if v != 0 && v == Abs(board[r+1][c]) && v == Abs(board[r+2][c]) {
				board[r][c], board[r+1][c], board[r+2][c] = -v, -v, -v
				todo = true
			}
		}
	}

	for c := 0; c < colLen; c++ {
		wr := rowLen - 1
		for r := rowLen-1; r >= 0; r-- {
			if board[r][c] > 0 {
				board[wr][c] = board[r][c]
				wr--
			}
		}
		for wr >= 0 {
			board[wr][c] = 0
			wr--
		}
	}

	if todo {
		return CandyCrush(board)
	}

	return board
}

func candyCrush(board [][]int) [][]int {
	var hadCrashed bool

	for i := range board {
		for j := range board[i] {
			if board[i][j] != 0 {
				verticals := make(map[int][]int)
				horizontals := make(map[int][]int)
				visited := make(map[string]bool)
				searchCandies(i, j, board, visited, verticals, horizontals)

				crashHorizontalCandies(horizontals, board)
				crashVerticalCandies(verticals, board)
				if len(verticals) > 0 || len(horizontals) > 0 {
					hadCrashed = true
				}
			}
		}
	}

	if hadCrashed {
		dropCandies(board)
		CandyCrush(board)
	}

	return board
}

func searchCandies(startRow, startCol int, board [][]int, visited map[string]bool, verticals map[int][]int, horizontals map[int][]int)  {
	if board[startRow][startCol] == 0 {
		return
	}

	stack := [][]int{{startRow, startCol}}
	var currentPosition []int
	for len(stack) > 0 {
		currentPosition, stack = stack[len(stack)-1], stack[:len(stack)-1]
		currentRow, currentCol := currentPosition[0], currentPosition[1]

		if !visited[generateKey(currentRow, currentCol)] {
			if len(horizontals[currentRow]) == 0 ||
				horizontals[currentRow][len(horizontals[currentRow])-1] == currentCol+1 ||
				horizontals[currentRow][len(horizontals[currentRow])-1] == currentCol-1 ||
				horizontals[currentRow][0] == currentCol+1 ||
				horizontals[currentRow][0] == currentCol-1 {
				horizontals[currentRow] = append(horizontals[currentRow], currentCol)
			}

			if len(verticals[currentCol]) == 0 ||
				verticals[currentCol][len(verticals[currentCol])-1] == currentRow+1 ||
				verticals[currentCol][len(verticals[currentCol])-1] == currentRow-1 ||
				verticals[currentCol][0] == currentRow+1 ||
				verticals[currentCol][0] == currentRow-1 {
				verticals[currentCol] = append(verticals[currentCol], currentRow)
			}
		}

		visited[generateKey(currentRow, currentCol)] = true
		neighbors := getNeighbors(board, currentRow, currentCol)
		for _, neighbor := range neighbors {
			row, col := neighbor[0], neighbor[1]
			if board[row][col] != board[startRow][startCol] || board[row][col] == 0 || visited[generateKey(row, col)] {
				continue
			}

			stack = append(stack, neighbor)
		}
	}

	for i := range horizontals {
		if len(horizontals[i]) < 3 {
			delete(horizontals, i)
		}
	}

	for i := range verticals {
		if len(verticals[i]) < 3 {
			delete(verticals, i)
		}
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

func crashVerticalCandies(verticals map[int][]int, board [][]int) {
	// vertically
	for col := range verticals {
		for _, row := range verticals[col] {
			board[row][col] = 0
		}
	}
}

func crashHorizontalCandies(horizontals map[int][]int, board [][]int) {
	// horizontally
	for row := range horizontals {
		for _, col := range horizontals[row] {
			board[row][col] = 0
		}
	}
}

func dropCandies(board [][]int) {
	for row := len(board)-1; row >= 0; row-- {
		for col := range board[row] {
			if board[row][col] == 0 {
				i := row
				j := col
				for i >= 0 && board[i][j] == 0 {
					i--
				}

				if i < 0 {
					board[row][col] = 0
				} else {
					board[row][col] = board[i][j]
					board[i][j] = 0
				}
			}
		}
	}
}

func generateKey(i, j int) string {
	return fmt.Sprintf("%d-%d", i, j)
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
