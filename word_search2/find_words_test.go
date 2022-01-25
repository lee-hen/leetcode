package word_search2

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_findWords(t *testing.T) {
	board := [][]byte{
		{'o','a','a','n'},
		{'e','t','a','e'},
		{'i','h','k','r'},
		{'i','f','l','v'},
	}
	words := []string{"oath","pea","eat","rain"}
	require.ElementsMatch(t, []string{"eat","oath"}, findWords(board, words))

	board = [][]byte{
		{'o','a','a','n'},
		{'e','t','a','e'},
		{'i','h','k','r'},
		{'i','f','l','v'},
	}
	words = []string{"pea","rain","hklf", "hf"}
	require.ElementsMatch(t, []string{"hklf", "hf"}, findWords(board, words))

	board = [][]byte{
		{'o','a','b','n'},
		{'o','t','a','e'},
		{'a','h','k','r'},
		{'a','f','l','v'},
	}
	words = []string{"oa", "oaa"}
	require.ElementsMatch(t, []string{"oa", "oaa"}, findWords(board, words))

	board = [][]byte{
		{'a','b'},
		{'c','d'},
	}
	words = []string{"abcb"}
	require.ElementsMatch(t, []string{}, findWords(board, words))

	board = [][]byte{
		{'o','a','a','n'},
		{'e','t','a','e'},
		{'i','h','k','r'},
		{'i','f','l','v'},
	}
	words = []string{"oath","pea","eat","rain","oathi","oathk","oathf","oate","oathii","oathfi","oathfii"}
	require.ElementsMatch(t, []string{"oath","oathk","oathf","oathfi","oathfii","oathi","oathii","oate","eat"}, findWords(board, words))
}
