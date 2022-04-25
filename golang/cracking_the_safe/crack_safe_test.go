package cracking_the_safe

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCrackSafe(t *testing.T) {
	n, k := 2, 2
	require.Contains(t, []string{
		"01100",
		"10011",
		"11001",
		"00110",
	}, CrackSafe(n, k))
}
