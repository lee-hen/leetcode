package isomorphic_strings

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	require.True(t, IsIsomorphic("egg", "add"))
	require.False(t, IsIsomorphic("foo", "bar"))
	require.True(t, IsIsomorphic("paper", "title"))

	require.False(t, isIsomorphic("bbbaaaba", "aaabbbba"))
	require.False(t, isIsomorphic("badc", "baba"))
	require.False(t, isIsomorphic("abcdefghijklmnopqrstuvwxyzva", "abcdefghijklmnopqrstuvwxyzck"))
}
