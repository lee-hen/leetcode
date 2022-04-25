package valid_parentheses

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsValid(t *testing.T) {
	s := "()"
	require.True(t, IsValid(s))
	s = "()[]{}"
	require.True(t, IsValid(s))
	s = "(]"
	require.False(t, IsValid(s))
	s = "([{}]){"
	require.False(t, IsValid(s))
	s = "([{}]){[]}"
	require.True(t, IsValid(s))
	s = "([{[}])"
	require.False(t, IsValid(s))
	s = "]})"
	require.False(t, IsValid(s))
}
