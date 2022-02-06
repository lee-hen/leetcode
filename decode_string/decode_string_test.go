package decode_string

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDecodeString(t *testing.T) {
	s := "3[a]2[bc]"
	require.Equal(t, decodeString(s), decodeStringTwoStack(s))

	s = "3[a2[c]]"
	require.Equal(t, decodeString(s), decodeStringTwoStack(s))

	s = "2[abc]3[cd]ef"
	require.Equal(t, decodeString(s), decodeStringTwoStack(s))

	s = "3[a]2[bc]"
	idx = 0
	require.Equal(t, decodeString(s), DecodeString(s))

	s = "3[a2[c]]"
	idx = 0
	require.Equal(t, decodeString(s), DecodeString(s))

	s = "2[abc]3[cd]ef"
	idx = 0
	require.Equal(t, decodeString(s), DecodeString(s))
}
