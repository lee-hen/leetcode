package alien_dictionary

import "strings"

type state int

const (
	unvisited state = iota
	visiting
	visited
)

// The input can contain words followed by their prefix, for example, abcd and then ab.
// These cases will never result in a valid alphabet (because in a valid alphabet, prefixes are always first).
// You'll need to make sure your solution detects these cases correctly.

func AlienOrder(words []string) string {
	reverseAdjList := make(map[byte][]byte)

	for _, word := range words {
		for j := 0; j < len(word); j++ {
			c := word[j]
			reverseAdjList[c] = make([]byte, 0)
		}
	}

	for i := 1; i < len(words); i++ {
		word1 := words[i-1]
		word2 := words[i]

		if len(word1) > len(word2) && strings.HasPrefix(word1, word2) {
			return ""
		}

		for j := 0; j < Min(len(word1), len(word2)); j++ {
			if word1[j] != word2[j] {
				reverseAdjList[word2[j]] = append(reverseAdjList[word2[j]], word1[j])
				break
			}
		}
	}

	var output = new(strings.Builder)
	mark := make(map[byte]state)
	for char := range reverseAdjList {
		hasCycle := dfs(char, reverseAdjList, output, mark)
		if hasCycle {
			return ""
		}
	}

	return output.String()
}

func dfs(char byte, g map[byte][]byte, output *strings.Builder, mark map[byte]state) bool {
	if st, ok := mark[char]; ok {
		return st == visiting
	}

	mark[char] = visiting
	for _, c := range g[char] {
		hasCycle := dfs(c, g, output, mark)
		if hasCycle {
			return true
		}
	}

	output.WriteByte(char)
	mark[char] = visited
	return false
}

func Min(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if curr > num {
			curr = num
		}
	}
	return curr
}
