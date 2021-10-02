package shortest_distance

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	wordsDict := []string{"practice", "makes", "perfect", "coding", "makes"}
	word1 := "makes"
	word2 := "makes"
	output := ShortestWordDistance(wordsDict, word1, word2)
	require.Equal(t, 3, output)
}

func TestCase2(t *testing.T) {
	wordsDict := []string{"practice", "makes", "perfect", "coding", "makes"}
	word1 := "makes"
	word2 := "coding"
	output := ShortestWordDistance(wordsDict, word1, word2)
	require.Equal(t, 1, output)
}
