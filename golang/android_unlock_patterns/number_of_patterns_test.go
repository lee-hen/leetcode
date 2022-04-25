package android_unlock_patterns

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNumberOfPatterns(t *testing.T) {
	require.Equal(t, 9, NumberOfPatterns(1, 1))
	require.Equal(t, 65, NumberOfPatterns(1, 2))
}

func Test_numberOfPatterns(t *testing.T) {
	require.Equal(t, 9, numberOfPatterns(1, 1))
	require.Equal(t, 56, numberOfPatterns(2, 2))
	require.Equal(t, 65, numberOfPatterns(1, 2))
	require.Equal(t, 140704, numberOfPatterns(9, 9))
	require.Equal(t, 320, numberOfPatterns(3, 3))
	require.Equal(t, 385, numberOfPatterns(1, 3))
}
