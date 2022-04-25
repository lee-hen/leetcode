package word_squares

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWordSquares(t *testing.T) {
	words := []string{"area", "lead", "wall", "lady", "ball"}
	require.ElementsMatch(t, [][]string{{"ball","area","lead","lady"},{"wall","area","lead","lady"}}, WordSquares(words))

	words = []string{"abat","baba","atan","atal"}
	require.ElementsMatch(t, [][]string{{"baba","abat","baba","atal"},{"baba","abat","baba","atan"}}, WordSquares(words))
}

func Test_wordSquares(t *testing.T) {
	words := []string{"abaa", "aaab", "baaa", "aaba"}
	expected := [][]string{{"aaab", "aaba", "abaa", "baaa"}, {"aaab", "abaa", "aaba", "baaa"}, {"aaba", "aaab", "baaa", "abaa"}, {"aaba", "abaa", "baaa", "aaab"}, {"abaa", "baaa", "aaab", "aaba"}, {"abaa", "baaa", "aaba", "aaab"}, {"baaa", "aaab", "aaba", "abaa"}, {"baaa", "aaba", "abaa", "aaab"}, {"baaa", "abaa", "aaab", "aaba"}, {"baaa", "abaa", "aaba", "aaab"}}
	actual := wordSquares(words)
	require.ElementsMatch(t, expected, actual)
}
