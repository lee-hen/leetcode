package longest_substring_without_repeating_characters

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "hsadfasdfasdfasdfsfdasdfsadfasdfdasfsd"
	require.Equal(t, 5, LengthOfLongestSubstring(s))

	s = "hhhsdfgfdsgfsdadfrttwergbfdshthe"
	require.Equal(t, 10, LengthOfLongestSubstring(s))

	s = "habcdahgf"
	require.Equal(t, 7, LengthOfLongestSubstring(s))

	s = "abcabcbb"
	require.Equal(t, 3, LengthOfLongestSubstring(s))

	s = "bbbbb"
	require.Equal(t, 1, LengthOfLongestSubstring(s))

	s = "pwwkew"
	require.Equal(t, 3, LengthOfLongestSubstring(s))
}
