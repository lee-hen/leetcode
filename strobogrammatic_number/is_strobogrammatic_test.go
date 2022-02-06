package strobogrammatic_number

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsStrobogrammatic(t *testing.T) {
	require.False(t, IsStrobogrammatic("2"))
	require.False(t, IsStrobogrammatic("6"))
	require.False(t, IsStrobogrammatic("659"))

	require.True(t, IsStrobogrammatic("69"))
}
