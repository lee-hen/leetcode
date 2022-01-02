package minimum_window_substring

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMinWindow(t *testing.T) {
	s1, s2 := "ADOBECODEBANC", "ABC"
	require.Equal(t, "BANC", MinWindow(s1, s2))

	s1, s2 = "a", "a"
	require.Equal(t, "a", MinWindow(s1, s2))

	s1, s2 = "a", "aa"
	require.Equal(t, "", MinWindow(s1, s2))

	s1, s2 = "a", "b"
	require.Equal(t, "", MinWindow(s1, s2))

	s1, s2 = "aa", "aa"
	require.Equal(t, "aa", MinWindow(s1, s2))

	s1, s2 = "bba", "ab"
	require.Equal(t, "ba", MinWindow(s1, s2))

	s1, s2 = "abcd$ef$axb$c$", "$$abf"
	require.Equal(t, "f$axb$", MinWindow(s1, s2))

	s1, s2 = "abcdefghijklmnopqrstuvwxyz", "aajjttwwxxzz"
	require.Equal(t, "", MinWindow(s1, s2))

	s1 = "abzacdwejxjftghiwjtklmnopqrstuvwxyz"
	s2 = "aajjttwwxxzz"
	require.Equal(t, "abzacdwejxjftghiwjtklmnopqrstuvwxyz", MinWindow(s1, s2))

	s1 = "a$fuu+afff+affaffa+a$Affab+a+a+$a$"
	s2 = "a+$aaAaaaa$++"
	require.Equal(t, "affa+a$Affab+a+a+$a", MinWindow(s1, s2))

	s1, s2 = "1456243561288281932363365412356789901!", "123!"
	require.Equal(t, "2356789901!", MinWindow(s1, s2))

	s1, s2 = "14562435612!88281932363365$412356789901", "$123!"
	require.Equal(t, "!88281932363365$", MinWindow(s1, s2))

	s1, s2 = "14562435612z!8828!193236!336!5$41!23!5!6789901#", "#!2z"
	require.Equal(t, "z!8828!193236!336!5$41!23!5!6789901#", MinWindow(s1, s2))
}
