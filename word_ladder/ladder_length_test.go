package word_ladder

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLadderLength(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot","dot","dog","lot","log","cog"}
	require.Equal(t, ladderLength(beginWord, endWord, wordList), LadderLength(beginWord,endWord, wordList))

	beginWord = "hit"
	endWord = "cog"
	wordList = []string{"hot","dot","dog","lot","log"}
	require.Equal(t, ladderLength(beginWord, endWord, wordList), LadderLength(beginWord,endWord, wordList))
}
