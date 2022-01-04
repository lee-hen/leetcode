package expressive_words

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExpressiveWords(t *testing.T) {
	words := []string{"hello", "hi", "helo"}
	s := "heeellooo"
	require.Equal(t, 1, ExpressiveWords(s, words))

	words = []string{"zzyy","zy","zyy"}
	s = "zzzzzyyyyy"
	require.Equal(t, 3, ExpressiveWords(s, words))

	words = []string{"helloe", "hi", "helo"}
	s = "heeelloooeeee"
	require.Equal(t, 1, ExpressiveWords(s, words))
}
