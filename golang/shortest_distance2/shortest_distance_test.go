package shortest_distance2

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	wordsDict := []string{"practice", "makes", "perfect", "coding", "makes"}
	word1 := "coding"
	word2 := "practice"
	shortestDistance := Constructor(wordsDict)
	output := shortestDistance.Shortest(word1, word2)
	require.Equal(t, 3, output)
}

func TestCase2(t *testing.T) {
	wordsDict := []string{"practice", "makes", "perfect", "coding", "makes"}
	word1 := "makes"
	word2 := "practice"
	shortestDistance := Constructor(wordsDict)
	output := shortestDistance.Shortest(word1, word2)
	require.Equal(t, 1, output)
}

