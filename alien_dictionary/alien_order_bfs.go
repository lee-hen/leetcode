package alien_dictionary

import (
	"container/list"
	"strings"
)

func alienOrder(words []string) string {
	adjList := make(map[byte][]byte)
	counts := make(map[byte]int)

	for _, word := range words {
		for i := 0; i < len(word); i++ {
			c := word[i]
			counts[c] = 0
			adjList[c] = make([]byte, 0)
		}
	}

	for i := 0; i < len(words) - 1; i++ {
		word1 := words[i]
		word2 := words[i+1]

		if len(word1) > len(word2) && strings.HasPrefix(word1, word2) {
			return ""
		}

		for j := 0; j < Min(len(word1), len(word2)); j++ {
			if word1[j] != word2[j] {
				adjList[word1[j]] = append(adjList[word1[j]], word2[j])
				counts[word2[j]]++
				break
			}
		}
	}

	sb := new(strings.Builder)
	queue := list.New()
	for c := range counts {
		if counts[c] == 0 {
			queue.PushBack(c)
		}
	}

	for queue.Len() != 0 {
		f := queue.Front()
		queue.Remove(f)
		c := f.Value.(byte)

		sb.WriteByte(c)
		for _, next := range adjList[c] {
			counts[next]--
			if counts[next] == 0 {
				queue.PushBack(next)
			}
		}
	}

	if sb.Len() < len(counts) {
		return ""
	}
	return sb.String()
}
