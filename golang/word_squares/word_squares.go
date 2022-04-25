package word_squares

import (
	"strings"
)

func WordSquares(words []string) [][]string {
	squares := make([][]string, 0)
	prefixHashTable := buildPrefixHashTable(words)
	for _, word := range words {
		backtracking(1, prefixHashTable, []string{word}, &squares)
	}
	return squares
}

func backtracking(step int, prefixHashTable map[string][]string, wordSquare []string, squares *[][]string) {
	if len(wordSquare) == len(wordSquare[0]) {
		cpy := make([]string, len(wordSquare))
		copy(cpy, wordSquare)
		*squares = append(*squares, cpy)
	}

	prefix := new(strings.Builder)
	for i := 0; i < len(wordSquare); i++ {
		if step < len(wordSquare[i]) {
			prefix.WriteByte(wordSquare[i][step])
		}
	}

	candidates := prefixHashTable[prefix.String()]
	for _, candidate := range candidates {
		wordSquare = append(wordSquare, candidate)
		backtracking(step+1, prefixHashTable, wordSquare, squares)
		wordSquare = wordSquare[:len(wordSquare)-1]
	}
}

func buildPrefixHashTable(words []string) map[string][]string {
	prefixHashTable := make(map[string][]string)

	for _, word := range words {
		for i := 1; i < len(words[0]); i++ {
			prefix := word[:i]
			prefixHashTable[prefix] = append(prefixHashTable[prefix], word)
		}
	}

	return prefixHashTable
}

// My solution, Accepted.
func wordSquares(words []string) [][]string {
	groupWords := make(map[byte][]string)
	for _, word := range words {
		groupWords[word[0]] = append(groupWords[word[0]], word)
	}

	squares := make([][]string, 0)
	keys := make(map[string]bool)
	for _, word := range words {
		makeWordSquares(1, groupWords, []string{word}, keys, &squares)
	}

	return squares
}

func makeWordSquares(idx int, groupWords map[byte][]string, square []string, keys map[string]bool, squares *[][]string)  {
	word := square[0]
	if len(square) == len(word) {
		if isComplete(square) {
			key := strings.Join(square, "")
			if !keys[key] {
				cpy := make([]string, len(square))
				copy(cpy, square)
				*squares = append(*squares, cpy)
				keys[key] = true
			}
			return
		}
	}

	for i := idx; i < len(word); i++ {
		candidates := groupWords[word[i]]
		for _, candidate := range candidates {
			if !isPartialOK(square) {
				return
			}

			square = append(square, candidate)
			makeWordSquares(idx+1, groupWords, square, keys, squares)
			square = square[:len(square)-1]
		}
	}
}

func isPartialOK(square []string) bool {
	if len(square) == 0 {
		return true
	}

	for i := 0; i < len(square); i++ {
		column := make([]byte, len(square))
		for j := 0; j < len(square); j++ {
			column[j] = square[j][i]
		}

		if string(column) != square[i][:len(column)] {
			return false
		}
	}
	return true
}

func isComplete(square []string) bool {
	for i := 0; i < len(square[0]); i++ {
		column := make([]byte, len(square))
		for j := 0; j < len(square); j++ {
			column[j] = square[j][i]
		}
		if string(column) != square[i] {
			return false
		}
	}

	return true
}
