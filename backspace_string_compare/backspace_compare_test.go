package backspace_string_compare

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBackspaceCompare(t *testing.T) {
	s1 := "ab#c"
	s2 := "ad#c"
	require.True(t, BackspaceCompare(s1, s2))

	s1 = "ab##"
	s2 = "c#d#"
	require.True(t, BackspaceCompare(s1, s2))

	s1 = "a#c"
	s2 = "b"
	require.False(t, BackspaceCompare(s1, s2))
}
