package valid_anagram

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	s1 := "anagram"
	s2 := "nagaram"
 	require.True(t, IsAnagram(s1, s2))

	s1 = "rat"
	s2 = "car"
	require.False(t, IsAnagram(s1, s2))
}
