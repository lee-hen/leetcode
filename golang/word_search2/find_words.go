package word_search2

import (
	"fmt"
	"github.com/lee-hen/leecode/lib"
	"strings"
)

func findWords(board [][]byte, words []string) []string {
	trie := lib.NewTrie(words)
	trieNode := trie.GetRoot()

	visited := make(map[string]bool)
	var dfs func(i, j int, board [][]byte, word *strings.Builder, node *lib.TrieNode) []string
	dfs = func(i, j int, board [][]byte, word *strings.Builder, parent *lib.TrieNode) []string {
		visited[fmt.Sprintf("%d-%d", i, j)] = true

		word.WriteByte(board[i][j])
		res := make([]string, 0)
		if parent.Terminates() {
			res = append(res, word.String())
			parent.SetTerminates(false)
		}

		neighbors := getNeighbors(i, j, board, parent)
		for _, neighbor := range neighbors {
			key := fmt.Sprintf("%d-%d", neighbor[0], neighbor[1])
			children := parent.GetChild(board[neighbor[0]][neighbor[1]])

			if !visited[key] {
				res = append(res, dfs(neighbor[0], neighbor[1], board, word, children)...)
				cpy := word.String()
				word.Reset()
				word.WriteString(cpy[:len(cpy)-1])

			}
		}
		visited[fmt.Sprintf("%d-%d", i, j)] = false
		return res
	}

	result := make([]string, 0)
	for i := range board {
		for j := range board[i] {
			if len(result) == len(words) {
				break
			}

			if children := trieNode.GetChild(board[i][j]); children != nil {
				result = append(result, dfs(i, j, board, new(strings.Builder), children)...)
			}
		}
	}

	return result
}

func getNeighbors(row, col int, board [][]byte, children *lib.TrieNode) [][]int {
	neighbors := make([][]int, 0)

	rowOffset := []int{-1, 0, 1, 0}
	colOffset := []int{0, 1, 0, -1}

	for idx := 0; idx < 4; idx++ {
		newRow := row + rowOffset[idx]
		newCol := col + colOffset[idx]

		if newRow < 0 || newRow >= len(board) || newCol < 0 || newCol >= len(board[row]) {
			continue
		}

		if children.GetChild(board[newRow][newCol]) != nil {
			neighbors = append(neighbors, []int{newRow, newCol})
		}
	}

	return neighbors
}
