package longest_substring_with_at_most_two_distinct_characters

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLengthOfLongestSubstringTwoDistinct(t *testing.T) {
	s := "eceba"
	require.Equal(t, 3, LengthOfLongestSubstringTwoDistinct(s))

	s = "ccaabbb"
	require.Equal(t, 5, LengthOfLongestSubstringTwoDistinct(s))

	s = "abaccc"
	require.Equal(t, 4, LengthOfLongestSubstringTwoDistinct(s))

	s = "ccaabbbddddad"
	require.Equal(t, 7, LengthOfLongestSubstringTwoDistinct(s))
}
