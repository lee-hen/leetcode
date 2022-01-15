package longest_palindromic_substring

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	require.Equal(t, "bab", LongestPalindrome("babad"))
	require.Equal(t, "bb", LongestPalindrome("cbbd"))
}
