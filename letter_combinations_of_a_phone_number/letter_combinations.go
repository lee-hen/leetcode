package letter_combinations_of_a_phone_number

import (
	"strings"
)

func LetterCombinations(digits string) []string {
	// If the input is empty, immediately return an empty answer array
	if len(digits) == 0 {
		return []string{}
	}

	// Initiate backtracking with an empty path and starting index of 0
	combinations := make([]string, 0)
	backtrack(0, digits, new(strings.Builder), &combinations)
	return combinations

}

func  backtrack(index int, digits string, path *strings.Builder, combinations *[]string) {
	// If the path is the same length as digits, we have a complete combination
	if path.Len() == len(digits) {
		*combinations = append(*combinations, path.String())
		return // Backtrack
	}

	if index >= len(digits) {
		return
	}

	// Get the letters that the current digit maps to, and loop through them
	possibleLetters := t9Letters[digits[index] - '0']
	for _, letter := range possibleLetters {
		// Add the letter to our current path
		path.WriteByte(letter)
		// Move on to the next digit
		backtrack(index + 1, digits, path, combinations)
		// Backtrack by removing the letter before moving onto the next
		cpy := path.String()
		path.Reset()
		path.WriteString(cpy[:len(cpy)-1])
	}
}

var t9Letters = [][]byte {
	nil, 					// 0
	nil, 					// 1
	{'a', 'b', 'c'}, 		// 2
	{'d', 'e', 'f'}, 		// 3
	{'g', 'h', 'i'}, 		// 4
	{'j', 'k', 'l'},		// 5
	{'m', 'n', 'o'},		// 6
	{'p', 'q', 'r', 's'}, 	// 7
	{'t', 'u', 'v'},		// 8
	{'w', 'x', 'y', 'z'}, 	// 9
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	return helper(0, digits)
}

func helper(idx int, digits string) []string {
	if idx >= len(digits) {
		return nil
	}

	combinations := helper(idx+1, digits)

	newCombinations := make([]string, 0)
	letters := t9Letters[digits[idx] - '0']

	for _, letter := range letters {
		if len(combinations) == 0 {
			newCombinations = append(newCombinations, string(letter))
		}

		for _, combination := range combinations {
			newCombinations = append(newCombinations, string(letter) + combination)
		}
	}

	return newCombinations
}
